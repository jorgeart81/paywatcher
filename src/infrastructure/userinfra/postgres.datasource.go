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

// Save implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) Save(user userdomain.User) (*userdomain.User, error) {
	db := pu.DB
	userEntity := model.ToUserEntity(&user)

	if err := db.Save(&userEntity).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, fmt.Errorf("user could not be created")
		}
		return nil, err
	}

	return userEntity.ToDomain(), nil
}

// GetUserById implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) GetUserById(id uuid.UUID) (*userdomain.User, error) {
	var userEntity model.User
	db := pu.DB

	if err := db.Find(&userEntity, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}

	return userEntity.ToDomain(), nil
}

// GetUserByEmail implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) GetUserByEmail(email string) (*userdomain.User, error) {
	var userEntity model.User
	db := pu.DB

	if err := db.Find(&userEntity, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return userEntity.ToDomain(), nil
}
