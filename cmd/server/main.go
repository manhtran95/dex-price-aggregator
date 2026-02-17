package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/manhtran95/dex-price-aggregator/internal/api"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize blockchain client
	client, err := blockchain.NewClient(cfg.EthereumRPC)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum: %v", err)
	}
	defer client.Close()

	// Initialize router
	router := api.NewRouter(client, cfg)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
