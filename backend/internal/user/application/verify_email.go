package application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// VerifyEmailRequest represents the request to verify a user's email.
type VerifyEmailRequest struct {
	Token string
}

// VerifyEmailResponse represents the response after verifying a user's email.
type VerifyEmailResponse struct {
	UserID string
}

// VerifyEmail is the use case for verifying a user's email.
type VerifyEmail struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
}

// NewVerifyEmail creates a new VerifyEmail use case.
func NewVerifyEmail(userRepo repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService) *VerifyEmail {
	return &VerifyEmail{
		UserRepository:      userRepo,
		OneTimeTokenService: oneTimeTokenService,
	}
}

// Execute handles the verification of a user's email.
func (uc *VerifyEmail) Execute(ctx context.Context, req VerifyEmailRequest) (*VerifyEmailResponse, error) {
	// Validate the token and get the user ID
	userIDStr, err := uc.OneTimeTokenService.VerifyToken(ctx, req.Token)
	if err != nil {
		return nil, fmt.Errorf("invalid or expired token: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	// Retrieve the user
	user, err := uc.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	// Mark email as verified
	user.VerifyEmail()

	// Update the user in the repository
	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &VerifyEmailResponse{UserID: user.ID.String()}, nil
}
