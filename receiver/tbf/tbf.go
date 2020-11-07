package tbf

import (
	"time"
)

// TokenBucket rate limiter
type TokenBucket struct {
	// parameters of tbf
	rate  float64
	burst float64

	// state
	capacity float64
	lastts   time.Time
}

// New creates a new TokenBucket rate limiter
func New(rate, burst float64) *TokenBucket {
	return &TokenBucket{
		rate:     rate / float64(time.Second),
		burst:    burst,
		capacity: burst,
		lastts:   time.Now(),
	}
}

// Get gets amount tokens from the bucket, non-blocking.
// Returns true if tokens were available, and false otherwise
func (t *TokenBucket) Get(amount float64) bool {
	t.replenish()

	if t.capacity >= amount {
		t.capacity -= amount
		return true
	}

	return false
}

func (t *TokenBucket) replenish() {
	now := time.Now()
	t.capacity += t.rate * float64(now.Sub(t.lastts))
	if t.capacity > t.burst {
		t.capacity = t.burst
	}
	t.lastts = now
}
