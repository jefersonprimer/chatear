package graph

import (
	"time"

	"github.com/jefersonprimer/chatear/backend/domain/entities"
	"github.com/jefersonprimer/chatear/backend/graph/model"
)

// timePtrToStringPtr converts a *time.Time to a *string, handling nil.
func timePtrToStringPtr(t *time.Time) *string {
	if t == nil {
		return nil
	}
	s := t.String()
	return &s
}

func toModelUser(user *entities.User) *model.User {
	return &model.User{
		ID:                user.ID.String(),
		Name:              user.Name,
		Email:             user.Email,
		CreatedAt:         user.CreatedAt.String(),
		UpdatedAt:         user.UpdatedAt.String(),
		IsEmailVerified:   user.IsEmailVerified,
		DeletedAt:         timePtrToStringPtr(user.DeletedAt),
		AvatarURL:         user.AvatarURL,
		DeletionDueAt:     timePtrToStringPtr(user.DeletionDueAt),
		LastLoginAt:       timePtrToStringPtr(user.LastLoginAt),
		IsDeleted:         user.IsDeleted,
		Gender:            (*model.Gender)(user.Gender),
	}
}