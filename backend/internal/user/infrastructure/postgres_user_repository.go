package infrastructure

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
)

// PostgresUserRepository is a PostgreSQL implementation of the UserRepository.

type PostgresUserRepository struct {
	DB *pgxpool.Pool
}

// NewPostgresUserRepository creates a new PostgresUserRepository.
func NewPostgresUserRepository(db *pgxpool.Pool) *PostgresUserRepository {
	return &PostgresUserRepository{DB: db}
}

// Create creates a new user in the database.
func (r *PostgresUserRepository) Create(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO users (id, name, email, password_hash, is_email_verified, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.DB.Exec(ctx, query, user.ID, user.Name, user.Email, user.PasswordHash, user.IsEmailVerified, user.CreatedAt, user.UpdatedAt)
	return err
}

// FindByID retrieves a user by their ID from the database.
func (r *PostgresUserRepository) FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	query := `SELECT id, name, email, password_hash, is_email_verified, created_at, updated_at, last_login_at, avatar_url, is_deleted, deleted_at, deletion_due_at FROM users WHERE id = $1`
	user := &entities.User{}
	err := r.DB.QueryRow(ctx, query, id).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.IsEmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt, &user.AvatarURL, &user.IsDeleted, &user.DeletedAt, &user.DeletionDueAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// FindByEmail retrieves a user by their email from the database.
func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := `SELECT id, name, email, password_hash, is_email_verified, created_at, updated_at, last_login_at, avatar_url, is_deleted, deleted_at, deletion_due_at FROM users WHERE email = $1`
	user := &entities.User{}
	err := r.DB.QueryRow(ctx, query, email).Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.IsEmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt, &user.AvatarURL, &user.IsDeleted, &user.DeletedAt, &user.DeletionDueAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Update updates a user in the database.
func (r *PostgresUserRepository) Update(ctx context.Context, user *entities.User) error {
	query := `UPDATE users SET name = $1, email = $2, password_hash = $3, is_email_verified = $4, is_deleted = $5, deleted_at = $6, created_at = $7, updated_at = $8, last_login_at = $9, avatar_url = $10, deletion_due_at = $11 WHERE id = $12`
	_, err := r.DB.Exec(ctx, query, user.Name, user.Email, user.PasswordHash, user.IsEmailVerified, user.IsDeleted, user.DeletedAt, user.CreatedAt, user.UpdatedAt, user.LastLoginAt, user.AvatarURL, user.DeletionDueAt, user.ID)
	return err
}

// Delete soft deletes a user in the database.
func (r *PostgresUserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `UPDATE users SET is_deleted = true, deleted_at = $1 WHERE id = $2`
	_, err := r.DB.Exec(ctx, query, time.Now(), id)
	return err
}

// FindSoftDeletedBefore retrieves all soft-deleted users from the database before a given time.
func (r *PostgresUserRepository) FindSoftDeletedBefore(ctx context.Context, t time.Time) ([]*entities.User, error) {
	query := `SELECT id, name, email, password_hash, is_email_verified, created_at, updated_at, last_login_at, avatar_url, is_deleted, deleted_at, deletion_due_at FROM users WHERE is_deleted = true AND deleted_at < $1`
	rows, err := r.DB.Query(ctx, query, t)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.IsEmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt, &user.AvatarURL, &user.IsDeleted, &user.DeletedAt, &user.DeletionDueAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// HardDelete permanently deletes a user from the database.
func (r *PostgresUserRepository) HardDelete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(ctx, query, id)
	return err
}

// FindAll retrieves all users from the database.
func (r *PostgresUserRepository) FindAll(ctx context.Context) ([]*entities.User, error) {
	query := `SELECT id, name, email, password_hash, is_email_verified, created_at, updated_at, last_login_at, avatar_url, is_deleted, deleted_at, deletion_due_at FROM users`
	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User
	for rows.Next() {
		user := &entities.User{}
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.PasswordHash, &user.IsEmailVerified, &user.CreatedAt, &user.UpdatedAt, &user.LastLoginAt, &user.AvatarURL, &user.IsDeleted, &user.DeletedAt, &user.DeletionDueAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
