package datasource

import (
	"errors"
	"fmt"
	"paywatcher/src/domain/entity"
	"paywatcher/src/infrastructure/database/schemas"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var _ datasource.CategoryDS = &PostgresCategoryDatasrc{}

type PostgresCategoryDatasrc struct {
	DB *gorm.DB
}

// Save implements datasource.CategoryDS.
func (p *PostgresCategoryDatasrc) Save(category entity.CategoryEnt, userID uuid.UUID) (*entity.CategoryEnt, error) {
	db := p.DB
	categorySchema := schemas.ToCategorySchema(&category)

	var user schemas.User
	db.Preload("Categories").First(&user, userID)

	user.Categories = append(user.Categories, *categorySchema)

	if err := db.Save(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return nil, fmt.Errorf("category could not be created")
		}
		return nil, err
	}
	return categorySchema.ToDomain(), nil
}

// GetCategories implements datasource.CategoryDS.
func (p *PostgresCategoryDatasrc) GetCategories(userID uuid.UUID) (*[]entity.CategoryEnt, error) {
	db := p.DB
	var user schemas.User
	var categories []entity.CategoryEnt

	if err := db.Preload("Categories").First(&user, userID).Error; err != nil {
		return nil, err
	}

	for _, category := range user.Categories {
		fmt.Println(category.ToDomain())
		categories = append(categories, *category.ToDomain())
	}

	return &categories, nil
}
