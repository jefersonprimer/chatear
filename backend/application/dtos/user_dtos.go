package dtos

// RegisterUserInput is the input for the register mutation.
type RegisterUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginInput is the input for the login mutation.
type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LogoutInput is the input for the logout mutation.
type LogoutInput struct {
	Token string `json:"token"`
}

// PasswordRecoveryInput is the input for the password recovery mutation.
type PasswordRecoveryInput struct {
	Email string `json:"email"`
}

// VerifyTokenAndResetPasswordInput is the input for the verify token and reset password mutation.
type VerifyTokenAndResetPasswordInput struct {
	Token       string `json:"token"`
	NewPassword string `json:"newPassword"`
}