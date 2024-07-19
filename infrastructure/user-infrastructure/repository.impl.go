package userinfrastructure

import (
	"paywatcher/database/model"
	userdomain "paywatcher/domain/user-domain"
)

type UserRepositoryImpl struct {
	Datasource userdomain.UserDatasource
}

func NewUserRepository(datasource userdomain.UserDatasource) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		Datasource: datasource,
	}
}

func (r *UserRepositoryImpl) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}
