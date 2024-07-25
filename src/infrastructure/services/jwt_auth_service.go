package services

import (
	"paywatcher/src/domain/services"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTAuth struct {
	JWTIssuer     string
	JWTAudience   string
	JWTSecret     string
	JWTExpiry     time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

func (a *JWTAuth) GenerateTokenPair(user *services.AuthUser) (services.TokenPairs, error) {
	// Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"sub":      user.ID,
		"aud":      a.JWTAudience,
		"iss":      a.JWTIssuer,
		"iat":      time.Now().UTC().Unix(),
		"type":     "JWT",
		"exp":      time.Now().UTC().Add(a.JWTExpiry).Unix(),
	})

	// Create a signed token
	signedAccessToken, err := token.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create a refresh token and set claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"iat":  time.Now().UTC().Unix(),
		"type": "JWT",
		"exp":  time.Now().UTC().Add(a.RefreshExpiry).Unix(),
	})

	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create services.TokenPairs and populate with signed tokens
	var tokenPairs = services.TokenPairs{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// Return services.TokenPairs
	return tokenPairs, nil
}
