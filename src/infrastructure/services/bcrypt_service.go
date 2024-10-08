package services

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

// var _ services.HashService = &BcryptService{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

// Compare implements services.HashService.
func (b *BcryptService) Compare(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Has implements services.HashService.
func (b *BcryptService) Has(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
