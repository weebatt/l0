package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"l0/internal/cache"
	"l0/internal/storage"

	"github.com/google/uuid"
)

func OrderHandler(orderCache *cache.OrderCache, orderStore storage.OrderStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("order_uid")
		id, err := uuid.Parse(idStr)
		if err != nil {
			http.Error(w, "invalid order_uid", http.StatusBadRequest)
			return
		}
		order, err := orderStore.GetOrder(context.Background(), id)
		if err != nil || order == nil {
			http.Error(w, "order not found", http.StatusNotFound)
			return
		}
		orderCache.Set(order)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(order)
	})
}
