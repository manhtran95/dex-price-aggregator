package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// newTestLimiter builds a limiter backed by the in-memory store.
// All test requests share the same store so quota is consumed correctly.
func newTestLimiter(limit int64, period time.Duration) *limiter.Limiter {
	store := memory.NewStore()
	rate := limiter.Rate{Period: period, Limit: limit}
	return limiter.New(store, rate)
}

// okHandler is a minimal 200 handler used as the downstream target.
var okHandler = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
})

// sendRequest fires one GET /api/v1/health request with a fixed IP so all
// requests in the test count against the same rate-limit bucket.
func sendRequest(t *testing.T, handler http.Handler) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/health", nil)
	req.RemoteAddr = "10.0.0.1:12345" // fixed IP → same bucket for every call
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	return rr
}

// TestRateLimit_11thRequestIsRejected verifies:
//   - requests 1-10 pass through (HTTP 200)
//   - request 11 is rejected (HTTP 429) with the canonical JSON error body
func TestRateLimit_11thRequestIsRejected(t *testing.T) {
	const rateLimit = 10

	lmt := newTestLimiter(rateLimit, time.Second)
	handler := RateLimitMiddleware(lmt)(okHandler)

	// First 10 requests must succeed.
	for i := 1; i <= rateLimit; i++ {
		rr := sendRequest(t, handler)
		assert.Equal(t, http.StatusOK, rr.Code, "request %d should be allowed", i)
	}

	// 11th request must be rate-limited.
	rr := sendRequest(t, handler)
	require.Equal(t, http.StatusTooManyRequests, rr.Code, "11th request should be rate-limited")

	// Decode and validate the JSON body.
	var body rateLimitExceededBody
	err := json.NewDecoder(rr.Body).Decode(&body)
	require.NoError(t, err, "response body should be valid JSON")

	assert.Equal(t, "Rate limit exceeded", body.Error)
	assert.Equal(t, "Too many requests. Try again in 1 second.", body.Message)
	assert.Equal(t, 1, body.RetryAfter)

	// Content-Type header should signal JSON.
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

// TestRateLimit_AllowsRequestAfterWindowExpires verifies that the quota
// resets after the rate-limit period so requests succeed again.
func TestRateLimit_AllowsRequestAfterWindowExpires(t *testing.T) {
	const rateLimit = 2
	const period = 200 * time.Millisecond

	lmt := newTestLimiter(rateLimit, period)
	handler := RateLimitMiddleware(lmt)(okHandler)

	// Exhaust the quota.
	for i := 0; i < rateLimit; i++ {
		rr := sendRequest(t, handler)
		assert.Equal(t, http.StatusOK, rr.Code)
	}

	// Next request hits the limit.
	rr := sendRequest(t, handler)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)

	// Wait for the window to expire then retry – should succeed.
	time.Sleep(period + 50*time.Millisecond)

	rr = sendRequest(t, handler)
	assert.Equal(t, http.StatusOK, rr.Code, "request after window reset should succeed")
}
