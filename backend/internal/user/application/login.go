package application

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	"github.com/jefersonprimer/chatear/backend/shared/errors"
)

// LoginRequest represents the request to log in a user.
type LoginRequest struct {
	Email    string
	Password string
}

// LoginResponse represents the response after a successful login.
type LoginResponse struct {
	User         *entities.User
	AccessToken  string
	RefreshToken string
}

// Login is the use case for user login.
type Login struct {
	UserRepository      repositories.UserRepository
	TokenService        services.TokenService
	RefreshTokenRepo    repositories.RefreshTokenRepository
}

// NewLogin creates a new Login use case.
func NewLogin(userRepo repositories.UserRepository, tokenService services.TokenService, refreshTokenRepo repositories.RefreshTokenRepository) *Login {
	return &Login{
		UserRepository:      userRepo,
		TokenService:        tokenService,
		RefreshTokenRepo:    refreshTokenRepo,
	}
}

// Execute handles the user login process.
func (uc *Login) Execute(ctx context.Context, req LoginRequest) (*LoginResponse, error) {
	// Retrieve the user by email
	user, err := uc.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	// Compare the provided password with the stored hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.ErrInvalidCredentials
	}

	// Check if the user's email is verified
	if !user.IsEmailVerified {
		return nil, errors.ErrEmailNotVerified
	}

	// Generate access token
	accessToken, err := uc.TokenService.GenerateAccessToken(user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	// Generate refresh token
	refreshToken, err := uc.TokenService.GenerateRefreshToken(user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// Store refresh token
	refreshTokenEntity := entities.NewRefreshToken(&user.ID, refreshToken, time.Now().Add(uc.TokenService.GetRefreshTokenTTL()), nil, nil)
	if err := uc.RefreshTokenRepo.CreateRefreshToken(ctx, refreshTokenEntity); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	// Update user's last login timestamp
	user.UpdateLastLogin()
	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user last login: %w", err)
	}

	return &LoginResponse{
		User:         user,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
