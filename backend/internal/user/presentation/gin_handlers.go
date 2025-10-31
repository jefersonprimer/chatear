package presentation

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
	"github.com/jefersonprimer/chatear/backend/internal/user/application"
	appErrors "github.com/jefersonprimer/chatear/backend/shared/errors"
	"github.com/jefersonprimer/chatear/backend/shared/auth"
)

// UserHandler holds the dependencies for user-related HTTP handlers.
type UserHandler struct {
	RegisterUser                *application.RegisterUser
	Login                       *application.Login
	VerifyEmail                 *application.VerifyEmail
	LogoutUser                  *application.LogoutUser
	PasswordReset            *application.PasswordReset
	VerifyTokenAndResetPassword *application.VerifyTokenAndResetPassword
	RecoverAccount              *application.RecoverAccount
	DeleteUser                  *application.DeleteUser
	RefreshToken                *application.RefreshToken
	OneTimeTokenService         services.OneTimeTokenService
	TokenService                services.TokenService
	BlacklistRepository         repositories.BlacklistRepository
	FrontendURL                 string
}

// NewUserHandlers initializes and registers user-related routes.
func NewUserHandlers(
	router *gin.RouterGroup,
	registerUser *application.RegisterUser,
	login *application.Login,
	verifyEmail *application.VerifyEmail,
	logoutUser *application.LogoutUser,
	passwordReset *application.PasswordReset,
	verifyTokenAndResetPassword *application.VerifyTokenAndResetPassword,
	recoverAccount *application.RecoverAccount,
	deleteUser *application.DeleteUser,
	refreshToken *application.RefreshToken,
	oneTimeTokenService services.OneTimeTokenService,
	tokenService services.TokenService,
	blacklistRepo repositories.BlacklistRepository,
	frontendURL string,
) {
	handler := &UserHandler{
		RegisterUser:                registerUser,
		Login:                       login,
		VerifyEmail:                 verifyEmail,
		LogoutUser:                  logoutUser,
		PasswordReset:            passwordReset,
		VerifyTokenAndResetPassword: verifyTokenAndResetPassword,
		RecoverAccount:              recoverAccount,
		DeleteUser:                  deleteUser,
		RefreshToken:                refreshToken,
		OneTimeTokenService:         oneTimeTokenService,
		TokenService:                tokenService,
		BlacklistRepository:         blacklistRepo,
		FrontendURL:                 frontendURL,
	}

	// Public routes
	router.POST("/register", handler.Register)
	router.POST("/login", handler.LoginHandler)
	router.GET("/verify-email", handler.VerifyEmailHandler)
	router.POST("/request-password-reset", handler.ResetPasswordHandler)
	router.GET("/password-reset-token", handler.HandlePasswordResetTokenRedirect) // Handles password reset token validation and redirect
	router.POST("/reset-password-confirm", handler.ResetPasswordConfirmHandler)
	router.POST("/recover-account", handler.RecoverAccountHandler)
	router.POST("/refresh-token", handler.RefreshTokenHandler)

	// Authenticated routes
	authenticated := router.Group("/")
	authenticated.Use(auth.AuthMiddleware(tokenService, blacklistRepo))
	{
		authenticated.POST("/logout", handler.Logout)
		authenticated.DELETE("/delete-account", handler.DeleteAccount)
	}
}

// RefreshTokenRequest represents the request to refresh a token.
type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

// ResetPasswordConfirmRequest represents the request to confirm password reset.
type ResetPasswordConfirmRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}

