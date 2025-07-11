package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"io"
	"l0/internal/cache"
	"l0/internal/config"
	"l0/internal/models"
	"l0/internal/server"
	"l0/internal/storage"
	"l0/pkg/db/postgres"
	"l0/pkg/kafka"
	"l0/pkg/logger"
	metrics "l0/pkg/metrics"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	validator "l0/internal/validator"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	if err := run(ctx, os.Args, os.Getenv, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

//go:embed migrations/*.sql
var embedMigrations embed.FS

func run(
	ctx context.Context,
	args []string,
	getenv func(string) string,
	stdout, stderr io.Writer,
) error {
	// Register collecting metrics
	metrics.Register()

	// Initialize logger
	ctx, err := logger.New(ctx)
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed to initialize logger", zap.Error(err))
	}

	// Initialize configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed with reading form .env or setup env vars into the structure ", zap.Error(err))
	}

	logger.GetFromContext(ctx).Info("successfully read and setup env vars")

	// Initialize connection to postgres
	db, err := postgres.New(cfg)
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed to open postgres connection", zap.Error(err))
	}

	if err := db.Ping(); err != nil {
		logger.GetFromContext(ctx).Fatal("failed to ping postgres connection", zap.Error(err))
	}

	logger.GetFromContext(ctx).Info("successfully connected to postgres")

	// Initialize migrations
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		logger.GetFromContext(ctx).Fatal("failed to set postgres dialect", zap.Error(err))
	}

	if err := goose.Up(db, "migrations"); err != nil {
		logger.GetFromContext(ctx).Fatal("failed to up migrations", zap.Error(err))
	}

	logger.GetFromContext(ctx).Info("successfully up migrations")

	// Initialize stores
	orderStore := storage.NewOrderStore(db)
	deliveryStore := storage.NewDeliveryStore(db)
	paymentStore := storage.NewPaymentStore(db)
	itemStore := storage.NewItemsStore(db)
	cacheSaverStore := storage.NewCacheSaverStore(db)
	orderCache := cache.NewOrderCache()

	// Restore cache from postgres
	RestoreCache(
		ctx,
		orderCache,
		orderStore,
		deliveryStore,
		paymentStore,
		itemStore,
		cacheSaverStore,
	)

	// Start consume messages from kafka
	ConsumeMessagesFromKafka(
		ctx,
		cfg,
		db,
		orderStore,
		deliveryStore,
		paymentStore,
		itemStore,
		cacheSaverStore,
	)

	// Initialize http server
	srv := server.NewServer(
		ctx,
		cfg,
		orderCache,
		orderStore,
		deliveryStore,
		paymentStore,
		itemStore,
		cacheSaverStore,
	)

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Handler: srv,
	}

	// Gracefull shutdown
	done := make(chan bool)

	go GracefulShutdown(ctx, server, db, done, cfg)

	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.GetFromContext(ctx).Fatal("failed to start http server", zap.Error(err))
	} else if err == nil {
		logger.GetFromContext(ctx).Info("http server started successfully")
	}

	<-done
	logger.GetFromContext(ctx).Info("server gracefully shutdown completed")
	return nil
}

func RestoreCache(
	ctx context.Context,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
	deliveryStore storage.DeliveryStore,
	paymentStore storage.PaymentStore,
	itemStore storage.ItemsStore,
	cacheSaverStore storage.CacheSaverStore,
) {
	orderUIDs, err := cacheSaverStore.GetAllCachedOrderUIDs(ctx)
	if err != nil {
		logger.GetFromContext(ctx).Error("failed to load cache_saver order_uids", zap.Error(err))
	} else {
		var orders []*models.Order
		for _, uid := range orderUIDs {
			order, err := orderStore.GetOrder(ctx, uid)
			if err == nil && order != nil {
				delivery, _ := deliveryStore.GetDelivery(ctx, uid)
				if delivery != nil {
					order.Delivery = *delivery
				}
				payment, _ := paymentStore.GetPayment(ctx, uid)
				if payment != nil {
					order.Payment = *payment
				}
				items, _ := itemStore.GetItems(ctx, uid)
				if items != nil {
					order.Items = *items
				}
				orders = append(orders, order)
			}
		}
		orderCache.Load(orders)
		logger.GetFromContext(ctx).Info("order cache loaded from cache_saver", zap.Int("count", len(orders)))
	}
}

