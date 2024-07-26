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
	ID       string
	Audience string
	Issuer   string
	IssuedAt int64
	Expires  int64
}

type Authenticator interface {
	GenerateTokenPair(user *AuthUser) (TokenPairs, error)
	GetTokenFromHeaderAndVerify(w http.ResponseWriter, r *http.Request) (*Claims, error)
	GetRefreshCookie(refreshToken string) *http.Cookie
	GetExpiredRefreshCookie() *http.Cookie
}
