package database

import (
	"fmt"
	"paywatcher/config"
	"paywatcher/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	var err error

	config := config.GetConfig()
	env := config.Env

	// Connect to database
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.SSLMODE, env.TIMEZONE, env.CONNECT_TIMEOUT)

	DB, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
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
