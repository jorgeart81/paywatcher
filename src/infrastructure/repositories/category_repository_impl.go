package repositories

import (
	"paywatcher/src/domain/datasource"
	"paywatcher/src/domain/entity"
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
func (c *CategoryRepositoryImpl) Save(user entity.CategoryEnt) (*entity.CategoryEnt, error) {
	panic("unimplemented")
}
