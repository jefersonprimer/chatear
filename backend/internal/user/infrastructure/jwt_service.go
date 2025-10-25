package infrastructure

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// JWTService is a JWT implementation of the domain.TokenService.
type JWTService struct {
	SecretKey       []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

// NewJWTService creates a new JWTService.
func NewJWTService(cfg *config.Config) services.TokenService {
	return &JWTService{
		SecretKey:       []byte(cfg.JwtSecret),
		AccessTokenTTL:  cfg.AccessTokenTTL,
		RefreshTokenTTL: cfg.RefreshTokenTTL,
	}
}

// CreateAccessToken creates a new access token for the given user.
func (s *JWTService) GenerateAccessToken(userID string) (string, error) {
	claims := &jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.AccessTokenTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.SecretKey)
}

// CreateRefreshToken creates a new refresh token for the given user.
func (s *JWTService) GenerateRefreshToken(userID string) (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

// VerifyToken verifies the given token and returns the user ID.
func (s *JWTService) VerifyToken(ctx context.Context, tokenString string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.SecretKey, nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	return uuid.Parse(claims.Subject)
}

// GetRefreshTokenTTL returns the refresh token TTL.
func (s *JWTService) GetRefreshTokenTTL() time.Duration {
	return s.RefreshTokenTTL
}

// GetTokenExpiration parses the token and returns its expiration time.
func (s *JWTService) GetTokenExpiration(tokenString string) (time.Time, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return s.SecretKey, nil
	})
	if err != nil {
		return time.Time{}, err
	}

	if !token.Valid {
		return time.Time{}, fmt.Errorf("invalid token")
	}

	return claims.ExpiresAt.Time, nil
}
