package application

import (
	"context"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// RefreshToken is a use case for refreshing a token.
type RefreshToken struct {
	RefreshTokenRepository repositories.RefreshTokenRepository
	TokenService           services.TokenService
	UserRepository         repositories.UserRepository
}

// NewRefreshToken creates a new RefreshToken use case.
func NewRefreshToken(refreshTokenRepository repositories.RefreshTokenRepository, tokenService services.TokenService, userRepository repositories.UserRepository) *RefreshToken {
	return &RefreshToken{
		RefreshTokenRepository: refreshTokenRepository,
		TokenService:           tokenService,
		UserRepository:         userRepository,
	}
}

// Execute refreshes a token and returns a new access token and refresh token.
func (uc *RefreshToken) Execute(ctx context.Context, token string) (*LoginResponse, error) {
	refreshToken, err := uc.RefreshTokenRepository.GetByToken(ctx, token)
	if err != nil {
		return nil, err
	}

	user, err := uc.UserRepository.FindByID(ctx, *refreshToken.UserID)
	if err != nil {
		return nil, err
	}

	accessToken, err := uc.TokenService.GenerateAccessToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	newRefreshToken, err := uc.TokenService.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, err
	}

	return &LoginResponse{AccessToken: accessToken, RefreshToken: newRefreshToken}, nil
}
