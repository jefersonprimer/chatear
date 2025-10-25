package repositories

import (
	"context"
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/entities"
)

// DeletionCapacityRepository is an interface for a deletion capacity repository.
type DeletionCapacityRepository interface {
	GetDeletionCapacity(ctx context.Context, date time.Time) (*entities.DeletionCapacity, error)
	IncrementDeletionCapacity(ctx context.Context, date time.Time) error
}
