package userinfra

import (
	"paywatcher/src/domain/userdomain"

	"github.com/google/uuid"
)

// var _ userdomain.UserRepository = &UserRepositoryImpl{}

type UserRepositoryImpl struct {
	Datasource userdomain.UserDatasource
}

var userRepository *UserRepositoryImpl

func NewUserRepository(datasource userdomain.UserDatasource) *UserRepositoryImpl {
	if userRepository == nil {
		userRepository = &UserRepositoryImpl{
			Datasource: datasource,
		}
	}
	return userRepository
}

// Save implements userdomain.UserRepository.
func (u *UserRepositoryImpl) Save(user userdomain.User) (*userdomain.User, error) {
	return u.Datasource.Save(user)
}

// GetUserById implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserById(id uuid.UUID) (*userdomain.User, error) {
	return u.Datasource.GetUserById(id)
}

// GetUserByEmail implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*userdomain.User, error) {
	return u.Datasource.GetUserByEmail(email)
}
