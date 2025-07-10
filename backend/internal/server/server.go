package server

import (
	"context"
	"l0/internal/cache"
	"l0/internal/config"
	"l0/internal/handlers"
	"l0/internal/storage"
	"net/http"
)

func NewServer(
	ctx context.Context,
	cfg *config.Config,
	orderCache *cache.OrderCache,
	orderStore storage.OrderStore,
) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		ctx,
		mux,
		cfg,
		orderCache,
		orderStore,
	)

	var handler http.Handler = mux
	handler = handlers.OrderCacheMiddleware(orderCache, handler)
	handler = handlers.CORSMiddleware(handler)
	return handler
}
