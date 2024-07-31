package schemas

import (
	"paywatcher/src/domain/entity"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Email     string    `gorm:"column:email;unique"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Role      []string  `gorm:"column:role;serializer:json"`
	Active    bool      `gorm:"column:active"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ToUserSchema(user *entity.UserEnt) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
		Active:   user.Active,
	}
}

func (e *User) ToDomain() *entity.UserEnt {
	return &entity.UserEnt{
		ID:       e.ID,
		Username: e.Username,
		Email:    e.Email,
		Password: e.Password,
		Role:     e.Role,
		Active:   e.Active,
	}
}
