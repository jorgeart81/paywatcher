package services

import "github.com/google/uuid"

type AuthUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type TokenPairs struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Authenticator interface {
	GenerateTokenPair(user *AuthUser) (TokenPairs, error)
}
