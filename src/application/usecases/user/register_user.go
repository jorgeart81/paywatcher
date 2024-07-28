package user

import (
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"
)

type RegisterUserUseCase struct {
	userRepo    repositories.UserRepository
	auth        services.Authenticator
	hashService services.HashService
}

func NewRegisterUserUseCase(userRepo repositories.UserRepository, auth services.Authenticator, hashService services.HashService) RegisterUserUseCase {
	return RegisterUserUseCase{
		userRepo:    userRepo,
		auth:        auth,
		hashService: hashService,
	}
}

func (uc *RegisterUserUseCase) Execute(user *entity.UserEnt) (*entity.UserEnt, services.TokenPairs, error) {
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
