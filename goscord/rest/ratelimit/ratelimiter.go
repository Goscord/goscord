package ratelimit

import "sync"

type RateLimiter struct {
	mu       sync.RWMutex
	ralimits map[string]*RateLimit
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		ralimits: make(map[string]*RateLimit),
	}
}

func (r *RateLimiter) Get(bucket string) *RateLimit {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if rl, ok := r.ralimits[bucket]; ok {
		return rl
	}

	return nil
}

func (r *RateLimiter) Set(bucket string, rl *RateLimit) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ralimits[bucket] = rl
}

func (r *RateLimiter) Remove(bucket string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(r.ralimits, bucket)
}

func (r *RateLimiter) Clear() {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.ralimits = make(map[string]*RateLimit)
}
