package services

import (
	"context"
	"time"
)

// OneTimeTokenService defines the interface for generating and verifying one-time tokens.
type OneTimeTokenService interface {
	GenerateToken(ctx context.Context, userID string) (string, error)
	VerifyToken(ctx context.Context, token string) (string, error)
	GetExpiry() time.Duration
}
