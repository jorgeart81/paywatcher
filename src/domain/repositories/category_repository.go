package repositories

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type CategoryRepository interface {
	Save(category entity.CategoryEnt, userID uuid.UUID) (*entity.CategoryEnt, error)
	GetCategories(userID uuid.UUID) (*[]entity.CategoryEnt, error)
}
