package ratelimiter

import (
	"context"
	"time"
)

type TokenBucketLimiter struct {
	tokenBucketCh chan struct{}
}

func NewTokenBucketLimiter(ctx context.Context, limit int, period time.Duration) *TokenBucketLimiter {
	limiter := &TokenBucketLimiter{
		tokenBucketCh: make(chan struct{}, limit),
	}

	for i := 0; i < limit; i++ {
		limiter.tokenBucketCh <- struct{}{}
	}

	replineshmentInterval := period.Nanoseconds() / int64(limit)
	go limiter.startPeriodicReplishement(ctx, time.Duration(replineshmentInterval))

	return limiter
}

func (l *TokenBucketLimiter) startPeriodicReplishement(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			l.tokenBucketCh <- struct{}{}
		}
	}
}

func (l *TokenBucketLimiter) Allow() bool {
	select {
	case <-l.tokenBucketCh:
		return true
	default:
		return false
	}
}
