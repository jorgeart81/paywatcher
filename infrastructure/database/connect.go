package database

import (
	"fmt"
	"paywatcher/infrastructure/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PotsgresDB struct {
	db *gorm.DB
}

// var PotsgresDB *gorm.DB
func (p PotsgresDB) GetDB() *gorm.DB {
	return p.db
}

func (p *PotsgresDB) Connect(dsn string) {
	var err error

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database because: %s!", err.Error()))
	}

	fmt.Println("database connected")

	// Schema migration
	fmt.Println("migrating schemas...")
	DB.AutoMigrate(model.UserEntity{})
	DB.AutoMigrate(model.PaymentEntity{})
	DB.AutoMigrate(model.CategoryEntity{})
	fmt.Println("migrated schemas!")

	p.db = DB
}
