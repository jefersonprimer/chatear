package events

import "time"

const (
	UserRegisteredSubject          = "user.registered"
	PasswordRecoveryRequestedSubject = "password.recovery.requested"
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

// PasswordRecoveryRequestedEvent is published when a user requests password recovery
type PasswordRecoveryRequestedEvent struct {
	UserID            string    `json:"userID"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	VerificationToken string    `json:"verificationToken"`
	Timestamp         time.Time `json:"timestamp"`
	AppURL            string    `json:"appURL"`
}

// AccountDeletionRequestedEvent is published when a user requests account deletion
type AccountDeletionRequestedEvent struct {
	UserID            string    `json:"userID"`
	Email             string    `json:"email"`
	Name              string    `json:"name"`
	RecoveryToken     string    `json:"recoveryToken"`
	Timestamp         time.Time `json:"timestamp"`
	AppURL            string    `json:"appURL"`
}

