package application

import "context"

type RateLimiter interface {
	Get(ctx context.Context, key string) (int, error)
	Increment(ctx context.Context, key string) error
	IsAllowed(ctx context.Context, key string) (bool, error)
}
