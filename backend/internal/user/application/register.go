package application

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/jackc/pgx/v5"
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	notificationApplication "github.com/jefersonprimer/chatear/backend/internal/notification/application"
	"github.com/jefersonprimer/chatear/backend/shared/errors"
	"github.com/jefersonprimer/chatear/backend/shared/events"
)

// RegisterUserRequest represents the request to register a new user.
type RegisterUserRequest struct {
	Name     string
	Email    string
	Password string
}

// RegisterUserResponse represents the response after registering a new user.
type RegisterUserResponse struct {
	UserID string
}

// RegisterUser is the use case for registering a new user.
type RegisterUser struct {
	UserRepository      repositories.UserRepository
	EventBus            repositories.EventBus
	OneTimeTokenService services.OneTimeTokenService
	EmailRateLimiter    notificationApplication.RateLimiter
}

// NewRegisterUser creates a new RegisterUser use case.
func NewRegisterUser(
	userRepo repositories.UserRepository,
	eventBus repositories.EventBus,
	oneTimeTokenService services.OneTimeTokenService,
	emailRateLimiter notificationApplication.RateLimiter,
) *RegisterUser {
	return &RegisterUser{
		UserRepository:      userRepo,
		EventBus:            eventBus,
		OneTimeTokenService: oneTimeTokenService,
		EmailRateLimiter:    emailRateLimiter,
	}
}

// Execute handles the registration of a new user.
func (uc *RegisterUser) Execute(ctx context.Context, req RegisterUserRequest) (*RegisterUserResponse, error) {
	// Check if user with the given email already exists
	existingUser, err := uc.UserRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("failed to check for existing user: %w", err)
	}
	if existingUser != nil {
		return nil, errors.ErrUserAlreadyExists
	}

	// Check email rate limit
	isAllowed, err := uc.EmailRateLimiter.IsAllowed(ctx, req.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to check email rate limit: %w", err)
	}
	if !isAllowed {
		return nil, errors.ErrTooManyEmailAttempts
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// Create new user entity
	user := entities.NewUser(req.Name, req.Email, string(hashedPassword))

	// Persist the user
	if err := uc.UserRepository.Create(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	// Generate verification token
	verificationToken, err := uc.OneTimeTokenService.GenerateToken(ctx, user.ID.String())
	if err != nil {
		return nil, fmt.Errorf("failed to generate verification token: %w", err)
	}

	// Increment email rate limiter
	if err := uc.EmailRateLimiter.Increment(ctx, req.Email); err != nil {
		// Log the error but don't return it, as user creation was successful
		fmt.Printf("failed to increment email rate limiter for user %s: %v\n", user.ID.String(), err)
	}

	// Publish UserRegisteredEvent
	userRegisteredEvent := events.UserRegisteredEvent{
		UserID:            user.ID.String(),
		Email:             user.Email,
		Name:              user.Name,
		VerificationToken: verificationToken,
		Timestamp:         time.Now(),
	}
	if err := uc.EventBus.Publish(ctx, events.UserRegisteredSubject, userRegisteredEvent); err != nil {
		// Log the error but don't return it, as user creation was successful
		fmt.Printf("failed to publish UserRegisteredEvent for user %s: %v\n", user.ID.String(), err)
	}

	return &RegisterUserResponse{UserID: user.ID.String()}, nil
}
