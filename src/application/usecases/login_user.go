package usecases

import (
	"errors"
	"paywatcher/src/application/auth"
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"

	"golang.org/x/crypto/bcrypt"
)

type LoginUserUseCase struct {
	userRepo   repositories.UserRepository
	auth       *auth.Auth
	hasService services.HashService
}

func NewLoginUserUseCase(userRepo repositories.UserRepository, auth *auth.Auth, hasService services.HashService) LoginUserUseCase {
	return LoginUserUseCase{
		userRepo:   userRepo,
		auth:       auth,
		hasService: hasService,
	}
}

func (uc *LoginUserUseCase) Execute(email, password string) (*entity.UserEnt, string, error) {
	repo := uc.userRepo

	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	jwtUser := auth.JwtUser{
		ID:       user.ID,
		Username: user.Username,
	}

	tokenPairs, err := uc.auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return nil, "", err
	}

	return user, tokenPairs.AccessToken, nil
}
