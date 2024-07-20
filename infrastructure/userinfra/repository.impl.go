package userinfra

import (
	"paywatcher/database/model"
	"paywatcher/domain/userdomain"

	"github.com/google/uuid"
)

// var _ userdomain.UserRepository = &UserRepositoryImpl{}

type UserRepositoryImpl struct {
	Datasource userdomain.UserDatasource
}

var userRepository *UserRepositoryImpl

// GetUserByEmail implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}

// GetUserById implements userdomain.UserRepository.
func (u *UserRepositoryImpl) GetUserById(id uuid.UUID) (*model.User, error) {
	panic("unimplemented")
}

func NewUserRepository(datasource userdomain.UserDatasource) *UserRepositoryImpl {
	if userRepository == nil {
		userRepository = &UserRepositoryImpl{
			Datasource: datasource,
		}
	}
	return userRepository
}
