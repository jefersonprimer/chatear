package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
)

// UserRepository defines the interface for interacting with user data.
type UserRepository interface {
	Create(ctx context.Context, user *entities.User) error
	FindByID(ctx context.Context, id uuid.UUID) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	FindAll(ctx context.Context) ([]*entities.User, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	FindSoftDeletedBefore(ctx context.Context, t time.Time) ([]*entities.User, error)
	HardDelete(ctx context.Context, id uuid.UUID) error
	UpdateAvatar(ctx context.Context, id uuid.UUID, avatarURL, avatarPublicID string) error
}
