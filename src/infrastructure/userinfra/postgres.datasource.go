package userinfra

import (
	"errors"
	"fmt"
	"paywatcher/src/domain/userdomain"
	"paywatcher/src/infrastructure/database/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var _ userdomain.UserDatasource = &PostgresUserDatasrc{}

type PostgresUserDatasrc struct {
	DB *gorm.DB
}

// GetUserById implements userdomain.UserDatasource.
func (u *PostgresUserDatasrc) GetUserById(id uuid.UUID) (*userdomain.User, error) {
	var user model.UserEntity
	db := u.DB

	err := db.Find(&user, "id = ?", id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}

	return user.ToDomain(), nil
}

// GetUserByEmail implements userdomain.UserDatasource.
func (u *PostgresUserDatasrc) GetUserByEmail(email string) (*userdomain.User, error) {
	panic("unimplemented")

}
