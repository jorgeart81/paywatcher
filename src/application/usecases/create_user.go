package usecases

import (
	"errors"
	"paywatcher/src/domain/userdomain"
)

type CreateUserUseCase struct {
	userRepository userdomain.UserRepository
}

func NewCreateUserUseCase(userRepository userdomain.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		userRepository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(user userdomain.User) (*userdomain.User, error) {
	repo := uc.userRepository

	u, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, err
	} else if u.Email == user.Email {
		return nil, errors.New("user already exists")
	}

	newUser, err := repo.Save(*user.NewUser())
	if err != nil {
		return nil, err
	}

	return newUser, nil
}
