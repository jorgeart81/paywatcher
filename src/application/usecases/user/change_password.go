package user

import (
	"errors"
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"
)

type ChangePasswordUseCase struct {
	userRepo    repositories.UserRepository
	hashService services.HashService
}

func NewChangePasswordUseCase(userRepo repositories.UserRepository, hashService services.HashService) LoginUserUseCase {
	return LoginUserUseCase{
		userRepo:    userRepo,
		hashService: hashService,
	}
}

func (uc *ChangePasswordUseCase) Execute(email, oldPassword, newPassword string) (*entity.UserEnt, error) {
	repo := uc.userRepo
	hashService := uc.hashService

	user, err := repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := hashService.Compare(user.Password, oldPassword); err != nil {
		return nil, errors.New("invalid credentials")
	}

	hashPass, err := hashService.Has(newPassword)
	if err != nil {
		return nil, err
	}

	user.Password = hashPass
	updatedUser, err := repo.Update(*user.UpdateUser())
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
