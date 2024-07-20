package userinfra

import (
	"paywatcher/database/model"
	"paywatcher/domain/userdomain"

	"github.com/google/uuid"
)

type UserRepositoryImpl struct {
	Datasource userdomain.UserDatasource
}

func NewUserRepository(datasource userdomain.UserDatasource) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Datasource: datasource,
	}
}

// GetUserById implements userdomain.UserRepository.
func (r *UserRepositoryImpl) GetUserById(id uuid.UUID) (*model.User, error) {
	panic("unimplemented")
}

// GetUserByEmail implements userdomain.UserRepository.
func (r *UserRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}
