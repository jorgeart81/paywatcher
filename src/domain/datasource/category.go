package datasource

import "paywatcher/src/domain/entity"

type CategoryDS interface {
	Save(user entity.CategoryEnt) (*entity.CategoryEnt, error)
}
