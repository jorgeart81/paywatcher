package usecases

import (
	"errors"
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"
)

type LoginUserUseCase struct {
	userRepo    repositories.UserRepository
	auth        services.Authenticator
	hashService services.HashService
}

func NewLoginUserUseCase(userRepo repositories.UserRepository, auth services.Authenticator, hashService services.HashService) LoginUserUseCase {
	return LoginUserUseCase{
		userRepo:    userRepo,
		auth:        auth,
		hashService: hashService,
	}
}

func (uc *LoginUserUseCase) Execute(email, password string) (*entity.UserEnt, services.TokenPairs, error) {
	repo := uc.userRepo

	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, services.TokenPairs{}, err
	}

	if err := uc.hashService.Compare(user.Password, password); err != nil {
		return nil, services.TokenPairs{}, errors.New("invalid credentials")
	}

	jwtUser := services.AuthUser{
		ID:       user.ID,
		Username: user.Username,
	}

	tokenPairs, err := uc.auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return nil, services.TokenPairs{}, err
	}

	return user, tokenPairs, nil
}
