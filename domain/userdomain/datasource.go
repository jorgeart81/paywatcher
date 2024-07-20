package userdomain

import (
	"paywatcher/database/model"

	"github.com/google/uuid"
)

type UserDatasource interface {
	GetUserById(id uuid.UUID) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
