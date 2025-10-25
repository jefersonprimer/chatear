package entities

import (
	"time"

	"github.com/google/uuid"
)

// EmailType represents the type of email being sent
type EmailType string

const (
	EmailTypeVerification  EmailType = "verification"
	EmailTypePasswordReset EmailType = "password_reset"
	EmailTypeAccountDeletion EmailType = "account_deletion"
)

// EmailSend represents an email send record
type EmailSend struct {
	ID           uuid.UUID              `json:"id"`
	UserID       *uuid.UUID             `json:"user_id,omitempty"`
	Type         EmailType              `json:"type"`
	Recipient    string                 `json:"recipient"`
	Subject      string                 `json:"subject"`
	TemplateName string                 `json:"template_name"`
	TemplateData map[string]interface{} `json:"template_data"`
	SentAt       time.Time              `json:"sent_at"`
}

// NewEmailSend creates a new email send record
func NewEmailSend(recipient, subject, templateName string, templateData map[string]interface{}) *EmailSend {
	return &EmailSend{
		ID:           uuid.New(),
		Recipient:    recipient,
		Subject:      subject,
		TemplateName: templateName,
		TemplateData: templateData,
		SentAt:       time.Now(),
	}
}
