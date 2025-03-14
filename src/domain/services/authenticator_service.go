package services

import (
	"github.com/google/uuid"
)

type AuthUser struct {
	ID            uuid.UUID `json:"id"`
	Username      string    `json:"username"`
	SecurityStamp string    `json:"securityStamp"`
}

type TokenPairs struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	Username      string
	ID            uuid.UUID
	SecurityStamp string
}

type Authenticator interface {
	GenerateTokenPair(user *AuthUser) (TokenPairs, error)
	VerifyToken(token string) (*Claims, error)
}
