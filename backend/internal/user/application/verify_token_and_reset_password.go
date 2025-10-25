package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	"golang.org/x/crypto/bcrypt"
)

// VerifyTokenAndResetPassword is a use case for verifying a token and resetting a user's password.
type VerifyTokenAndResetPassword struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
}

// NewVerifyTokenAndResetPassword creates a new VerifyTokenAndResetPassword use case.
func NewVerifyTokenAndResetPassword(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService) *VerifyTokenAndResetPassword {
	return &VerifyTokenAndResetPassword{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
	}
}

// Execute verifies a token and resets a user's password.
func (uc *VerifyTokenAndResetPassword) Execute(ctx context.Context, token string, newPassword string) (*entities.User, error) {
	userIDStr, err := uc.OneTimeTokenService.VerifyToken(ctx, token)
	if err != nil {
		return nil, err
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return nil, err
	}

	user, err := uc.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.PasswordHash = string(hashedPassword)
	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
