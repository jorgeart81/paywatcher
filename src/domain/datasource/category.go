package datasource

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type CategoryDS interface {
	Save(category entity.CategoryEnt, userID uuid.UUID) (*entity.CategoryEnt, error)
}
