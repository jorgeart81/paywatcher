package userdomain

import (
	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user User) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
