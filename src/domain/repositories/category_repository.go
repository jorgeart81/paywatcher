package repositories

import "paywatcher/src/domain/entity"

type CategoryRepository interface {
	Save(user entity.CategoryEnt) (*entity.CategoryEnt, error)
}
