package datasource

import (
	"errors"
	"fmt"
	"paywatcher/src/domain/entity"
	"paywatcher/src/infrastructure/database/schemas"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// var _ userdomain.UserDatasource = &PostgresUserDatasrc{}

type PostgresUserDatasrc struct {
	DB *gorm.DB
}

// Save implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) Save(user entity.UserEnt) (*entity.UserEnt, error) {
	var pgErr *pgconn.PgError
	db := pu.DB
	userSchema := schemas.ToUserSchema(&user)

	if err := db.Save(&userSchema).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, fmt.Errorf("user could not be created")
		}
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return nil, fmt.Errorf("duplicate key, user could not be created")
		}
		return nil, err
	}

	return userSchema.ToDomain(), nil
}

// GetUserById implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) GetUserById(id uuid.UUID) (*entity.UserEnt, error) {
	var userSchema schemas.User
	db := pu.DB

	if err := db.First(&userSchema, "id = ?", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with ID %s not found", id)
		}
		return nil, err
	}

	return userSchema.ToDomain(), nil
}

// GetUserByEmail implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) GetUserByEmail(email string) (*entity.UserEnt, error) {
	var userSchema schemas.User
	db := pu.DB

	if err := db.First(&userSchema, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("user with email %s not found", email)
		}
		return nil, err
	}

	return userSchema.ToDomain(), nil
}

// Save implements userdomain.UserDatasource.
func (pu *PostgresUserDatasrc) Update(user entity.UserEnt) (*entity.UserEnt, error) {
	db := pu.DB
	userSchema := schemas.User{}
	entityToUserSchema := schemas.ToUserSchema(&user)

	if err := db.Model(&userSchema).
		Clauses(clause.Returning{}).
		Where("id = ?", user.ID).
		Select("username", "email", "password", "role", "active").
		Updates(entityToUserSchema).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, fmt.Errorf("user could not be updated")
		}
		return nil, err
	}

	userSchema.ID = user.ID
	return userSchema.ToDomain(), nil
}

// SoftDelete implements datasource.UserDS.
func (pu *PostgresUserDatasrc) SoftDelete(id uuid.UUID) error {
	db := pu.DB
	now := time.Now()
	userSchema := schemas.User{}

	if err := db.Model(&userSchema).
		Clauses(clause.Returning{}).
		Where("id = ?", id).
		Select("active", "deleted_at").
		Updates(schemas.User{Active: false, DeletedAt: &now}).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return fmt.Errorf("user could not be updated")
		}
		return err
	}

	if userSchema.DeletedAt == nil && userSchema.Active {
		return fmt.Errorf("user could not be updated")
	}

	return nil
}
