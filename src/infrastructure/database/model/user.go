package model

import (
	"paywatcher/src/domain/userdomain"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func ToUserEntity(user *userdomain.User) *User {
	return &User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	}
}

func (e *User) ToDomain() *userdomain.User {
	return &userdomain.User{
		ID:       e.ID,
		Name:     e.Name,
		Email:    e.Email,
		Username: e.Username,
		Password: e.Password,
	}
}
