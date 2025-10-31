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

// PasswordRecoveryConsumer consumes password recovery events and sends a magic link email.
type PasswordRecoveryConsumer struct {
	emailService *application.EmailService
}

// NewPasswordRecoveryConsumer creates a new PasswordRecoveryConsumer.
func NewPasswordRecoveryConsumer(emailService *application.EmailService) *PasswordRecoveryConsumer {
	return &PasswordRecoveryConsumer{
		emailService: emailService,
	}
}

// Consume consumes password recovery events from NATS.
func (c *PasswordRecoveryConsumer) Consume(ctx context.Context, msg *nats.Msg) {
	var event events.PasswordRecoveryRequestedEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("Error unmarshalling password recovery event: %v", err)
		return
	}

	magicLink := fmt.Sprintf("%s/recover-account?token=%s", event.AppURL, event.VerificationToken)

	if err := c.emailService.SendMagicLinkEmail(ctx, event.Email, event.UserID, "Password Reset", magicLink); err != nil {
		log.Printf("Error sending password recovery email for user %s: %v", event.UserID, err)
	}
}
