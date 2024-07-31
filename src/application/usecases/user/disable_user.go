package user

import (
	"errors"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"

	"github.com/google/uuid"
)

type DisableUserUseCase struct {
	userRepo    repositories.UserRepository
	hashService services.HashService
}

func NewDisableUserUseCase(userRepo repositories.UserRepository, hashService services.HashService) DisableUserUseCase {
	return DisableUserUseCase{
		userRepo:    userRepo,
		hashService: hashService,
	}
}

func (uc *DisableUserUseCase) Execute(id uuid.UUID, password string) (bool, error) {
	repo := uc.userRepo
	hashService := uc.hashService

	user, err := repo.GetUserById(id)
	if err != nil {
		return false, err
	}
	// Confirm with password
	if err := hashService.Compare(user.Password, password); err != nil {
		return false, errors.New("invalid credentials")
	}
	// Change status to inactive
	user.Active = false
	userDB, err := repo.Update(id, *user.UpdateUser())
	if err != nil {
		return false, err
	}

	return !userDB.Active, nil
}
