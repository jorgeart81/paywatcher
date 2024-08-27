package repositories

import (
	"paywatcher/src/domain/datasource"
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

// var _ repositories.CategoryRepository = &CategoryRepositoryImpl{}

type CategoryRepositoryImpl struct {
	Datasource datasource.CategoryDS
}

var categoryRepository *CategoryRepositoryImpl

func NewcCategoryRepository(datasource datasource.CategoryDS) *CategoryRepositoryImpl {
	if categoryRepository == nil {
		categoryRepository = &CategoryRepositoryImpl{
			Datasource: datasource,
		}
	}
	return categoryRepository
}

// Save implements repositories.CategoryRepository.
func (c *CategoryRepositoryImpl) Save(cagegory entity.CategoryEnt, userID uuid.UUID) (*entity.CategoryEnt, error) {
	return c.Datasource.Save(cagegory, userID)
}
