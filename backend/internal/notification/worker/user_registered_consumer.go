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

// UserRegisteredConsumer consumes user registration events and sends a verification email.
type UserRegisteredConsumer struct {
	emailService *application.EmailService
	appURL       string
}

// NewUserRegisteredConsumer creates a new UserRegisteredConsumer.
func NewUserRegisteredConsumer(emailService *application.EmailService, appURL string) *UserRegisteredConsumer {
	return &UserRegisteredConsumer{
		emailService: emailService,
		appURL:       appURL,
	}
}

// Consume consumes user registration events from NATS.
func (c *UserRegisteredConsumer) Consume(msg *nats.Msg) {
	var event events.UserRegisteredEvent
	if err := json.Unmarshal(msg.Data, &event); err != nil {
		log.Printf("Error unmarshalling user registered event: %v", err)
		return
	}

	verificationLink := fmt.Sprintf("%s/verify-email?token=%s", c.appURL, event.VerificationToken)

	if err := c.emailService.SendMagicLinkEmail(context.Background(), event.Email, event.UserID, "Verify your email", verificationLink); err != nil {
		log.Printf("Error sending verification email for user %s: %v", event.UserID, err)
	}
}
