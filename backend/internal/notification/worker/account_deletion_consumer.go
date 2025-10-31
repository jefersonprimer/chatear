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

// AccountDeletionConsumer consumes account deletion events and sends a recovery email.
type AccountDeletionConsumer struct {
	emailService *application.EmailService
}

// NewAccountDeletionConsumer creates a new AccountDeletionConsumer.
func NewAccountDeletionConsumer(emailService *application.EmailService) *AccountDeletionConsumer {
	return &AccountDeletionConsumer{
		emailService: emailService,
	}
}

// Consume consumes account deletion events from NATS.
func (c *AccountDeletionConsumer) Consume(ctx context.Context, msg *nats.Msg) {
	var event events.AccountDeletionRequestedEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("Error unmarshalling account deletion event: %v", err)
		return
	}

	recoveryLink := fmt.Sprintf("%s/recover-account?token=%s", event.FrontendURL, event.RecoveryToken)

	if err := c.emailService.SendMagicLinkEmail(ctx, event.Email, event.UserID, "Account Deletion Confirmation and Recovery", recoveryLink); err != nil {
		log.Printf("Error sending account deletion recovery email for user %s: %v", event.UserID, err)
	}
}
