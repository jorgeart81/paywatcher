package userdomain

import "paywatcher/database/model"

type UserDatasource interface {
	GetUserById(id int) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
