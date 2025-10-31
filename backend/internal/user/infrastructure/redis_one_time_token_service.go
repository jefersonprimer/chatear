package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// RedisOneTimeTokenService is a Redis implementation of the OneTimeTokenService.
type RedisOneTimeTokenService struct {
	RedisClient *redis.Client
	Expiry      time.Duration
}

// NewRedisOneTimeTokenService creates a new RedisOneTimeTokenService.
func NewRedisOneTimeTokenService(redisClient *redis.Client, cfg *config.Config) services.OneTimeTokenService {
	return &RedisOneTimeTokenService{
		RedisClient: redisClient,
		Expiry:      cfg.MagicLinkExpiry,
	}
}

// GenerateToken generates a new one-time token and stores it in Redis with a TTL.
func (s *RedisOneTimeTokenService) GenerateToken(ctx context.Context, userID string) (string, error) {
	token := uuid.New().String()
	key := fmt.Sprintf("one_time_token:%s", token)

	err := s.RedisClient.Set(ctx, key, userID, s.Expiry).Err()
	if err != nil {
		return "", fmt.Errorf("failed to store one-time token in Redis: %w", err)
	}
	return token, nil
}

// VerifyToken verifies a one-time token and returns the associated user ID.
// The token is deleted from Redis after successful verification.
func (s *RedisOneTimeTokenService) VerifyToken(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("one_time_token:%s", token)

	userID, err := s.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("one-time token not found or expired")
	}
	if err != nil {
		return "", fmt.Errorf("failed to retrieve one-time token from Redis: %w", err)
	}

	// Delete the token after use
	err = s.RedisClient.Del(ctx, key).Err()
	if err != nil {
		return "", fmt.Errorf("failed to delete one-time token from Redis: %w", err)
	}

	return userID, nil
}

// PeekToken checks if a one-time token exists and returns the associated user ID without deleting it.
func (s *RedisOneTimeTokenService) PeekToken(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("one_time_token:%s", token)

	userID, err := s.RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("one-time token not found or expired")
	}
	if err != nil {
		return "", fmt.Errorf("failed to retrieve one-time token from Redis: %w", err)
	}

	return userID, nil
}

// GetExpiry returns the configured expiry duration for one-time tokens.
func (s *RedisOneTimeTokenService) GetExpiry() time.Duration {
	return s.Expiry
}
