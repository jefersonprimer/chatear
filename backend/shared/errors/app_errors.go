package errors

import "errors"

var (
	ErrNotFound             = errors.New("not found")
	ErrRateLimitExceeded    = errors.New("rate limit exceeded")
	ErrExpired              = errors.New("expired")
	ErrUserAlreadyExists    = errors.New("user with this email already exists")
	ErrTooManyEmailAttempts = errors.New("too many email attempts, please try again later")
	ErrInvalidCredentials   = errors.New("invalid credentials")
	ErrEmailNotVerified     = errors.New("email not verified")
	ErrInvalidToken         = errors.New("invalid token")
	ErrTokenExpired         = errors.New("token expired")
	ErrUserNotFound         = errors.New("user not found")
)
