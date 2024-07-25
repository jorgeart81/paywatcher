package database

import (
	"fmt"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/database/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var logger *config.Logger

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
	logger = config.GetLogger("database")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		db.Host, db.Port, db.User, db.Password, db.DBName, db.SSLMode, db.Timezone, db.ConnectTimeout)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database because: %s!", err.Error()))
	}

	logger.Info("connected to database")

	// Schema migration
	logger.Info("migrating schemas")
	DB.AutoMigrate(model.User{})
	DB.AutoMigrate(model.Payment{})
	DB.AutoMigrate(model.Category{})
	logger.Info("migrated schemas!")

	return DB
}
