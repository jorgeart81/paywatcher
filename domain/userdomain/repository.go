package userdomain

import (
	"paywatcher/database/model"

	"github.com/google/uuid"
)

type UserRepository interface {
	GetUserById(id uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
