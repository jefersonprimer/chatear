package domain

import (
	"time"
)

// EmailSend represents an email that has been sent or is pending to be sent.
type EmailSend struct {
	ID           string                 `json:"id"`
	Recipient    string                 `json:"recipient"`
	Subject      string                 `json:"subject"`
	TemplateName string                 `json:"templateName"`
	TemplateData map[string]interface{} `json:"templateData"`
	Body         string                 `json:"body"` // Rendered email body
	SentAt       time.Time              `json:"sentAt"`
	Status       string                 `json:"status"` // e.g., "pending", "sent", "failed"
	ErrorMessage string                 `json:"errorMessage,omitempty"`
}