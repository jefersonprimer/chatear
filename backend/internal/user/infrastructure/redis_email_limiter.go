package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/internal/notification/application"
)

// RedisEmailLimiter is a Redis implementation of the RateLimiter for emails.
type RedisEmailLimiter struct {
	RedisClient *redis.Client
	MaxEmails   int
}

// NewRedisEmailLimiter creates a new RedisEmailLimiter.
func NewRedisEmailLimiter(redisClient *redis.Client, cfg *config.Config) application.RateLimiter {
	return &RedisEmailLimiter{
		RedisClient: redisClient,
		MaxEmails:   cfg.MaxEmailsPerDay,
	}
}

// Get retrieves the current count for a given key.
func (l *RedisEmailLimiter) Get(ctx context.Context, key string) (int, error) {
	count, err := l.RedisClient.Get(ctx, key).Int()
	if err == redis.Nil {
		return 0, nil
	}
	if err != nil {
		return 0, fmt.Errorf("failed to get rate limit count from Redis: %w", err)
	}
	return count, nil
}

// Increment increments the count for a given key and sets an expiry if it's a new key.
func (l *RedisEmailLimiter) Increment(ctx context.Context, key string) error {
	pipe := l.RedisClient.Pipeline()
	incr := pipe.Incr(ctx, key)
	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to increment rate limit count in Redis: %w", err)
	}

	// If the key was just created (incr.Val() == 1), set its expiry to the end of the current UTC day.
	// This ensures the counter resets daily at midnight UTC.
	if incr.Val() == 1 {
		now := time.Now().UTC()
		midnight := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, time.UTC)
		durationUntilMidnight := midnight.Sub(now)
		if durationUntilMidnight <= 0 {
			// If it's already past midnight, set expiry for next midnight
			midnight = midnight.Add(24 * time.Hour)
			durationUntilMidnight = midnight.Sub(now)
		}
		if err := l.RedisClient.Expire(ctx, key, durationUntilMidnight).Err(); err != nil {
			return fmt.Errorf("failed to set expiry for rate limit key in Redis: %w", err)
		}
	}

	return nil
}

// IsAllowed checks if the given key is within the allowed rate limit.
func (l *RedisEmailLimiter) IsAllowed(ctx context.Context, key string) (bool, error) {
	count, err := l.Get(ctx, key)
	if err != nil {
		return false, err
	}
	return count < l.MaxEmails, nil
}
