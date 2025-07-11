package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"l0/internal/cache"
	"l0/internal/storage"
	"l0/pkg/logger"
	"l0/pkg/metrics"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func OrderHandler(
	ctx context.Context,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
	deliveryStore storage.DeliveryStore,
	paymentStore storage.PaymentStore,
	itemsStore storage.ItemsStore,
	cacheSaverStore storage.CacheSaverStore,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("order_uid")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid order_uid", http.StatusBadRequest)
			return
		}

		startDB := time.Now()

		order, err := orderStore.GetOrder(context.Background(), id)
		if err != nil {
			logger.GetFromContext(ctx).Error("failed to get order", zap.String("order_uid", idStr), zap.Error(err))
			http.Error(w, "failed to get order", http.StatusInternalServerError)
			return
		} else if order != nil {
			logger.GetFromContext(ctx).Info("order found in postgres", zap.Any("order", *order))
		}

		delivery, err := deliveryStore.GetDelivery(context.Background(), id)
		if err != nil {
			logger.GetFromContext(ctx).Error("failed to get delivery", zap.String("order_uid", idStr), zap.Error(err))
		} else if delivery != nil {
			logger.GetFromContext(ctx).Info("delivery found in postgres", zap.Any("delivery", *delivery))
			order.Delivery = *delivery
		}

		payment, err := paymentStore.GetPayment(context.Background(), id)
		if err != nil {
			logger.GetFromContext(ctx).Error("failed to get delivery", zap.String("order_uid", idStr), zap.Error(err))
		} else if payment != nil {
			logger.GetFromContext(ctx).Info("payment found in postgres", zap.Any("payment", *payment))
			order.Payment = *payment
		}

		items, err := itemsStore.GetItems(context.Background(), id)
		if err != nil {
			logger.GetFromContext(ctx).Error("failed to get items", zap.String("order_uid", idStr), zap.Error(err))
		} else if items != nil {
			logger.GetFromContext(ctx).Info("items found in postgres", zap.Any("items", *items))
			order.Items = *items
		}

		orderCache.Set(order)
		_ = cacheSaverStore.AddOrderUID(r.Context(), order.OrderUID)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)

		dbDuration := time.Since(startDB).Seconds()
		metrics.DBLookupDuration.Observe(dbDuration)
		metrics.OrderDBLookupDuration.Set(dbDuration)

		logger.GetFromContext(ctx).Info("order take from postgres", zap.String("order_uid", idStr))
	})
}
