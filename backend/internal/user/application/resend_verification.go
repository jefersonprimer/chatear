package application

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	"github.com/jefersonprimer/chatear/backend/shared/events"
)

// ResendVerificationEmail is a use case for resending the verification email.
type ResendVerificationEmail struct {
	UserRepository      repositories.UserRepository
	OneTimeTokenService services.OneTimeTokenService
	EventBus            repositories.EventBus
	AppURL              string
}

// NewResendVerificationEmail creates a new ResendVerificationEmail use case.
func NewResendVerificationEmail(userRepository repositories.UserRepository, oneTimeTokenService services.OneTimeTokenService, eventBus repositories.EventBus, appURL string) *ResendVerificationEmail {
	return &ResendVerificationEmail{
		UserRepository:      userRepository,
		OneTimeTokenService: oneTimeTokenService,
		EventBus:            eventBus,
		AppURL:              appURL,
	}
}

// Execute finds a user by email, generates a new verification token, and sends it.
func (uc *ResendVerificationEmail) Execute(ctx context.Context, email string) error {
	user, err := uc.UserRepository.FindByEmail(ctx, email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.IsEmailVerified {
		return errors.New("email already verified")
	}

	token, err := uc.OneTimeTokenService.GenerateToken(ctx, user.ID.String())
	if err != nil {
		return err
	}

	emailRequest := events.EmailSendRequest{
		Recipient: user.Email,
		Subject:   "Email Verification",
		Body:      fmt.Sprintf("Click here to verify your email: %s/verify-email?token=%s", uc.AppURL, token),
	}
	emailDataBytes, err := json.Marshal(emailRequest)
	if err != nil {
		return err
	}

	if err := uc.EventBus.Publish(ctx, "email.send", emailDataBytes); err != nil {
		return err
	}

	return nil
}
