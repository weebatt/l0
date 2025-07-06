package main

import (
	"context"
	"database/sql"
	"embed"
	"l0/internal/config"
	"l0/pkg/db/postgres"
	"l0/pkg/logger"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	// Initialize logger
	ctx := context.Background()
	ctx, err := logger.New(ctx)
	if err != nil {
		logger.GetFromContext(ctx).Fatal("failed to initialize logger", zap.Error(err))
	}

	// Set up env var
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

	// Initialize http server
	logger.GetFromContext(ctx).Info("starting server...")

	server := config.NewServer(cfg.Server)
	done := make(chan bool)

	go GracefulShutdown(server, db, done, cfg)

	if err := server.ListenAndServe(); err != nil {
		logger.GetFromContext(ctx).Fatal("failed to start http server", zap.Error(err))
	} else {
		logger.GetFromContext(ctx).Info("http server started successfully")
	}

	<-done
	logger.GetFromContext(ctx).Info("server gracefully shutdown completed")
}

func GracefulShutdown(serverApi *http.Server, db *sql.DB, done chan bool, cfg *config.Config) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	defer stop()

	<-ctx.Done()

	logger.GetFromContext(ctx).Info("shutting down server gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.ShutdownTimeout))
	defer cancel()
	if err := serverApi.Shutdown(ctx); err != nil {
		logger.GetFromContext(ctx).Error("failed to shutdown http server gracefully", zap.Error(err))
	} else {
		logger.GetFromContext(ctx).Info("http server shutdown gracefully")
	}

	if err := db.Close(); err != nil {
		logger.GetFromContext(ctx).Error("failed to close database connection", zap.Error(err))
	} else {
		logger.GetFromContext(ctx).Info("successfully closed database connection")
	}

	done <- true
}
