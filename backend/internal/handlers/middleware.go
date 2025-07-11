package handlers

import (
	"context"
	"encoding/json"
	"l0/internal/cache"
	"l0/internal/metrics"
	"time"

	"l0/pkg/logger"
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

func OrderCacheMiddleware(ctx context.Context, orderCache *cache.OrderCache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("order_uid")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid order_uid", http.StatusBadRequest)
			return
		}

		start := time.Now()
		order, ok := orderCache.Get(id)

		if ok {
			logger.GetFromContext(ctx).Info("order found in cache", zap.String("order_uid", idStr))
			metrics.OrderCacheHits.Inc()

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)

			cacheDuration := time.Since(start).Seconds()
			metrics.CacheLookupDuration.Observe(cacheDuration)
			metrics.OrderCacheLookupDuration.Set(cacheDuration)
			return
		}

		logger.GetFromContext(ctx).Info("order not found in cache", zap.String("order_uid", idStr))
		metrics.OrderCacheMisses.Inc()

		next.ServeHTTP(w, r)

		logger.GetFromContext(ctx).Info("order loaded from DB", zap.String("order_uid", idStr))
	})
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
