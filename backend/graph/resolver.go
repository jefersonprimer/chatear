package graph

import (
	"github.com/jefersonprimer/chatear/backend/application/usecases"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	notificationApplication "github.com/jefersonprimer/chatear/backend/internal/notification/application"
	userApplication "github.com/jefersonprimer/chatear/backend/internal/user/application"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	RegisterUserUseCase    *userApplication.RegisterUser
	LoginUseCase           *userApplication.Login
	VerifyEmailUseCase     *userApplication.VerifyEmail
	LogoutUser             *userApplication.LogoutUser
	ResetPassword        *userApplication.PasswordReset
	DeleteUser             *userApplication.DeleteUser
	RecoverAccount         *userApplication.RecoverAccount
	RefreshToken           *userApplication.RefreshToken
	GetUsersUseCase        usecases.UserUseCases
	TokenService           services.TokenService
	OneTimeTokenService    services.OneTimeTokenService
	EmailRateLimiter       notificationApplication.RateLimiter
	EventBus               repositories.EventBus
	UserDeletionRepository repositories.UserDeletionRepository
	UserRepository         repositories.UserRepository
	AvatarUsecases         *usecases.AvatarUsecases
}

