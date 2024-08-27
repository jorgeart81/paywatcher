package category

import (
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"

	"github.com/google/uuid"
)

type CreateCategoryUseCase struct {
	categoryRepo repositories.CategoryRepository
}

func NewCreateCategoryUseCase(categoryRepo repositories.CategoryRepository) CreateCategoryUseCase {
	return CreateCategoryUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *CreateCategoryUseCase) Execute(category *entity.CategoryEnt, userID uuid.UUID) (*entity.CategoryEnt, error) {
	repo := uc.categoryRepo

	newCategory, err := repo.Save(*category.NewCategoryEnt(), userID)
	if err != nil {
		return nil, err
	}

	return newCategory, nil
}
