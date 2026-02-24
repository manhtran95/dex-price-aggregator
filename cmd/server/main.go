package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/manhtran95/dex-price-aggregator/internal/aggregator"
	"github.com/manhtran95/dex-price-aggregator/internal/api"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/config"
	"github.com/manhtran95/dex-price-aggregator/internal/dex"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	client, err := blockchain.NewClient(cfg.EthereumRPC)
	if err != nil {
		log.Fatalf("Failed to connect to Ethereum: %v", err)
	}
	defer client.Close()

	// Initialize DEXes
	uniswapV2, err := dex.NewUniswapV2(client.Client())
	if err != nil {
		log.Fatalf("Failed to initialize Uniswap V2: %v", err)
	}

	uniswapV3, err := dex.NewUniswapV3(client.Client())
	if err != nil {
		log.Fatalf("Failed to initialize Uniswap V3: %v", err)
	}

	sushiSwap, err := dex.NewSushiSwap(client.Client())
	if err != nil {
		log.Fatalf("Failed to initialize SushiSwap: %v", err)
	}

	dexes := []dex.DEX{
		uniswapV2,
		uniswapV3,
		sushiSwap,
	}

	// Initialize aggregator
	agg := aggregator.NewAggregator(dexes, client.Client())

	// Initialize router - pass aggregator in
	router := api.NewRouter(client, cfg, agg)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
