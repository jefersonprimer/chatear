package worker

import (
	"context"
	"encoding/json"
	"log"

	"github.com/jefersonprimer/chatear/backend/internal/notification/domain"
	"github.com/jefersonprimer/chatear/backend/shared/events"
	"github.com/nats-io/nats.go"
)

// EmailConsumer consumes email send events and sends emails.
type EmailConsumer struct {
	emailSender domain.Sender
}

// NewEmailConsumer creates a new EmailConsumer.
func NewEmailConsumer(emailSender domain.Sender) *EmailConsumer {
	return &EmailConsumer{
		emailSender: emailSender,
	}
}

// Consume consumes email send events from NATS.
func (c *EmailConsumer) Consume(ctx context.Context, msg *nats.Msg) {
	var req events.EmailSendRequest
	if err := json.Unmarshal(msg.Data, &req); err != nil {
		log.Printf("Error unmarshalling email send request: %v", err)
		return
	}

	email := &domain.EmailSend{
		Recipient:    req.Recipient,
		Subject:      req.Subject,
		Body:         req.Body,
		TemplateName: req.TemplateName,
	}

	if err := c.emailSender.Send(ctx, email); err != nil {
		log.Printf("Error sending email: %v", err)
	}
}
