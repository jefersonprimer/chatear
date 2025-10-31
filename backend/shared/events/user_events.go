package events

import "time"

const (
	UserRegisteredSubject          = "user.registered"
	PasswordResetRequestedSubject = "password.reset.requested"
	AccountDeletionRequestedSubject  = "account.deletion.requested"
)

// UserRegisteredEvent is published when a new user registers
type UserRegisteredEvent struct {
	UserID            string    `json:"userID"`	
	Email             string    `json:"email"`
	Timestamp         time.Time `json:"timestamp"`
	VerificationToken string    `json:"verificationToken"`
	Name              string    `json:"name"`
}

// PasswordResetRequestedEvent is published when a user requests password reset
type PasswordResetRequestedEvent struct {
	UserID            string    `json:"userID"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	VerificationToken string    `json:"verificationToken"`
	Timestamp         time.Time `json:"timestamp"`
	FrontendURL       string    `json:"frontendURL"`
}

// AccountDeletionRequestedEvent is published when a user requests account deletion
type AccountDeletionRequestedEvent struct {
	UserID            string    `json:"userID"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	RecoveryToken     string    `json:"recoveryToken"`
	Timestamp         time.Time `json:"timestamp"`
	FrontendURL       string    `json:"frontendURL"`
}

