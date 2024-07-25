package database

import (
	"fmt"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/database/schemas"

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

func (p *PotsgresDB) Connect() (*gorm.DB, error) {
	logger := config.GetLogger("database")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode, p.Timezone, p.ConnectTimeout)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database because: %s!", err.Error()))
	}
	logger.Info("connected to database")

	// Schemas migration
	logger.Info("migrating schemas")
	msgMigrationErr := "postgres automigration error"

	if err := db.AutoMigrate(schemas.User{}); err != nil {
		logger.Errorf("%s %v", msgMigrationErr, err)
		return nil, err
	}
	if err := db.AutoMigrate(schemas.Payment{}); err != nil {
		logger.Errorf("%s %v", msgMigrationErr, err)
		return nil, err
	}
	if err := db.AutoMigrate(schemas.Category{}); err != nil {
		logger.Errorf("%s %v", msgMigrationErr, err)
		return nil, err
	}

	logger.Info("migrated schemas!")

	return db, nil
}
