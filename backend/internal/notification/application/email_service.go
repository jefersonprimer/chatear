package application

import (
	"context"
	"fmt"
	"time"


	"github.com/jefersonprimer/chatear/backend/domain/services"
	"github.com/jefersonprimer/chatear/backend/shared/errors"
	notificationDomain "github.com/jefersonprimer/chatear/backend/internal/notification/domain"
)

// EmailService orchestrates sending different types of emails.
type EmailService struct {
	emailSender         notificationDomain.Sender
	oneTimeTokenService services.OneTimeTokenService
	appURL              string
	magicLinkExpiry     time.Duration
	emailRateLimiter    RateLimiter
}

// NewEmailService creates a new EmailService.
func NewEmailService(emailSender notificationDomain.Sender, oneTimeTokenService services.OneTimeTokenService, appURL string, magicLinkExpiry time.Duration, emailRateLimiter RateLimiter) *EmailService {
	return &EmailService{
		emailSender:         emailSender,
		oneTimeTokenService: oneTimeTokenService,
		appURL:              appURL,
		magicLinkExpiry:     magicLinkExpiry,
		emailRateLimiter:    emailRateLimiter,
	}
}

// SendMagicLinkEmail sends a magic link email for email verification, password recovery, or account recovery.
func (s *EmailService) SendMagicLinkEmail(ctx context.Context, recipient, userID, subject, link string) error {
	// Enforce rate limit
	isAllowed, err := s.emailRateLimiter.IsAllowed(ctx, recipient)
	if err != nil {
		return fmt.Errorf("failed to check email rate limit: %w", err)
	}
	if !isAllowed {
		return errors.ErrTooManyEmailAttempts
	}

	data := map[string]interface{}{
		"Subject":   subject,
		"Recipient": recipient,
		"Link":      link,
		"ExpiryMinutes": s.magicLinkExpiry.Minutes(),
	}

	emailSend := &notificationDomain.EmailSend{
		Recipient:    recipient,
		Subject:      subject,
		TemplateName: "magic_link.html", // Use a generic magic link template
		TemplateData: data,
	}

	if err := s.emailSender.Send(ctx, emailSend); err != nil {
		return fmt.Errorf("failed to send magic link email: %w", err)
	}

	return nil
}
