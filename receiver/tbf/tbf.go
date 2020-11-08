package tbf

import (
	"time"
)

// TokenBucket rate limiter
type TokenBucket struct {
	// parameters of tbf
	rate  float32
	burst float32

	// state
	capacity float32
	lastts   time.Time
}

// New creates a new TokenBucket rate limiter
func New(rate, burst float32) *TokenBucket {
	return &TokenBucket{
		rate:     rate / float32(time.Second),
		burst:    burst,
		capacity: burst,
		lastts:   time.Now(),
	}
}

// Get gets amount tokens from the bucket, non-blocking.
// Returns true if tokens were available, and false otherwise
func (t *TokenBucket) Get(amount float32) bool {
	t.replenish()

	if t.capacity >= amount {
		t.capacity -= amount
		return true
	}

	return false
}

func (t *TokenBucket) replenish() {
	now := time.Now()
	t.capacity += t.rate * float32(now.Sub(t.lastts))
	if t.capacity > t.burst {
		t.capacity = t.burst
	}
	t.lastts = now
}
