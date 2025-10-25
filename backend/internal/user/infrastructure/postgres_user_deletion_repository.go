package infrastructure

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// PostgresUserDeletionRepository is a PostgreSQL implementation of the UserDeletionRepository.
type PostgresUserDeletionRepository struct {
	db *pgxpool.Pool
}

// NewPostgresUserDeletionRepository creates a new PostgresUserDeletionRepository.
func NewPostgresUserDeletionRepository(db *pgxpool.Pool) repositories.UserDeletionRepository {
	return &PostgresUserDeletionRepository{
		db: db,
	}
}

// Create creates a new user deletion record.
func (r *PostgresUserDeletionRepository) Create(ctx context.Context, userDeletion *entities.UserDeletion) error {
	query := `INSERT INTO user_deletions (id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := r.db.Exec(ctx, query, userDeletion.ID, userDeletion.UserID, userDeletion.ScheduledDate, userDeletion.Status, userDeletion.RecoveryToken, userDeletion.RecoveryTokenExpiresAt)
	return err
}

// GetByID retrieves a user deletion record by its ID.
func (r *PostgresUserDeletionRepository) GetByID(ctx context.Context, id uuid.UUID) (*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE id = $1`
	deletion := &entities.UserDeletion{}
	err := r.db.QueryRow(ctx, query, id).Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
	if err != nil {
		return nil, err
	}
	return deletion, nil
}

// GetByUserID retrieves a user deletion record by user ID.
func (r *PostgresUserDeletionRepository) GetByUserID(ctx context.Context, userID uuid.UUID) (*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE user_id = $1`
	deletion := &entities.UserDeletion{}
	err := r.db.QueryRow(ctx, query, userID).Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
	if err != nil {
		return nil, err
	}
	return deletion, nil
}

// GetByToken retrieves a user deletion record by token.
func (r *PostgresUserDeletionRepository) GetByToken(ctx context.Context, token string) (*entities.UserDeletion, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetByRecoveryToken retrieves a user deletion record by recovery token.
func (r *PostgresUserDeletionRepository) GetByRecoveryToken(ctx context.Context, recoveryToken string) (*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE recovery_token = $1`
	deletion := &entities.UserDeletion{}
	err := r.db.QueryRow(ctx, query, recoveryToken).Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
	if err != nil {
		return nil, err
	}
	return deletion, nil
}

// CancelByUserID cancels a user deletion record by user ID.
func (r *PostgresUserDeletionRepository) CancelByUserID(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE user_deletions SET status = $1 WHERE user_id = $2`
	_, err := r.db.Exec(ctx, query, entities.UserDeletionStatusCancelled, userID)
	return err
}

// Delete deletes a user deletion record by its ID.
func (r *PostgresUserDeletionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM user_deletions WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}

// Update updates a user deletion record.
func (r *PostgresUserDeletionRepository) Update(ctx context.Context, userDeletion *entities.UserDeletion) error {
	query := `UPDATE user_deletions SET user_id = $1, scheduled_date = $2, status = $3, recovery_token = $4, recovery_token_expires_at = $5 WHERE id = $6`
	_, err := r.db.Exec(ctx, query, userDeletion.UserID, userDeletion.ScheduledDate, userDeletion.Status, userDeletion.RecoveryToken, userDeletion.RecoveryTokenExpiresAt, userDeletion.ID)
	return err
}

// GetScheduledDeletions retrieves user deletion records scheduled for a specific date.
func (r *PostgresUserDeletionRepository) GetScheduledDeletions(ctx context.Context, scheduledDate time.Time) ([]*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE scheduled_date = $1`
	rows, err := r.db.Query(ctx, query, scheduledDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deletions []*entities.UserDeletion
	for rows.Next() {
		deletion := &entities.UserDeletion{}
		err := rows.Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
		if err != nil {
			return nil, err
		}
		deletions = append(deletions, deletion)
	}

	return deletions, nil
}

func (r *PostgresUserDeletionRepository) FindScheduledBefore(ctx context.Context, date time.Time) ([]*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE scheduled_date <= $1 AND status = $2`
	rows, err := r.db.Query(ctx, query, date, entities.UserDeletionStatusScheduled)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deletions []*entities.UserDeletion
	for rows.Next() {
		deletion := &entities.UserDeletion{}
		err := rows.Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
		if err != nil {
			return nil, err
		}
		deletions = append(deletions, deletion)
	}

	return deletions, nil
}

// GetByStatus retrieves user deletion records by status.
func (r *PostgresUserDeletionRepository) GetByStatus(ctx context.Context, status entities.UserDeletionStatus, limit, offset int) ([]*entities.UserDeletion, error) {
	query := `SELECT id, user_id, scheduled_date, status, recovery_token, recovery_token_expires_at FROM user_deletions WHERE status = $1 LIMIT $2 OFFSET $3`
	rows, err := r.db.Query(ctx, query, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var deletions []*entities.UserDeletion
	for rows.Next() {
		deletion := &entities.UserDeletion{}
		err := rows.Scan(&deletion.ID, &deletion.UserID, &deletion.ScheduledDate, &deletion.Status, &deletion.RecoveryToken, &deletion.RecoveryTokenExpiresAt)
		if err != nil {
			return nil, err
		}
		deletions = append(deletions, deletion)
	}

	return deletions, nil
}
