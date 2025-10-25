package application

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	"github.com/jefersonprimer/chatear/backend/shared/events"
)

// DeleteUserRequest represents the request to delete a user.
type DeleteUserRequest struct {
	UserID string
}

// DeleteUser is a use case for deleting a user.
type DeleteUser struct {
	UserRepository         repositories.UserRepository
	OneTimeTokenService    services.OneTimeTokenService
	EventBus               repositories.EventBus
	UserDeletionRepository repositories.UserDeletionRepository
	AppURL                 string
}

// NewDeleteUser creates a new DeleteUser use case.
func NewDeleteUser(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService, eventBus repositories.EventBus, userDeletionRepository repositories.UserDeletionRepository, appURL string) *DeleteUser {
	return &DeleteUser{
		UserRepository:         userRepository,
		OneTimeTokenService:    oneTimeTokenService,
		EventBus:               eventBus,
		UserDeletionRepository: userDeletionRepository,
		AppURL:                 appURL,
	}
}

// Execute soft deletes a user and sends a recovery token.
func (uc *DeleteUser) Execute(ctx context.Context, req DeleteUserRequest) error {
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	user, err := uc.UserRepository.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	user.MarkAsDeleted()

	if err := uc.UserRepository.Update(ctx, user); err != nil {
		return err
	}

	recoveryToken, err := uc.OneTimeTokenService.GenerateToken(ctx, user.ID.String())
	if err != nil {
		return err
	}

	recoveryTokenExpiresAt := time.Now().Add(uc.OneTimeTokenService.GetExpiry())

	userDeletion := &entities.UserDeletion{
		ID:                     uuid.New(),
		UserID:                 user.ID,
		ScheduledDate:          time.Now().Add(90 * 24 * time.Hour), // 90 days for permanent deletion
		Status:                 entities.UserDeletionStatusScheduled,
		RecoveryToken:          &recoveryToken,
		RecoveryTokenExpiresAt: &recoveryTokenExpiresAt,
	}

	if err := uc.UserDeletionRepository.Create(ctx, userDeletion); err != nil {
		return err
	}

	accountDeletionEvent := events.AccountDeletionRequestedEvent{
		UserID:            user.ID.String(),
		Email:             user.Email,
		Name:              user.Name,
		RecoveryToken:     recoveryToken,
		Timestamp:         time.Now(),
		AppURL:            uc.AppURL,
	}

	if err := uc.EventBus.Publish(ctx, events.AccountDeletionRequestedSubject, accountDeletionEvent); err != nil {
		return fmt.Errorf("failed to publish AccountDeletionRequestedEvent: %w", err)
	}

	return nil
}
