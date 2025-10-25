package infrastructure

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// PostgresDeletionCapacityRepository is a PostgreSQL implementation of the DeletionCapacityRepository.
type PostgresDeletionCapacityRepository struct {
	db *pgxpool.Pool
}

// NewPostgresDeletionCapacityRepository creates a new PostgresDeletionCapacityRepository.
func NewPostgresDeletionCapacityRepository(db *pgxpool.Pool) repositories.DeletionCapacityRepository {
	return &PostgresDeletionCapacityRepository{
		db: db,
	}
}

// GetDeletionCapacity retrieves the deletion capacity for a given date.
func (r *PostgresDeletionCapacityRepository) GetDeletionCapacity(ctx context.Context, date time.Time) (*entities.DeletionCapacity, error) {
	query := `SELECT date, count, max_limit FROM deletion_capacities WHERE date = $1`
	row := r.db.QueryRow(ctx, query, date)

	capacity := &entities.DeletionCapacity{}
	err := row.Scan(&capacity.Day, &capacity.Count, &capacity.MaxLimit)
	if err != nil {
		return nil, err
	}

	return capacity, nil
}

// IncrementDeletionCapacity increments the deletion capacity for a given date.
func (r *PostgresDeletionCapacityRepository) IncrementDeletionCapacity(ctx context.Context, date time.Time) error {
	query := `INSERT INTO deletion_capacities (date, count, max_limit) VALUES ($1, 1, 10) ON CONFLICT (date) DO UPDATE SET count = deletion_capacities.count + 1`
	_, err := r.db.Exec(ctx, query, date)
	return err
}
