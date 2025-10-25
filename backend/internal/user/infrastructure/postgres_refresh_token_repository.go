package infrastructure

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// PostgresRefreshTokenRepository is a PostgreSQL implementation of the RefreshTokenRepository.
type PostgresRefreshTokenRepository struct {
	db *pgxpool.Pool
}

// NewPostgresRefreshTokenRepository creates a new PostgresRefreshTokenRepository.
func NewPostgresRefreshTokenRepository(db *pgxpool.Pool) repositories.RefreshTokenRepository {
	return &PostgresRefreshTokenRepository{
		db: db,
	}
}

// CreateRefreshToken creates a new refresh token in the database.
func (r *PostgresRefreshTokenRepository) CreateRefreshToken(ctx context.Context, refreshToken *entities.RefreshToken) error {
	query := `INSERT INTO refresh_tokens (id, user_id, token, expires_at) VALUES ($1, $2, $3, $4)`
	_, err := r.db.Exec(ctx, query, refreshToken.ID, refreshToken.UserID, refreshToken.Token, refreshToken.ExpiresAt)
	return err
}

// GetRefreshTokensByUserID retrieves all refresh tokens for a user.
func (r *PostgresRefreshTokenRepository) GetRefreshTokensByUserID(ctx context.Context, userID uuid.UUID) ([]*entities.RefreshToken, error) {
	query := `SELECT id, user_id, token, expires_at, revoked FROM refresh_tokens WHERE user_id = $1`
	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []*entities.RefreshToken
	for rows.Next() {
		token := &entities.RefreshToken{}
		err := rows.Scan(&token.ID, &token.UserID, &token.Token, &token.ExpiresAt, &token.Revoked)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}

// GetByToken retrieves a refresh token by its token string.
func (r *PostgresRefreshTokenRepository) GetByToken(ctx context.Context, token string) (*entities.RefreshToken, error) {
	query := `SELECT id, user_id, token, expires_at, revoked FROM refresh_tokens WHERE token = $1`
	refreshToken := &entities.RefreshToken{}
	err := r.db.QueryRow(ctx, query, token).Scan(&refreshToken.ID, &refreshToken.UserID, &refreshToken.Token, &refreshToken.ExpiresAt, &refreshToken.Revoked)
	if err != nil {
		return nil, err
	}
	return refreshToken, nil
}

// RevokeAllUserTokens revokes all refresh tokens for a user.
func (r *PostgresRefreshTokenRepository) RevokeAllUserTokens(ctx context.Context, userID uuid.UUID) error {
	query := `UPDATE refresh_tokens SET revoked = true WHERE user_id = $1`
	_, err := r.db.Exec(ctx, query, userID)
	return err
}
