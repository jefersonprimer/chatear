package auth

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

type contextKey string

const (
	ContextKeyUserID       contextKey = "userID"
	ContextKeyRefreshToken contextKey = "refreshToken"
	ContextKeyAccessToken  contextKey = "accessToken"
)

// AuthMiddleware creates a Gin middleware for JWT authentication.
func AuthMiddleware(tokenService services.TokenService, blacklistRepo repositories.BlacklistRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			return
		}

		tokenString := authHeader
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		
		isBlacklisted, err := blacklistRepo.Check(c.Request.Context(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to check token blacklist"})
			return
		}

		if isBlacklisted {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token has been revoked"})
			return
		}

		userID, err := tokenService.VerifyToken(c.Request.Context(), tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Store userID in Gin context
		c.Set(string(ContextKeyUserID), userID)

		// Extract refresh token from header
		refreshToken := c.GetHeader("X-Refresh-Token")

		// Store userID, accessToken, and refreshToken in request context for GraphQL resolvers
		ctx := context.WithValue(c.Request.Context(), ContextKeyUserID, userID)
		ctx = context.WithValue(ctx, ContextKeyAccessToken, tokenString)
		ctx = context.WithValue(ctx, ContextKeyRefreshToken, refreshToken)
		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

// OptionalAuthMiddleware tries to authenticate the user and add the user ID to the context,
// but does not fail if the user is not authenticated.
func OptionalAuthMiddleware(tokenService services.TokenService, blacklistRepo repositories.BlacklistRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := authHeader
			if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
				tokenString = tokenString[7:]
			}

			isBlacklisted, err := blacklistRepo.Check(c.Request.Context(), tokenString)
			if err != nil {
				// Log the error, but don't block the request
			}

			if !isBlacklisted {
				userID, err := tokenService.VerifyToken(c.Request.Context(), tokenString)
				if err == nil {
					// Store userID in Gin context
					c.Set(string(ContextKeyUserID), userID)

					// Store userID, accessToken, and refreshToken in request context for GraphQL resolvers
					ctx := context.WithValue(c.Request.Context(), ContextKeyUserID, userID)
					ctx = context.WithValue(ctx, ContextKeyAccessToken, tokenString)
					c.Request = c.Request.WithContext(ctx)
				}
			}
		}

		c.Next()
	}
}

// GetUserIDFromContext extracts the UserID from the context.
func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	userID, ok := ctx.Value(ContextKeyUserID).(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("user ID not found in context")
	}
	return userID, nil
}
