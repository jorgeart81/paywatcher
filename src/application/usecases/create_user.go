package usecases

import (
	"paywatcher/src/domain/userdomain"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase struct {
	userRepo userdomain.UserRepository
}

func NewCreateUserUseCase(userRepo userdomain.UserRepository) CreateUserUseCase {
	return CreateUserUseCase{
		userRepo: userRepo,
	}
}

func (uc *CreateUserUseCase) Execute(user userdomain.User) (*userdomain.User, error) {
	repo := uc.userRepo

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	// Save user
	newUser, err := repo.Save(*user.NewUser())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
