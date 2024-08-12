package repositories

import (
	"paywatcher/src/domain/datasource"
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

// var _ userdomain.UserRepository = &UserRepositoryImpl{}

type UserRepositoryImpl struct {
	Datasource datasource.UserDS
}

var userRepository *UserRepositoryImpl

func NewUserRepository(datasource datasource.UserDS) *UserRepositoryImpl {
	if userRepository == nil {
		userRepository = &UserRepositoryImpl{
			Datasource: datasource,
		}
	}
	return userRepository
}

// Save implements userdomain.UserRepository.
func (u *UserRepositoryImpl) Save(user entity.UserEnt) (*entity.UserEnt, error) {
	return u.Datasource.Save(user)
}

// GetUserById implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserById(id uuid.UUID) (*entity.UserEnt, error) {
	return u.Datasource.GetUserById(id)
}

// GetUserByEmail implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*entity.UserEnt, error) {
	return u.Datasource.GetUserByEmail(email)
}

// Save implements userdomain.UserRepository.
func (u *UserRepositoryImpl) Update(id uuid.UUID, user entity.UserEnt) (*entity.UserEnt, error) {
	return u.Datasource.Update(id, user)
}
