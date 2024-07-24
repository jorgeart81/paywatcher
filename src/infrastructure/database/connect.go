package database

import (
	"fmt"
	"paywatcher/src/infrastructure/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PotsgresDB struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	SSLMode        string
	Timezone       string
	ConnectTimeout int
}

func (db *PotsgresDB) Connect() *gorm.DB {
	var err error

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode, db.Timezone, db.ConnectTimeout)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database because: %s!", err.Error()))
	}

	fmt.Println("database connected")

	// Schema migration
	fmt.Println("migrating schemas...")
	DB.AutoMigrate(model.User{})
	DB.AutoMigrate(model.Payment{})
	DB.AutoMigrate(model.Category{})
	fmt.Println("migrated schemas!")

	return DB
}
