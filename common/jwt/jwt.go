package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	audience = "koki-platform-families"
	issuer   = "koki-sa"
	subject  = "api-auth"
)

var (
	errInvalidTokenClaims           = errors.New("invalid token claims")
	errUnexpectedTokenSigningMethod = errors.New("unexpected token signing method")
)

// TokenManager is a JSON web token manager
type TokenManager struct {
	accessTokenSecretKey  string
	refreshTokenSecretKey string
	accessTokenDuration   time.Duration
	refreshTokenDuration  time.Duration
}

// UserForToken is bridge for bussinessUser to authUser
type UserForToken struct {
	Email       string `json:"email"`
	UserID      string `json:"user_id"`
	PhoneNumber string `json:"phonenumber"`
}

// CustomClaims is a custom JWT claims that contains some user's information
type CustomClaims struct {
	jwt.StandardClaims
	UserForToken
}

// NewJWTManager returns a new JWT manager
func NewJWTManager(accessTokenSecretKey, refreshTokenSecretKey string, accessTokenDuration,
	refreshTokenDuration time.Duration) *TokenManager {
	return &TokenManager{accessTokenSecretKey, refreshTokenSecretKey,
		accessTokenDuration, refreshTokenDuration}
}

// GenerateAccessToken generates and signs a new token for a user
func (manager *TokenManager) GenerateAccessToken(user *UserForToken) (string, error) {
	t := time.Now().UTC()
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			Issuer:    issuer,
			Audience:  audience,
			IssuedAt:  t.Unix(),
			ExpiresAt: t.Add(manager.accessTokenDuration).Unix(),
		},
		UserForToken: UserForToken{
			Email:       user.Email,
			UserID:      user.UserID,
			PhoneNumber: user.PhoneNumber,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.accessTokenSecretKey))
}

// GenerateRefreshToken generates and signs a new token for a user
func (manager *TokenManager) GenerateRefreshToken(user *UserForToken) (string, error) {
	t := time.Now().UTC()
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   subject,
			ExpiresAt: t.Add(manager.refreshTokenDuration).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.refreshTokenSecretKey))
}

// Verify verifies the access token string and return a user claim if the token is valid
func (manager *TokenManager) Verify(accessToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errUnexpectedTokenSigningMethod
			}

			return []byte(manager.accessTokenSecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errInvalidTokenClaims
	}

	return claims, nil
}

// VerifyRefresh verifies the refresh token string and return new access token if the token is valid
func (manager *TokenManager) VerifyRefresh(refreshToken string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, errUnexpectedTokenSigningMethod
			}

			return []byte(manager.refreshTokenSecretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("invalid token: %w", err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errInvalidTokenClaims
	}

	return claims, nil
}