// ResetPasswordConfirmHandler handles the POST request to confirm password reset.
func (h *UserHandler) ResetPasswordConfirmHandler(c *gin.Context) {
	var req ResetPasswordConfirmRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.VerifyTokenAndResetPassword.Execute(c.Request.Context(), req.Token, req.NewPassword)
	if err != nil {
		if errors.Is(err, appErrors.ErrInvalidToken) || errors.Is(err, appErrors.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		if errors.Is(err, appErrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to reset password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// HandlePasswordResetTokenRedirect handles the GET request for password reset token validation and redirection.
func (h *UserHandler) HandlePasswordResetTokenRedirect(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		errorURL := fmt.Sprintf("%s/auth/reset-password?error=token_missing", h.FrontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Validate the token
	_, err := h.OneTimeTokenService.PeekToken(c.Request.Context(), token)
	if err != nil {
		errorURL := fmt.Sprintf("%s/auth/reset-password?error=invalid_token", h.FrontendURL)
		c.Redirect(http.StatusFound, errorURL)
		return
	}

	// Redirect to the frontend password reset page
	resetPasswordURL := fmt.Sprintf("%s/auth/reset-password?token=%s", h.FrontendURL, token)
	c.Redirect(http.StatusFound, resetPasswordURL)
}

// Register handles user registration.
func (h *UserHandler) Register(c *gin.Context) {
	var req application.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.RegisterUser.Execute(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, appErrors.ErrUserAlreadyExists) {
			c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
			return
		}
		if errors.Is(err, appErrors.ErrTooManyEmailAttempts) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many email attempts, please try again later"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully. Please verify your email."})
}

// LoginHandler handles user login.
func (h *UserHandler) LoginHandler(c *gin.Context) {
	var req application.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.Login.Execute(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, appErrors.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// VerifyEmailHandler handles email verification.
func (h *UserHandler) VerifyEmailHandler(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.Redirect(http.StatusFound, fmt.Sprintf("%s/auth/verify-email?error=token_missing", h.FrontendURL))
		return
	}

	verifyReq := application.VerifyEmailRequest{
		Token: token,
	}

	_, err := h.VerifyEmail.Execute(c.Request.Context(), verifyReq)
	if err != nil {
		if errors.Is(err, appErrors.ErrInvalidToken) || errors.Is(err, appErrors.ErrTokenExpired) {
			c.Redirect(http.StatusFound, fmt.Sprintf("%s/auth/verify-email?error=invalid_token", h.FrontendURL))
			return
		}
		c.Redirect(http.StatusFound, fmt.Sprintf("%s/auth/verify-email?error=verification_failed", h.FrontendURL))
		return
	}

	c.Redirect(http.StatusFound, fmt.Sprintf("%s/auth/sucess", h.FrontendURL))
}

// Logout handles user logout.
func (h *UserHandler) Logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token is missing"})
		return
	}

	err := h.LogoutUser.Execute(c.Request.Context(), token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to logout"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Logged out successfully"})
}

// ResetPasswordHandler handles password reset requests.
func (h *UserHandler) ResetPasswordHandler(c *gin.Context) {
	var req application.PasswordResetRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.PasswordReset.Execute(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, appErrors.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return	
		}
		if errors.Is(err, appErrors.ErrTooManyEmailAttempts) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many password reset attempts, please try again later"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send password reset email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent if user exists"})
}

// DeleteAccount handles user account deletion requests.
func (h *UserHandler) DeleteAccount(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := h.DeleteUser.Execute(c.Request.Context(), application.DeleteUserRequest{UserID: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initiate account deletion"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account deletion initiated. Please check your email for confirmation."})
}

// RecoverAccountHandler handles account recovery requests.
func (h *UserHandler) RecoverAccountHandler(c *gin.Context) {
	var req application.RecoverAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.RecoverAccount.Execute(c.Request.Context(), req)
	if err != nil {
		if errors.Is(err, appErrors.ErrInvalidToken) || errors.Is(err, appErrors.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to recover account"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account recovered successfully"})
}

// RefreshTokenHandler handles refresh token requests.
func (h *UserHandler) RefreshTokenHandler(c *gin.Context) {
	var req RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.RefreshToken.Execute(c.Request.Context(), req.RefreshToken)
	if err != nil {
		if errors.Is(err, appErrors.ErrInvalidToken) || errors.Is(err, appErrors.ErrTokenExpired) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired refresh token"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
