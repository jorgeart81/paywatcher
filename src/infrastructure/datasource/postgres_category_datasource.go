package datasource

import (
	"paywatcher/src/domain/entity"

	"gorm.io/gorm"
)

// var _ datasource.CategoryDS = &PostgresCategoryDatasrc{}

type PostgresCategoryDatasrc struct {
	DB *gorm.DB
}

// Save implements datasource.CategoryDS.
func (p *PostgresCategoryDatasrc) Save(user entity.CategoryEnt) (*entity.CategoryEnt, error) {
	panic("unimplemented")
}
