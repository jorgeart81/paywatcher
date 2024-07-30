package user

import (
	"errors"
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"
	"paywatcher/src/domain/services"

	"github.com/google/uuid"
)

type ChangePasswordUseCase struct {
	userRepo    repositories.UserRepository
	auth        services.Authenticator
	hashService services.HashService
}

func NewChangePasswordUseCase(userRepo repositories.UserRepository, auth services.Authenticator, hashService services.HashService) ChangePasswordUseCase {
	return ChangePasswordUseCase{
		userRepo:    userRepo,
		hashService: hashService,
		auth:        auth,
	}
}

func (uc *ChangePasswordUseCase) Execute(id uuid.UUID, currentPassword, newPassword string) (*entity.UserEnt, error) {
	repo := uc.userRepo
	hashService := uc.hashService

	user, err := repo.GetUserById(id)
	if err != nil {
		return nil, err
	}

	if err := hashService.Compare(user.Password, currentPassword); err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := hashService.Compare(user.Password, newPassword); err == nil {
		return nil, errors.New("the new password must be different from the current one")
	}

	hashPass, err := hashService.Has(newPassword)
	if err != nil {
		return nil, err
	}

	user.Password = hashPass
	updatedUser, err := repo.Update(id, *user.UpdateUser())
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
