package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
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

// Get retorna a contagem atual de ações no rate limit
func (r *RedisRateLimiter) Get(ctx context.Context, key string) (int, error) {
	key = fmt.Sprintf("rate_limit:%s", key)
	count, err := r.client.ZCard(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("error getting rate limit count from redis: %w", err)
	}
	return int(count), nil
}

// Increment adiciona uma entrada e remove antigas, controlando o rate limit
func (r *RedisRateLimiter) Increment(ctx context.Context, key string) error {
	key = fmt.Sprintf("rate_limit:%s", key)
	now := time.Now().UnixNano()

	pipe := r.client.TxPipeline()

	// Remove entradas antigas (janela de 24h)
	maxAge := now - int64(24*time.Hour)
	pipe.ZRemRangeByScore(ctx, key, "-inf", fmt.Sprintf("%d", maxAge))

	// Adiciona nova entrada
	pipe.ZAdd(ctx, key, redis.Z{
		Score:  float64(now),
		Member: now,
	})

	// Define tempo de expiração para a chave
	pipe.Expire(ctx, key, 24*time.Hour)

	_, err := pipe.Exec(ctx)
	if err != nil {
		return fmt.Errorf("error executing redis pipeline for increment: %w", err)
	}
	return nil
}

// IsAllowed verifica se a ação é permitida segundo o limite configurado
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

