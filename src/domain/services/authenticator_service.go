package services

import (
	"net/http"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type TokenPairs struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	Username string
	ID       uuid.UUID
}

type Authenticator interface {
	GenerateTokenPair(user *AuthUser) (TokenPairs, error)
	VerifyToken(token string) (*Claims, error)
	GetRefreshCookie(refreshToken string) *http.Cookie
	GetExpiredRefreshCookie() *http.Cookie
}
