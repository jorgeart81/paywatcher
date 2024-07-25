package server

import (
	"log"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/database"
	"paywatcher/src/presentation/router"
)

var logger *config.Logger

func Start() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()

	config.Initialize()
	logger = config.GetLogger("server")
	logger.Info("starting the server")

	// Connect to database
	db := config.Database
	postgresDB := database.PotsgresDB{
		Host:           db.Host,
		Port:           db.Port,
		User:           db.User,
		Password:       db.Password,
		DBName:         db.DBName,
		SSLMode:        db.SSLMode,
		Timezone:       db.Timezone,
		ConnectTimeout: db.ConnectTimeout,
	}
	DB := postgresDB.Connect()

	if DB == nil {
		log.Fatal("failed to connect to database")
	}

	// Start server
	serv := config.Server
	logger.Infof("starting server on port %d", serv.Port)
	router.Initialize(serv.Port, serv.Host, serv.GinMode, DB)
}
