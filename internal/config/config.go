package config

import (
	"os"
)

type Config struct {
	EthereumRPC string
	Port        string
	Environment string
}

func Load() (*Config, error) {
	return &Config{
		EthereumRPC: getEnv("ETHEREUM_RPC", "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"),
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}