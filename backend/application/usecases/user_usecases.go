package usecases

import (
	"context"

	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

// UserUseCases defines the interface for user use cases
type UserUseCases interface {
	GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetUsers(ctx context.Context) ([]*entities.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// UserUseCasesImpl implements UserUseCases
type UserUseCasesImpl struct {
	userRepo repositories.UserRepository
}

// NewUserUseCases creates a new user use cases instance
func NewUserUseCases(
	userRepo repositories.UserRepository,
) UserUseCases {
	return &UserUseCasesImpl{
		userRepo: userRepo,
	}
}

// GetUser retrieves a user by ID
func (u *UserUseCasesImpl) GetUser(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	user, err := u.userRepo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUsers retrieves all users
func (u *UserUseCasesImpl) GetUsers(ctx context.Context) ([]*entities.User, error) {
	users, err := u.userRepo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// UpdateUser updates a user
func (u *UserUseCasesImpl) UpdateUser(ctx context.Context, id uuid.UUID, updates map[string]interface{}) error {
	// Implementation would go here
	return nil
}

// DeleteUser deletes a user
func (u *UserUseCasesImpl) DeleteUser(ctx context.Context, id uuid.UUID) error {
	// Implementation would go here
	return nil
}
