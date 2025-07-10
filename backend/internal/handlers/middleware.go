package handlers

import (
	"encoding/json"
	"l0/internal/cache"
	"l0/internal/metrics"
	"net/http"

	"github.com/google/uuid"
)

func OrderCacheMiddleware(orderCache *cache.OrderCache, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("order_uid")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid order_uid", http.StatusBadRequest)
			return
		}
		order, ok := orderCache.Get(id)
		if ok {
			metrics.OrderCacheHits.Inc()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(order)
			return
		}
		metrics.OrderCacheMisses.Inc()
		next.ServeHTTP(w, r)
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
