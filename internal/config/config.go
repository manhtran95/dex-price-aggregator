package config

import (
	"os"
	"strconv"
)

type Config struct {
	EthereumRPC string
	Port        string
	Environment string

	RedisAddress  string
	RedisPassword string
	RedisDB       int
}

func Load() (*Config, error) {
	return &Config{
		EthereumRPC:   getEnv("ETHEREUM_MAINNET_RPC", "https://mainnet.infura.io/v3/YOUR-PROJECT-ID"),
		Port:          getEnv("PORT", "8080"),
		Environment:   getEnv("ENVIRONMENT", "development"),
		RedisAddress:  getEnv("REDIS_ADDRESS", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvInt("REDIS_DB", 0),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
