package infrastructure

import (
	"context"
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/internal/notification/application"
	"github.com/jefersonprimer/chatear/backend/internal/notification/domain"
)

// SMTPSender sends emails using SMTP.
type SMTPSender struct {
	host           string
	port           string
	username       string
	password       string
	from           string
	templateParser application.TemplateParser
}

// NewSMTPSender creates a new SMTPSender.
func NewSMTPSender(cfg *config.Config, templateParser application.TemplateParser) *SMTPSender {
	return &SMTPSender{
		host:           cfg.SMTPHost,
		port:           strconv.Itoa(cfg.SMTPPort),
		username:       cfg.SMTPUser,
		password:       cfg.SMTPPass,
		from:           cfg.SMTPFrom,
		templateParser: templateParser,
	}
}

// Send sends an email using SMTP.
func (s *SMTPSender) Send(ctx context.Context, email *domain.EmailSend) error {
	// Render the email body from the template
	body, err := s.templateParser.ParseTemplate(email.TemplateName, email.TemplateData)
	if err != nil {
		return fmt.Errorf("failed to render email template: %w", err)
	}

	auth := smtp.PlainAuth("", s.username, s.password, s.host)

	msg := []byte("From: " + s.from + "\r\n" +
		"To: " + email.Recipient + "\r\n" +
		"Subject: " + email.Subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n" +
		"\r\n" +
		body + "\r\n")

	addr := fmt.Sprintf("%s:%s", s.host, s.port)

	return smtp.SendMail(addr, auth, s.from, []string{email.Recipient}, msg)
}
