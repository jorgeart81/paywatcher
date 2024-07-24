package server

import (
	"log"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/database"
	"paywatcher/src/presentation/router"
)

func Start() {
	var conf config.Config
	conf.Init()

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
		log.Fatalf("Failed to connect to database")
	}

	// Start server
	serv := config.Server
	router.Initialize(serv.Port, serv.Host, serv.GinMode, DB)

}
