package application

import (
	"context"
	"fmt"
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// LogoutUser is a use case for logging out a user.
type LogoutUser struct {
	RefreshTokenRepository repositories.RefreshTokenRepository
	BlacklistRepository    repositories.BlacklistRepository
	TokenService           services.TokenService
}

// NewLogoutUser creates a new LogoutUser use case.
func NewLogoutUser(refreshTokenRepository repositories.RefreshTokenRepository, blacklistRepository repositories.BlacklistRepository, tokenService services.TokenService) *LogoutUser {
	return &LogoutUser{
		RefreshTokenRepository: refreshTokenRepository,
		BlacklistRepository:    blacklistRepository,
		TokenService:           tokenService,
	}
}

// Execute logs out a user by invalidating all their tokens.
func (uc *LogoutUser) Execute(ctx context.Context, accessToken string) error {
	userID, err := uc.TokenService.VerifyToken(ctx, accessToken)
	if err != nil {
		return err
	}

	// Add the access token to the blacklist
	accessTokenExp, err := uc.TokenService.GetTokenExpiration(accessToken)
	if err != nil {
		return fmt.Errorf("failed to get access token expiration: %w", err)
	}

	err = uc.BlacklistRepository.Add(ctx, accessToken, time.Until(accessTokenExp))
	if err != nil {
		return fmt.Errorf("failed to add access token to blacklist: %w", err)
	}

	refreshTokens, err := uc.RefreshTokenRepository.GetRefreshTokensByUserID(ctx, userID)
	if err != nil {
		return err
	}

	for _, refreshToken := range refreshTokens {
		// Add the refresh token to the blacklist
		expiration := time.Until(refreshToken.ExpiresAt)
		err = uc.BlacklistRepository.Add(ctx, refreshToken.Token, expiration)
		if err != nil {
			// Log the error but continue
		}
	}

	// Revoke all refresh tokens for the user
	return uc.RefreshTokenRepository.RevokeAllUserTokens(ctx, userID)
}
