package usecases

import (
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"
)

type CreateUserUseCase struct {
	userRepo    repositories.UserRepository
	auth        services.Authenticator
	hashService services.HashService
}

func NewCreateUserUseCase(userRepo repositories.UserRepository, auth services.Authenticator, hashService services.HashService) CreateUserUseCase {
	return CreateUserUseCase{
		userRepo:    userRepo,
		auth:        auth,
		hashService: hashService,
	}
}

func (uc *CreateUserUseCase) Execute(user *entity.UserEnt) (*entity.UserEnt, services.TokenPairs, error) {
	repo := uc.userRepo

	// Hash the password
	hashedPassword, err := uc.hashService.Has(user.Password)
	if err != nil {
		return nil, services.TokenPairs{}, err
	}
	user.Password = hashedPassword

	// Save user
	newUser, err := repo.Save(*user.NewUser())
	if err != nil {
		return nil, services.TokenPairs{}, err
	}

	jwtUser := services.AuthUser{
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	tokenPairs, err := uc.auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return nil, services.TokenPairs{}, err
	}

	return newUser, tokenPairs, nil
}
