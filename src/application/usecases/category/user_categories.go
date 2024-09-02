package category

import (
	"paywatcher/src/domain/entity"
	"paywatcher/src/domain/repositories"

	"github.com/google/uuid"
)

type UserCategoriesUseCase struct {
	categoryRepo repositories.CategoryRepository
}

func NewUserCategoriesUseCase(categoryRepo repositories.CategoryRepository) UserCategoriesUseCase {
	return UserCategoriesUseCase{
		categoryRepo: categoryRepo,
	}
}

func (uc *UserCategoriesUseCase) Execute(userID uuid.UUID) (*[]entity.CategoryEnt, error) {
	repo := uc.categoryRepo

	categories, err := repo.GetCategories(userID)
	if err != nil {
		return nil, err
	}

	return categories, nil
}
