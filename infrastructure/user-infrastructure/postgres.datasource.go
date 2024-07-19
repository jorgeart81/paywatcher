package userinfrastructure

import (
	"paywatcher/database/model"

	"gorm.io/gorm"
)

type UserDatasource struct {
	DB *gorm.DB
}

func (r *UserDatasource) GetUserById(id int) (*model.User, error) {
	panic("unimplemented")
}

func (r *UserDatasource) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}
