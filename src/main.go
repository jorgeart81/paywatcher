package main

import (
	"fmt"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/database"
	"paywatcher/src/presentation"
)

func main() {
	var config config.Config
	config.Load()
	env := config.GetEnvs()

	// Connect to database
	var postgresDB database.PotsgresDB
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s timezone=%s connect_timeout=%d",
		env.DB_HOST, env.DB_PORT, env.DB_USER, env.DB_PASSWORD, env.DB_NAME, env.SSLMODE, env.TIMEZONE, env.CONNECT_TIMEOUT)
	postgresDB.Connect(dsn)

	// Start server
	server := presentation.Server{
		Port: env.APP_PORT,
		Host: env.DB_HOST,
		DB:   postgresDB.GetDB(),
	}

	server.Start()
}
