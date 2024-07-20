package userinfra

import (
	"errors"
	"fmt"
	"paywatcher/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDatasource struct {
	DB *gorm.DB
}

func (datasource *UserDatasource) GetUserById(id uuid.UUID) (*model.User, error) {
	var user model.User
	db := datasource.DB

	err := db.Find(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}

	return &user, nil
}

func (datasource *UserDatasource) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")
}
