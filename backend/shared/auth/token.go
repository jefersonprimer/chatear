package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jefersonprimer/chatear/backend/config"
	"github.com/jefersonprimer/chatear/backend/domain/repositories"
	"github.com/jefersonprimer/chatear/backend/domain/services"
)

// Claims defines the structure of our JWT claims
type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenService struct {
	refreshTokenRepo repositories.RefreshTokenRepository
	jwtSecret        []byte
	cfg              *config.Config
}

// NewTokenService creates a new TokenService
func NewTokenService(refreshTokenRepo repositories.RefreshTokenRepository, jwtSecret string, cfg *config.Config) services.TokenService {
	return &TokenService{
		refreshTokenRepo: refreshTokenRepo,
		jwtSecret:        []byte(jwtSecret),
		cfg:              cfg,
	}
}

// GenerateAccessToken generates a new access token.
func (s *TokenService) GenerateAccessToken(userID string) (string, error) {
	expirationTime := time.Now().Add(s.cfg.AccessTokenTTL)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", fmt.Errorf("failed to sign access token: %w", err)
	}
	return tokenString, nil
}

// VerifyToken verifies an access token and returns the associated user ID.
func (s *TokenService) VerifyToken(ctx context.Context, tokenString string) (uuid.UUID, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return uuid.Nil, fmt.Errorf("failed to parse access token: %w", err)
	}

	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid access token")
	}

	userID, err := uuid.Parse(claims.UserID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid user ID in token: %w", err)
	}

	return userID, nil
}

// GetTokenExpiration extracts the expiration time from an access token.
func (s *TokenService) GetTokenExpiration(tokenString string) (time.Time, error) {
	claims := &Claims{}
	_, _, err := new(jwt.Parser).ParseUnverified(tokenString, claims)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse unverified token: %w", err)
	}

	if claims.ExpiresAt == nil {
		return time.Time{}, fmt.Errorf("token has no expiration claim")
	}

	return claims.ExpiresAt.Time, nil
}

// GenerateRefreshToken generates a new refresh token.
func (s *TokenService) GenerateRefreshToken(userID string) (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes for refresh token: %w", err)
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// GetRefreshTokenTTL returns the refresh token TTL.
func (s *TokenService) GetRefreshTokenTTL() time.Duration {
	return s.cfg.RefreshTokenTTL
}
