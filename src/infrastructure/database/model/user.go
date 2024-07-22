package model

import (
	"paywatcher/src/domain/userdomain"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Email     string    `gorm:"column:email;unique"`
	Username  string    `gorm:"column:username"`
	Password  string    `gorm:"column:password"`
	Role      []string  `gorm:"column:role;serializer:json"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func ToUserEntity(user *userdomain.User) *User {
	return &User{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
		Role:     ToStringRoles(user.Role),
	}
}

func (e *User) ToDomain() *userdomain.User {
	return &userdomain.User{
		ID:       e.ID,
		Username: e.Username,
		Email:    e.Email,
		Password: e.Password,
		Role:     ToDomainRoles(e.Role),
	}
}

func ToDomainRoles(roles []string) []userdomain.Role {
	domainRoles := make([]userdomain.Role, len(roles))
	for i, role := range roles {
		domainRoles[i] = userdomain.Role(role)
	}
	return domainRoles
}

func ToStringRoles(roles []userdomain.Role) []string {
	stringRoles := make([]string, len(roles))
	for i, role := range roles {
		stringRoles[i] = string(role)
	}
	return stringRoles
}
