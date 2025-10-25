package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jefersonprimer/chatear/backend/internal/notification/domain"
)

// PostgresEmailSendRepository implements domain.Repository for PostgreSQL.
type PostgresEmailSendRepository struct {
	DB *pgxpool.Pool
}

// NewPostgresEmailSendRepository creates a new PostgresEmailSendRepository.
func NewPostgresEmailSendRepository(db *pgxpool.Pool) *PostgresEmailSendRepository {
	return &PostgresEmailSendRepository{DB: db}
}

// Save inserts or updates an EmailSend record in the database.
func (r *PostgresEmailSendRepository) Save(ctx context.Context, emailSend *domain.EmailSend) error {
	templateDataJSON, err := json.Marshal(emailSend.TemplateData)
	if err != nil {
		return fmt.Errorf("failed to marshal template data: %w", err)
	}

	query := `
		INSERT INTO email_sends (id, recipient, subject, template_name, template_data, body, sent_at, status, error_message)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		ON CONFLICT (id) DO UPDATE SET
			recipient = EXCLUDED.recipient,
			subject = EXCLUDED.subject,
			template_name = EXCLUDED.template_name,
			template_data = EXCLUDED.template_data,
			body = EXCLUDED.body,
			sent_at = EXCLUDED.sent_at,
			status = EXCLUDED.status,
			error_message = EXCLUDED.error_message
	`

	_, err = r.DB.Exec(ctx, query,
		emailSend.ID,
		emailSend.Recipient,
		emailSend.Subject,
		emailSend.TemplateName,
		templateDataJSON,
		emailSend.Body,
		emailSend.SentAt,
		emailSend.Status,
		emailSend.ErrorMessage,
	)
	if err != nil {
		return fmt.Errorf("failed to save email send record: %w", err)
	}

	return nil
}

// GetByID retrieves an EmailSend record by its ID.
func (r *PostgresEmailSendRepository) GetByID(ctx context.Context, id string) (*domain.EmailSend, error) {
	emailSend := &domain.EmailSend{}
	var templateDataJSON []byte

	query := `
		SELECT id, recipient, subject, template_name, template_data, body, sent_at, status, error_message
		FROM email_sends
		WHERE id = $1
	`

	err := r.DB.QueryRow(ctx, query, id).Scan(
		&emailSend.ID,
		&emailSend.Recipient,
		&emailSend.Subject,
		&emailSend.TemplateName,
		&templateDataJSON,
		&emailSend.Body,
		&emailSend.SentAt,
		&emailSend.Status,
		&emailSend.ErrorMessage,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to get email send record by ID: %w", err)
	}

	if err := json.Unmarshal(templateDataJSON, &emailSend.TemplateData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal template data: %w", err)
	}

	return emailSend, nil
}

// GetByRecipient retrieves a list of EmailSend records by recipient.
func (r *PostgresEmailSendRepository) GetByRecipient(ctx context.Context, recipient string, limit int) ([]*domain.EmailSend, error) {
	query := `
		SELECT id, recipient, subject, template_name, template_data, body, sent_at, status, error_message
		FROM email_sends
		WHERE recipient = $1
		ORDER BY sent_at DESC
		LIMIT $2
	`

	rows, err := r.DB.Query(ctx, query, recipient, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get email send records by recipient: %w", err)
	}
	defer rows.Close()

	var emailSends []*domain.EmailSend
	for rows.Next() {
		emailSend := &domain.EmailSend{}
		var templateDataJSON []byte
		err := rows.Scan(
			&emailSend.ID,
			&emailSend.Recipient,
			&emailSend.Subject,
			&emailSend.TemplateName,
			&templateDataJSON,
			&emailSend.Body,
			&emailSend.SentAt,
			&emailSend.Status,
			&emailSend.ErrorMessage,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan email send record: %w", err)
		}

		if err := json.Unmarshal(templateDataJSON, &emailSend.TemplateData); err != nil {
			return nil, fmt.Errorf("failed to unmarshal template data: %w", err)
		}
		emailSends = append(emailSends, emailSend)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating through email send records: %w", err)
	}

	return emailSends, nil
}
