package server

import (
	"context"
	"l0/internal/cache"
	"l0/internal/config"
	"l0/internal/handlers"
	"l0/internal/storage"
	"net/http"
)

func addRoutes(
	ctx context.Context,
	mux *http.ServeMux,
	cfg *config.Config,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
	deliveryStore storage.DeliveryStore,
	paymentStore storage.PaymentStore,
	itemsStore storage.ItemsStore,
	cacheSaverStore storage.CacheSaverStore,
) {
	mux.Handle("/order", handlers.OrderHandler(ctx, orderCache, orderStore, deliveryStore, paymentStore, itemsStore, cacheSaverStore))
}
