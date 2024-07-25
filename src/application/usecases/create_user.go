package usecases

import (
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"

	"golang.org/x/crypto/bcrypt"
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

func (uc *CreateUserUseCase) Execute(user *entity.UserEnt) (*entity.UserEnt, string, error) {
	repo := uc.userRepo

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", err
	}
	user.Password = string(hashedPassword)

	// Save user
	newUser, err := repo.Save(*user.NewUser())
	if err != nil {
		return nil, "", err
	}

	jwtUser := services.AuthUser{
		ID:       newUser.ID,
		Username: newUser.Username,
	}

	tokenPairs, err := uc.auth.GenerateTokenPair(&jwtUser)
	if err != nil {
		return nil, "", err
	}

	return newUser, tokenPairs.AccessToken, nil
}