func ConsumeMessagesFromKafka(
	ctx context.Context,
	cfg *config.Config,
	db *sql.DB,
	orderStore storage.OrderStore,
	deliveryStore storage.DeliveryStore,
	paymentStore storage.PaymentStore,
	itemStore storage.ItemsStore,
	cacheSaverStore storage.CacheSaverStore,
) {
	go func() {
		kafka.StartOrderConsumer(ctx, cfg, func(ctx context.Context, order *models.Order) error {
			tx, err := db.BeginTx(ctx, nil)
			if err != nil {
				logger.GetFromContext(ctx).Error("failed to begin transaction", zap.Error(err))
				return err
			}

			err = validator.ValidateOrder(order)
			if err != nil {
				tx.Rollback()
				logger.GetFromContext(ctx).Error("failed in validation order", zap.Error(err))
				return err
			}

			err = orderStore.CreateOrder(
				ctx,
				order.OrderUID,
				order.TrackNumber,
				order.Entry,
				order.Locale,
				order.CustomerID,
				order.DeliveryService,
				order.Shardkey,
				order.SmID,
				order.DateCreated,
				order.OofShard,
				order.InternalSignature,
			)
			if err != nil {
				tx.Rollback()
				logger.GetFromContext(ctx).Error("failed to save order from kafka", zap.Error(err))
				return err
			}

			if (order.Delivery != models.Delivery{}) {
				order.Delivery.OrderUID = order.OrderUID

				err = validator.ValidateDelivery(&order.Delivery)

				if err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed in validation delivery", zap.Error(err))
					return err
				}

				if err := deliveryStore.CreateDelivery(ctx, &order.Delivery); err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed to save delivery from kafka", zap.Error(err))
					return err
				}
			}

			if (order.Payment != models.Payment{}) {
				order.Payment.Transaction = order.OrderUID

				err = validator.ValidatePayment(&order.Payment)
				if err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed iin validation payment", zap.Error(err))
					return err
				}

				if err := paymentStore.CreatePayment(ctx, &order.Payment); err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed to save payment from kafka", zap.Error(err))
					return err
				}
			}

			if len(order.Items) > 0 {
				for i := range order.Items {
					order.Items[i].OrderUID = order.OrderUID
				}

				err = validator.ValidateItems(order.Items)
				if err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed in validation items", zap.Error(err))
					return err
				}

				if err := itemStore.CreateItems(ctx, order.Items); err != nil {
					tx.Rollback()
					logger.GetFromContext(ctx).Error("failed to save items from kafka", zap.Error(err))
					return err
				}
			}

			if err := tx.Commit(); err != nil {
				logger.GetFromContext(ctx).Error("failed to commit transaction", zap.Error(err))
				return err
			}

			logger.GetFromContext(ctx).Info("Order and related data saved from kafka", zap.Any("order_uid", order.OrderUID))
			return nil
		})
	}()

}

func GracefulShutdown(origCtx context.Context, serverApi *http.Server, db *sql.DB, done chan bool, cfg *config.Config) {
	ctx, stop := signal.NotifyContext(origCtx, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer stop()

	<-ctx.Done()

	logger.GetFromContext(origCtx).Info("shutting down server gracefully")

	ctxTimeout, cancel := context.WithTimeout(origCtx, time.Duration(cfg.Server.ShutdownTimeout))
	defer cancel()
	if err := serverApi.Shutdown(ctxTimeout); err != nil {
		logger.GetFromContext(origCtx).Error("failed to shutdown http server gracefully", zap.Error(err))
	} else {
		logger.GetFromContext(origCtx).Info("http server shutdown gracefully")
	}

	if err := db.Close(); err != nil {
		logger.GetFromContext(origCtx).Error("failed to close database connection", zap.Error(err))
	} else {
		logger.GetFromContext(origCtx).Info("successfully closed database connection")
	}

	done <- true
	logger.GetFromContext(origCtx).Info("Kafka config", zap.Any("brokers", cfg.Kafka.Brokers), zap.String("topic", cfg.Kafka.Topic))
}
