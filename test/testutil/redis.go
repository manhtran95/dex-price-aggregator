package testutil

import (
	"testing"

	"github.com/alicebob/miniredis/v2"
	"github.com/manhtran95/dex-price-aggregator/internal/cache"
)

func NewTestRedisCache(t *testing.T) *cache.RedisCache {
	// Start in-memory Redis
	mr, err := miniredis.Run()
	if err != nil {
		t.Fatalf("Failed to start miniredis: %v", err)
	}

	// Clean up after test
	t.Cleanup(func() {
		mr.Close()
	})

	// Create cache client
	redisCache, err := cache.NewRedisCache(mr.Addr(), "", 0)
	if err != nil {
		t.Fatalf("Failed to create Redis cache: %v", err)
	}

	t.Cleanup(func() {
		redisCache.Close()
	})

	return redisCache
}
