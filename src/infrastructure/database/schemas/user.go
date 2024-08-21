package schemas

import (
	"paywatcher/src/domain/entity"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id" gorm:"primaryKey;type:uuid"`
	Email     string     `gorm:"column:email;unique;not null"`
	Username  string     `gorm:"column:username"`
	Password  string     `gorm:"column:password"`
	Role      []string   `gorm:"column:role;serializer:json"`
	Active    bool       `gorm:"column:active;default:true"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`

	Categories []Category `gorm:"foreignKey:UserID"`
}

func ToUserSchema(user *entity.UserEnt) *User {
	return &User{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Password:  user.Password,
		Role:      user.Role,
		Active:    user.Active,
		DeletedAt: user.DeletedAt,
	}
}

func (e *User) ToDomain() *entity.UserEnt {
	return &entity.UserEnt{
		ID:        e.ID,
		Username:  e.Username,
		Email:     e.Email,
		Password:  e.Password,
		Role:      e.Role,
		Active:    e.Active,
		DeletedAt: e.DeletedAt,
	}
}
