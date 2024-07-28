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
	dbConf := config.Database
	postgresDB := database.PotsgresDB{
		Host:           dbConf.Host,
		Port:           dbConf.Port,
		User:           dbConf.User,
		Password:       dbConf.Password,
		DBName:         dbConf.DBName,
		SSLMode:        dbConf.SSLMode,
		Timezone:       dbConf.Timezone,
		ConnectTimeout: dbConf.ConnectTimeout,
	}
	db, err := postgresDB.Connect()

	if err != nil {
		logger.Errorf("error initializing database: %v", err)
		return
	}
	if db == nil {
		log.Fatal("failed to connect to database")
	}

	// Start server
	serv := config.Server
	logger.Infof("starting server on port %d", serv.Port)
	router.Initialize(serv.Port, serv.Host, serv.GinMode, db, serv.CorsAllowOrigin)
}
