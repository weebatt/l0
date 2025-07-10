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

func addRoutes(
	ctx context.Context,
	mux *http.ServeMux,
	cfg *config.Config,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
) {
	mux.Handle("/health", handlers.HandleHealth(ctx))
	mux.Handle("/order", handlers.OrderHandler(orderCache, orderStore))
	mux.Handle("/metrics", promhttp.Handler())
}
