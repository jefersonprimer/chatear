package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/jefersonprimer/chatear/backend/internal/notification/application"
	"github.com/jefersonprimer/chatear/backend/shared/events"
	"github.com/nats-io/nats.go"
)

// PasswordResetConsumer consumes password reset events and sends a magic link email.
type PasswordResetConsumer struct {
	emailService *application.EmailService
}

// NewPasswordResetConsumer creates a new PasswordResetConsumer.
func NewPasswordResetConsumer(emailService *application.EmailService) *PasswordResetConsumer {
	return &PasswordResetConsumer{
		emailService: emailService,
	}
}

// Consume consumes password reset events from NATS.
func (c *PasswordResetConsumer) Consume(ctx context.Context, msg *nats.Msg) {
	var event events.PasswordResetRequestedEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("Error unmarshalling password reset event: %v", err)
		return
	}

	magicLink := fmt.Sprintf("%s/auth/reset-password?token=%s", event.FrontendURL, event.VerificationToken)

	if err := c.emailService.SendMagicLinkEmail(ctx, event.Email, event.UserID, "Password Reset", magicLink); err != nil {
		log.Printf("Error sending password reset email for user %s: %v", event.UserID, err)
	}
}
