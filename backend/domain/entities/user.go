package entities

import (
	"time"

	"github.com/google/uuid"
)

// User represents a user entity in the domain
type User struct {
	ID                uuid.UUID  `json:"id"`
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	PasswordHash      string     `json:"-"`
	CreatedAt         time.Time  `json:"created_at"`
	UpdatedAt         time.Time  `json:"updated_at"`
	IsEmailVerified   bool       `json:"is_email_verified"`
	DeletedAt         *time.Time `json:"deleted_at,omitempty"`
	AvatarURL         *string    `json:"avatar_url,omitempty"`
	AvatarPublicID    *string    `json:"avatar_public_id,omitempty"`
	DeletionDueAt     *time.Time `json:"deletion_due_at,omitempty"`
	LastLoginAt       *time.Time `json:"last_login_at,omitempty"`
	IsDeleted         bool       `json:"is_deleted"`
	Gender            *string    `json:"gender,omitempty"`
}

// NewUser creates a new user entity
func NewUser(name, email, passwordHash, gender string) *User {
	return &User{
		ID:              uuid.New(),
		Name:            name,
		Email:           email,
		PasswordHash:    passwordHash,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		IsEmailVerified: false,
		IsDeleted:       false,
		Gender:          &gender,
	}
}

// Validate validates the user entity
func (u *User) Validate() error {
	// Add validation logic here
	return nil
}

// MarkAsDeleted marks the user as deleted
func (u *User) MarkAsDeleted() {
	now := time.Now()
	u.IsDeleted = true
	u.DeletedAt = &now
	u.UpdatedAt = now
}

// Restore restores a soft-deleted user
func (u *User) Restore() {
	u.IsDeleted = false
	u.DeletedAt = nil
	u.UpdatedAt = time.Now()
}

// UpdateLastLogin updates the last login timestamp
func (u *User) UpdateLastLogin() {
	now := time.Now()
	u.LastLoginAt = &now
	u.UpdatedAt = now
}

// VerifyEmail marks the user's email as verified
func (u *User) VerifyEmail() {
	u.IsEmailVerified = true
	u.UpdatedAt = time.Now()
}