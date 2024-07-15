package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type configDatabase struct {
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
}

func GetConfig() configDatabase {
	var config configDatabase

	err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	config.DB_HOST = os.Getenv("DB_HOST")
	config.DB_PORT = parseInt(os.Getenv("DB_PORT"), "error parsing DB_PORT")
	config.DB_USER = os.Getenv("DB_USER")
	config.DB_PASSWORD = os.Getenv("DB_PASSWORD")
	config.DB_NAME = os.Getenv("DB_NAME")

	return config
}

func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %w", err)
	}
	return nil
}

func parseInt(value string, errorMessage string) int {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println(errorMessage+":", err)
		return 0
	}
	return intValue
}
