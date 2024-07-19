package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type environmentVariables struct {
	APP_HOST        string
	APP_PORT        int
	DB_HOST         string
	DB_PORT         int
	DB_USER         string
	DB_PASSWORD     string
	DB_NAME         string
	SSLMODE         string
	TIMEZONE        string
	CONNECT_TIMEOUT int
}

func (c *Config) loadEnv() (*environmentVariables, error) {

	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}

	envs := &environmentVariables{
		APP_HOST:        os.Getenv("APP_HOST"),
		APP_PORT:        parseInt(os.Getenv("APP_PORT"), "error parsing APP_PORT"),
		DB_HOST:         os.Getenv("DB_HOST"),
		DB_PORT:         parseInt(os.Getenv("DB_PORT"), "error parsing DB_PORT"),
		DB_USER:         os.Getenv("DB_USER"),
		DB_PASSWORD:     os.Getenv("DB_PASSWORD"),
		DB_NAME:         os.Getenv("DB_NAME"),
		SSLMODE:         os.Getenv("SSLMODE"),
		TIMEZONE:        os.Getenv("TIMEZONE"),
		CONNECT_TIMEOUT: parseInt(os.Getenv("CONNECT_TIMEOUT"), "error parsing CONNECT_TIMEOUT"),
	}

	return envs, nil
}

func parseInt(value string, errorMessage string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(errorMessage+":", err)
		return 0
	}
	return intValue
}
