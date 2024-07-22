package userdomain

import (
	"github.com/google/uuid"
)

type UserDatasource interface {
	Save(user User) (*User, error)
	GetUserById(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
