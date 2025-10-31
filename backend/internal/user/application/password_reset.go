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

// PasswordResetRequest represents the request to reset a user's password.
type PasswordResetRequest struct {
	Email string
}

// PasswordReset is a use case for resetting a user's password.
type PasswordReset struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
	EventBus            repositories.EventBus
	EmailLimiter        notificationApp.RateLimiter
	AppURL              string
}

// NewPasswordReset creates a new PasswordReset use case.
func NewPasswordReset(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService, eventBus repositories.EventBus, emailLimiter notificationApp.RateLimiter, appURL string) *PasswordReset {
	return &PasswordReset{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
		EventBus:            eventBus,
		EmailLimiter:        emailLimiter,
		AppURL:              appURL,
	}
}

// Execute sends a password reset email to the user.
func (uc *PasswordReset) Execute(ctx context.Context, req PasswordResetRequest) error {
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

	passwordResetEvent := events.PasswordResetRequestedEvent{
		UserID:            user.ID.String(),
		Email:             user.Email,
		Name:              user.Name,
		VerificationToken: token,
		Timestamp:         time.Now(),
		FrontendURL:       uc.AppURL,
	}

	if err := uc.EventBus.Publish(ctx, events.PasswordResetRequestedSubject, passwordResetEvent); err != nil {
		return fmt.Errorf("failed to publish PasswordResetRequestedEvent: %w", err)
	}

	return nil
}
