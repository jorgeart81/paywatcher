package userdomain

import (
	"github.com/google/uuid"
)

type UserDatasource interface {
	GetUserById(id uuid.UUID) (*User, error)
	GetUserByEmail(email string) (*User, error)
}
