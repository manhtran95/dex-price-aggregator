package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	// Test connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisCache{client: client}, nil
}

func (r *RedisCache) Close() error {
	return r.client.Close()
}

// Pool Reserves Caching
type ReservesCache struct {
	Reserve0  string `json:"reserve0"`
	Reserve1  string `json:"reserve1"`
	Timestamp int64  `json:"timestamp"`
}

func (r *RedisCache) GetReserves(ctx context.Context, poolAddress common.Address) (*big.Int, *big.Int, error) {
	key := fmt.Sprintf("reserves:%s", poolAddress.Hex())

	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil, nil // Cache miss
	}
	if err != nil {
		return nil, nil, err
	}

	var cached ReservesCache
	if err := json.Unmarshal([]byte(val), &cached); err != nil {
		return nil, nil, err
	}

	reserve0 := new(big.Int)
	reserve1 := new(big.Int)
	reserve0.SetString(cached.Reserve0, 10)
	reserve1.SetString(cached.Reserve1, 10)

	return reserve0, reserve1, nil
}

func (r *RedisCache) SetReserves(
	ctx context.Context,
	poolAddress common.Address,
	reserve0, reserve1 *big.Int,
	ttl time.Duration,
) error {
	key := fmt.Sprintf("reserves:%s", poolAddress.Hex())

	cached := ReservesCache{
		Reserve0:  reserve0.String(),
		Reserve1:  reserve1.String(),
		Timestamp: time.Now().Unix(),
	}

	data, err := json.Marshal(cached)
	if err != nil {
		return err
	}

	return r.client.Set(ctx, key, data, ttl).Err()
}

// Token Metadata Caching
type TokenMetadata struct {
	Symbol   string `json:"symbol"`
	Name     string `json:"name"`
	Decimals uint8  `json:"decimals"`
}

func (r *RedisCache) GetTokenMetadata(ctx context.Context, tokenAddress common.Address) (*TokenMetadata, error) {
	key := fmt.Sprintf("token:%s", tokenAddress.Hex())

	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil // Cache miss
	}
	if err != nil {
		return nil, err
	}

	var metadata TokenMetadata
	if err := json.Unmarshal([]byte(val), &metadata); err != nil {
		return nil, err
	}

	return &metadata, nil
}

func (r *RedisCache) SetTokenMetadata(
	ctx context.Context,
	tokenAddress common.Address,
	symbol, name string,
	decimals uint8,
) error {
	key := fmt.Sprintf("token:%s", tokenAddress.Hex())

	metadata := TokenMetadata{
		Symbol:   symbol,
		Name:     name,
		Decimals: decimals,
	}

	data, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	// Token metadata rarely changes - cache for 24 hours
	return r.client.Set(ctx, key, data, 24*time.Hour).Err()
}

// Gas Price Caching
func (r *RedisCache) GetGasPrice(ctx context.Context) (float64, error) {
	val, err := r.client.Get(ctx, "gas:price:gwei").Result()
	if err == redis.Nil {
		return 0, nil // Cache miss
	}
	if err != nil {
		return 0, err
	}

	var gasPrice float64
	if err := json.Unmarshal([]byte(val), &gasPrice); err != nil {
		return 0, err
	}

	return gasPrice, nil
}

func (r *RedisCache) SetGasPrice(ctx context.Context, gasPriceGwei float64) error {
	data, err := json.Marshal(gasPriceGwei)
	if err != nil {
		return err
	}

	// Gas price changes frequently - cache for 30 seconds
	return r.client.Set(ctx, "gas:price:gwei", data, 30*time.Second).Err()
}

// Generic cache for any data
func (r *RedisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	}
	return val, err
}

func (r *RedisCache) Set(ctx context.Context, key string, value interface{}, ttl time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, key, data, ttl).Err()
}

// Clear cache (useful for testing/debugging)
func (r *RedisCache) FlushAll(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}
