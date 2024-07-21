package model

import (
	"paywatcher/domain/userdomain"
	"time"

	"github.com/google/uuid"
)

type UserEntity struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func ToUserEntity(user *userdomain.User) *UserEntity {
	return &UserEntity{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
}

func (e *UserEntity) ToDomain() *userdomain.User {
	return &userdomain.User{
		ID:       e.ID,
		Name:     e.Name,
		Email:    e.Email,
		Username: e.Username,
		Password: e.Password,
	}
}
