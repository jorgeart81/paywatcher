package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type envVars struct {
	APP_HOST          string
	APP_PORT          int
	GIN_MODE          string
	CORS_ALLOW_ORIGIN string
	// Database
	DB_HOST         string
	DB_PORT         int
	DB_USER         string
	DB_PASSWORD     string
	DB_NAME         string
	SSLMODE         string
	TIMEZONE        string
	CONNECT_TIMEOUT int
	// JWT
	DOMAIN        string
	COOKIE_DOMAIN string
	JWT_SECRET    string
	JWT_ISSUER    string
	JWT_AUDIENCE  string
}

func loadEnv() (*envVars, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	envs := envVars{
		APP_HOST:          os.Getenv("APP_HOST"),
		APP_PORT:          parseInt(os.Getenv("APP_PORT"), "error parsing APP_PORT"),
		GIN_MODE:          os.Getenv("GIN_MODE"),
		CORS_ALLOW_ORIGIN: os.Getenv("CORS_ALLOW_ORIGIN"),

		DB_HOST:         os.Getenv("DB_HOST"),
		DB_PORT:         parseInt(os.Getenv("DB_PORT"), "error parsing DB_PORT"),
		DB_USER:         os.Getenv("DB_USER"),
		DB_PASSWORD:     os.Getenv("DB_PASSWORD"),
		DB_NAME:         os.Getenv("DB_NAME"),
		SSLMODE:         os.Getenv("SSLMODE"),
		TIMEZONE:        os.Getenv("TIMEZONE"),
		CONNECT_TIMEOUT: parseInt(os.Getenv("CONNECT_TIMEOUT"), "error parsing CONNECT_TIMEOUT"),

		DOMAIN:        os.Getenv("DOMAIN"),
		COOKIE_DOMAIN: os.Getenv("COOKIE_DOMAIN"),
		JWT_SECRET:    os.Getenv("JWT_SECRET"),
		JWT_ISSUER:    os.Getenv("JWT_ISSUER"),
		JWT_AUDIENCE:  os.Getenv("JWT_AUDIENCE"),
	}

	return &envs, nil
}

func parseInt(value string, errorMessage string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(errorMessage+":", err)
		return 0
	}
	return intValue
}
