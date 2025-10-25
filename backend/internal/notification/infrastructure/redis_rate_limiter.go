
package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/jefersonprimer/chatear/backend/config"
)

type RedisRateLimiter struct {
	client *redis.Client
	cfg    *config.Config
}

func NewRedisRateLimiter(client *redis.Client, cfg *config.Config) *RedisRateLimiter {
	return &RedisRateLimiter{
		client: client,
		cfg:    cfg,
	}
}

func (r *RedisRateLimiter) Get(ctx context.Context, key string) (int, error) {
	key = fmt.Sprintf("rate_limit:%s", key)
	count, err := r.client.ZCard(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("error getting rate limit count from redis: %w", err)
	}
	return int(count), nil
}

func (r *RedisRateLimiter) Increment(ctx context.Context, key string) error {
	key = fmt.Sprintf("rate_limit:%s", key)
	now := time.Now().UnixNano()

	pipe := r.client.TxPipeline()

	// Remove old entries
	maxAge := now - int64(24*time.Hour) // 24 hours window
	pipe.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", maxAge))

	// Add new entry
	pipe.ZAdd(ctx, key, &redis.Z{
		Score:  float64(now),
		Member: now,
	})

	// Set expiry for the key
	pipe.Expire(ctx, key, 24*time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("error executing redis pipeline for increment: %w", err)
	}
	return nil
}

func (r *RedisRateLimiter) IsAllowed(ctx context.Context, key string) (bool, error) {
	if !r.cfg.RateLimitEnabled {
		return true, nil
	}

	count, err := r.Get(ctx, key)
	if err != nil {
		return false, err
	}

	return count < r.cfg.MaxEmailsPerDay, nil
}
