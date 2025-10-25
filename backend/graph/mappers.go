package graph

import (
	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/graph/model"
)

func toModelUser(user *entities.User) *model.User {
	modelUser := &model.User{
		ID:              user.ID.String(),
		Name:            user.Name,
		Email:           user.Email,
		CreatedAt:       user.CreatedAt.String(),
		UpdatedAt:       user.UpdatedAt.String(),
		IsEmailVerified: user.IsEmailVerified,
		IsDeleted:       user.IsDeleted,
	}

	if user.DeletedAt != nil {
		deletedAtStr := user.DeletedAt.String()
		modelUser.DeletedAt = &deletedAtStr
	}
	if user.AvatarURL != nil {
		modelUser.AvatarURL = user.AvatarURL
	}
	if user.DeletionDueAt != nil {
		deletionDueAtStr := user.DeletionDueAt.String()
		modelUser.DeletionDueAt = &deletionDueAtStr
	}
	if user.LastLoginAt != nil {
		lastLoginAtStr := user.LastLoginAt.String()
		modelUser.LastLoginAt = &lastLoginAtStr
	}

	return modelUser
}
