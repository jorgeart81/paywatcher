package main

import (
	"paywatcher/config"
	"paywatcher/database"
	"paywatcher/presentation"
)

func main() {
	config := config.GetConfig()
	env := config.Env

	database.Connect()

	// Start server
	server := presentation.Server{
		Port: env.APP_PORT,
		Host: env.DB_HOST,
	}
	server.Start()
}
