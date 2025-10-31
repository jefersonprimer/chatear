package usecases

import (
	"context"
	"io"

	"github.com/google/uuid"

	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/internal/user/services"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
)

type AvatarUsecases struct {
	userRepo         repositories.UserRepository
	cloudinaryService *services.CloudinaryService
}

func NewAvatarUsecases(userRepo repositories.UserRepository, cloudinaryService *services.CloudinaryService) *AvatarUsecases {
	return &AvatarUsecases{
		userRepo:         userRepo,
		cloudinaryService: cloudinaryService,
	}
}

func (uc *AvatarUsecases) UploadAvatar(ctx context.Context, userID uuid.UUID, file io.Reader) (*entities.User, error) {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	url, publicID, err := uc.cloudinaryService.UploadAvatar(ctx, file, user.ID.String())
	if err != nil {
		return nil, err
	}

	err = uc.userRepo.UpdateAvatar(ctx, user.ID, url, publicID)
	if err != nil {
		return nil, err
	}

	return uc.userRepo.FindByID(ctx, userID)
}

func (uc *AvatarUsecases) DeleteAvatar(ctx context.Context, userID uuid.UUID) error {
	user, err := uc.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.AvatarPublicID == nil {
		return nil
	}

	err = uc.cloudinaryService.DeleteAvatar(ctx, *user.AvatarPublicID)
	if err != nil {
		return err
	}

	return uc.userRepo.UpdateAvatar(ctx, user.ID, "", "")
}
