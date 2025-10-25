package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/jefersonprimer/chatear/backend/internal/notification/domain"
)

var ErrRateLimitExceeded = errors.New("rate limit exceeded")

type EmailSender struct {
	repository  domain.Repository
	sender      domain.Sender
	rateLimiter RateLimiter
}

func NewEmailSender(repository domain.Repository, sender domain.Sender, rateLimiter RateLimiter) *EmailSender {
	return &EmailSender{
		repository:  repository,
		sender:      sender,
		rateLimiter: rateLimiter,
	}
}

func (s *EmailSender) Send(ctx context.Context, emailSend *domain.EmailSend) error {
	allowed, err := s.rateLimiter.IsAllowed(ctx, emailSend.Recipient)
	if err != nil {
		return err
	}
	if !allowed {
		return ErrRateLimitExceeded
	}

	// Increment the rate limit counter after it's confirmed that the email can be sent
	if err := s.rateLimiter.Increment(ctx, emailSend.Recipient); err != nil {
		return fmt.Errorf("failed to increment email rate limit: %w", err)
	}

	// Ensure ID is set if not already
	if emailSend.ID == "" {
		emailSend.ID = uuid.New().String()
	}
	emailSend.SentAt = time.Now()
	emailSend.Status = "pending"

	if err := s.sender.Send(ctx, emailSend); err != nil {
		emailSend.Status = "failed"
		emailSend.ErrorMessage = err.Error()
		if saveErr := s.repository.Save(ctx, emailSend); saveErr != nil {
			return saveErr
		}
		return err
	}

	emailSend.Status = "sent"
	if err := s.repository.Save(ctx, emailSend); err != nil {
		return err
	}

	return nil
}
