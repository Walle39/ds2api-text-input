package account

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mu              sync.Mutex
	lastRequestTime map[string]time.Time
	interval        time.Duration
}

func NewRateLimiter(intervalMs int) *RateLimiter {
	interval := time.Duration(intervalMs) * time.Millisecond
	if interval <= 0 {
		interval = 500 * time.Millisecond
	}
	return &RateLimiter{
		lastRequestTime: map[string]time.Time{},
		interval:        interval,
	}
}

func (r *RateLimiter) Wait(accountID string) {
	if accountID == "" {
		return
	}
	r.mu.Lock()
	defer r.mu.Unlock()

	lastTime, exists := r.lastRequestTime[accountID]
	if exists {
		elapsed := time.Since(lastTime)
		if elapsed < r.interval {
			time.Sleep(r.interval - elapsed)
		}
	}
	r.lastRequestTime[accountID] = time.Now()
}

func (r *RateLimiter) UpdateInterval(intervalMs int) {
	r.mu.Lock()
	defer r.mu.Unlock()
	interval := time.Duration(intervalMs) * time.Millisecond
	if interval > 0 {
		r.interval = interval
	}
}