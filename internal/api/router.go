package api

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
	"github.com/rs/cors"
)

func NewRouter(client *blockchain.Client, cfg *config.Config) http.Handler {
	r := mux.NewRouter()

	// Initialize handlers
	h := NewHandlers(client, cfg)

	// API routes
	api := r.PathPrefix("/api/v1").Subrouter()
	api.HandleFunc("/health", h.HealthCheck).Methods("GET")
	api.HandleFunc("/quote", h.GetQuote).Methods("POST")
	api.HandleFunc("/compare", h.ComparePrices).Methods("POST")

	// CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

	return c.Handler(r)
}
