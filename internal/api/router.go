package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manhtran95/dex-price-aggregator/internal/aggregator"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/cache"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
	"github.com/rs/cors"
)

func NewRouter(client *blockchain.Client, cfg *config.Config, agg *aggregator.Aggregator, redisCache *cache.RedisCache) http.Handler {
	r := mux.NewRouter()

	// redisCache may be nil (Redis unavailable) – NewRateLimiter falls back to in-memory store.
	rateLimiter := NewRateLimiter(redisCache)
	r.Use(RateLimitMiddleware(rateLimiter))

	h := NewHandlers(client, cfg, agg)

	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", h.HealthCheck).Methods("GET")
	api.HandleFunc("/quote", h.GetRouteWithGas).Methods("POST")
	api.HandleFunc("/compare", h.ComparePrices).Methods("POST")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	return c.Handler(r)
}
