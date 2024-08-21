package user

import (
	"errors"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"

	"github.com/google/uuid"
)

type SoftDeleteUserUseCase struct {
	userRepo    repositories.UserRepository
	hashService services.HashService
}

func NewSoftDeleteUserUseCase(userRepo repositories.UserRepository, hashService services.HashService) SoftDeleteUserUseCase {
	return SoftDeleteUserUseCase{
		userRepo:    userRepo,
		hashService: hashService,
	}
}

func (uc *SoftDeleteUserUseCase) Execute(id uuid.UUID, password string) error {
	repo := uc.userRepo
	hashService := uc.hashService

	user, err := repo.GetUserById(id)
	if err != nil {
		return err
	}

	if !user.Active {
		return errors.New("invalid credentials")
	}
	// Confirm with password
	if err := hashService.Compare(user.Password, password); err != nil {
		return errors.New("invalid credentials")
	}
	err = repo.SoftDelete(user.ID)
	if err != nil {
		return err
	}

	return nil
}
