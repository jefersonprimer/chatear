package services

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// TokenService defines the interface for token-related operations.
type TokenService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	VerifyToken(ctx context.Context, tokenString string) (uuid.UUID, error)
	GetTokenExpiration(tokenString string) (time.Time, error)
	GetRefreshTokenTTL() time.Duration
}