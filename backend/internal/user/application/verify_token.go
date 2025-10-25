package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// VerifyToken is a use case for verifying a token.
type VerifyToken struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
}

// NewVerifyToken creates a new VerifyToken use case.
func NewVerifyToken(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService) *VerifyToken {
	return &VerifyToken{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
	}
}

// Execute verifies a token and performs the corresponding action.
func (uc *VerifyToken) Execute(ctx context.Context, token string) error {
	userIDStr, err := uc.OneTimeTokenService.VerifyToken(ctx, token)
	if err != nil {
		return err
	}
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}

	user, err := uc.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.IsEmailVerified = true
	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return err
	}

	return nil
}
