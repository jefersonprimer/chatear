package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
)

// RefreshTokenRepository is an interface for a refresh token repository.
type RefreshTokenRepository interface {
	CreateRefreshToken(ctx context.Context, refreshToken *entities.RefreshToken) error
	GetRefreshTokensByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.RefreshToken, error)
	GetByToken(ctx context.Context, token string) (*entities.RefreshToken, error)
	RevokeAllUserTokens(ctx context.Context, userID uuid.UUID) error
}
