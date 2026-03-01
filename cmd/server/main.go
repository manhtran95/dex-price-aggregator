package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/manhtran95/dex-price-aggregator/internal/aggregator"
	"github.com/manhtran95/dex-price-aggregator/internal/api"
	"github.com/manhtran95/dex-price-aggregator/internal/blockchain"
	"github.com/manhtran95/dex-price-aggregator/internal/cache"
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

	// Initialize Redis cache
	redisCache, err := cache.NewRedisCache(cfg.RedisAddress, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Printf("Warning: Redis not available, running without cache: %v", err)
		redisCache = nil // Run without cache
	} else {
		defer redisCache.Close()
		log.Printf("✓ Redis cache connected at %s", cfg.RedisAddress)
	}

	// Initialize DEXes
	uniswapV2, err := dex.NewUniswapV2(client.Client(), redisCache)
	if err != nil {
		log.Fatalf("Failed to initialize Uniswap V2: %v", err)
	}

	uniswapV3, err := dex.NewUniswapV3(client.Client(), redisCache)
	if err != nil {
		log.Fatalf("Failed to initialize Uniswap V3: %v", err)
	}

	sushiSwap, err := dex.NewSushiSwap(client.Client(), redisCache)
	if err != nil {
		log.Fatalf("Failed to initialize SushiSwap: %v", err)
	}

	curve, err := dex.NewCurve(client.Client(), redisCache)
	if err != nil {
		log.Fatalf("Failed to initialize Curve: %v", err)
	}

	dexes := []dex.DEX{
		uniswapV2,
		uniswapV3,
		sushiSwap,
		curve,
	}

	// Initialize aggregator
	agg := aggregator.NewAggregator(dexes, client.Client(), redisCache)

	// Initialize router
	router := api.NewRouter(client, cfg, agg, redisCache)

	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
