package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/manhtran95/dex-price-aggregator/internal/cache"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/middleware/stdlib"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	limiterRedis "github.com/ulule/limiter/v3/drivers/store/redis"
)

// NewRateLimiter creates an IP-based rate limiter.
// When redisCache is non-nil its underlying go-redis client is used for a
// distributed store; otherwise an in-process memory store is used.
func NewRateLimiter(redisCache *cache.RedisCache) *limiter.Limiter {
	var store limiter.Store
	var err error

	if redisCache != nil {
		// *redis.Client (go-redis/v9) satisfies the ulule/limiter redis.Client interface.
		store, err = limiterRedis.NewStoreWithOptions(redisCache.Client(), limiter.StoreOptions{
			Prefix:   "rate_limit",
			MaxRetry: 3,
		})
		if err != nil {
			panic(err)
		}
	} else {
		// Single-server fallback
		store = memory.NewStore()
	}

	// Rate: 10 requests per second
	rate := limiter.Rate{
		Period: 1 * time.Second,
		Limit:  10,
	}

	return limiter.New(store, rate)
}

// rateLimitExceededBody is the canonical error body returned on 429.
type rateLimitExceededBody struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	RetryAfter int    `json:"retry_after"`
}

// RateLimitMiddleware wraps the rate limiter and returns a structured JSON
// error when the limit is reached.
func RateLimitMiddleware(lmt *limiter.Limiter) func(http.Handler) http.Handler {
	limitReached := stdlib.LimitReachedHandler(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(rateLimitExceededBody{
			Error:      "Rate limit exceeded",
			Message:    "Too many requests. Try again in 1 second.",
			RetryAfter: 1,
		})
	})

	middleware := stdlib.NewMiddleware(lmt, stdlib.WithLimitReachedHandler(limitReached))

	return func(next http.Handler) http.Handler {
		return middleware.Handler(next)
	}
}
