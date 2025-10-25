package application

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// RecoverAccountRequest represents the request to recover a user account with a new password.
type RecoverAccountRequest struct {
	Token       string
	NewPassword string
}

// RecoverAccount is a use case for recovering a user account by setting a new password.
type RecoverAccount struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
}

// NewRecoverAccount creates a new RecoverAccount use case.
func NewRecoverAccount(userRepository repositories.UserRepository, _ repositories.UserDeletionRepository, oneTimeTokenService services.OneTimeTokenService) *RecoverAccount {
	return &RecoverAccount{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
	}
}

// Execute handles the recovery of a user account by setting a new password.
func (uc *RecoverAccount) Execute(ctx context.Context, req RecoverAccountRequest) error {
	// Validate the token and get the user ID
	userIDStr, err := uc.OneTimeTokenService.VerifyToken(ctx, req.Token)
	if err != nil {
		return fmt.Errorf("invalid or expired token: %w", err)
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return fmt.Errorf("invalid user ID in token: %w", err)
	}

	// Retrieve the user
	user, err := uc.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// Hash the new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash new password: %w", err)
	}

	// Update the user's password hash
	user.PasswordHash = string(hashedPassword)

	// Restore the user account
	user.Restore()

	// Update the user in the repository
	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return fmt.Errorf("failed to update user password and restore account: %w", err)
	}

	return nil
}
