package application

import (
	"context"
	"fmt"
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	notificationApp "github.com/jefersonprimer/chatear/backend/internal/notification/application"
	"github.com/jefersonprimer/chatear/backend/shared/events"
	"github.com/jefersonprimer/chatear/backend/shared/errors"
)

// PasswordRecoveryRequest represents the request to recover a user's password.
type PasswordRecoveryRequest struct {
	Email string
}

// PasswordRecovery is a use case for recovering a user's password.
type PasswordRecovery struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
	EventBus            repositories.EventBus
	EmailLimiter        notificationApp.RateLimiter
	AppURL              string
}

// NewPasswordRecovery creates a new PasswordRecovery use case.
func NewPasswordRecovery(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService, eventBus repositories.EventBus, emailLimiter notificationApp.RateLimiter, appURL string) *PasswordRecovery {
	return &PasswordRecovery{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
		EventBus:            eventBus,
		EmailLimiter:        emailLimiter,
		AppURL:              appURL,
	}
}

// Execute sends a password recovery email to the user.
func (uc *PasswordRecovery) Execute(ctx context.Context, req PasswordRecoveryRequest) error {
	isAllowed, err := uc.EmailLimiter.IsAllowed(ctx, req.Email)
	if err != nil {
		return fmt.Errorf("failed to check email rate limit: %w", err)
	}

	if !isAllowed {
		return errors.ErrRateLimitExceeded
	}

	user, err := uc.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil {
		return err
	}

	token, err := uc.OneTimeTokenService.GenerateToken(ctx, user.ID.String())
	if err != nil {
		return err
	}

	if err := uc.EmailLimiter.Increment(ctx, req.Email); err != nil {
		// Log or handle increment error, but proceed with sending the email
	}

	passwordRecoveryEvent := events.PasswordRecoveryRequestedEvent{
		UserID:            user.ID.String(),
		Email:             user.Email,
		Name:              user.Name,
		VerificationToken: token,
		Timestamp:         time.Now(),
		AppURL:            uc.AppURL,
	}

	if err := uc.EventBus.Publish(ctx, events.PasswordRecoveryRequestedSubject, passwordRecoveryEvent); err != nil {
		return fmt.Errorf("failed to publish PasswordRecoveryRequestedEvent: %w", err)
	}

	return nil
}
