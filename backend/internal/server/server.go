package server

import (
	"context"
	"l0/internal/cache"
	"l0/internal/config"
	"l0/internal/handlers"
	"l0/internal/storage"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewServer(
	ctx context.Context,
	cfg *config.Config,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
	deliveryStore storage.DeliveryStore,
	paymentStore storage.PaymentStore,
	itemsStore storage.ItemsStore,
	cacheSaverStore storage.CacheSaverStore,
) http.Handler {
	apiMux := http.NewServeMux()
	addRoutes(
		ctx,
		apiMux,
		cfg,
		orderCache,
		orderStore,
		deliveryStore,
		paymentStore,
		itemsStore,
		cacheSaverStore,
	)

	var apiHandler http.Handler = apiMux
	apiHandler = handlers.OrderCacheMiddleware(ctx, orderCache, apiHandler)
	apiHandler = handlers.CORSMiddleware(apiHandler)

	rootMux := http.NewServeMux()
	rootMux.Handle("/metrics", promhttp.Handler())
	rootMux.Handle("/health", handlers.HandleHealth(ctx))
	rootMux.Handle("/", apiHandler)

	return rootMux
}
