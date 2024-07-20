package userinfra

import (
	"errors"
	"fmt"
	"paywatcher/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var _ userdomain.UserDatasource = &PostgresUserDatasrc{}

type PostgresUserDatasrc struct {
	DB *gorm.DB
}

// GetUserById implements userdomain.UserDatasource.
func (u *PostgresUserDatasrc) GetUserById(id uuid.UUID) (*model.User, error) {
	var user model.User
	db := u.DB

	err := db.Find(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByEmail implements userdomain.UserDatasource.
func (u *PostgresUserDatasrc) GetUserByEmail(email string) (*model.User, error) {
	panic("unimplemented")

}
