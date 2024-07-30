package repositories

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type UserRepository interface {
	Save(user entity.UserEnt) (*entity.UserEnt, error)
	GetUserById(id uuid.UUID) (*entity.UserEnt, error)
	GetUserByEmail(email string) (*entity.UserEnt, error)
	Update(user entity.UserEnt) (*entity.UserEnt, error)
}
