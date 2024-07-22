package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type Auth struct {
	JWTIssuer     string
	JWTAudience   string
	JWTSecret     string
	JWTExpiry     time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type JwtUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type TokenPairs struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *Auth) GenerateTokenPair(user *JwtUser) (TokenPairs, error) {
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
		return TokenPairs{}, err
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
		return TokenPairs{}, err
	}

	// Create TokenPairs and populate with signed tokens
	var tokenPairs = TokenPairs{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// Return TokenPairs
	return tokenPairs, nil
}
