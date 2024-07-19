package config

import (
	"log"
)

type Config struct {
	Env *environmentVariables
}

func GetConfig() *Config {
	var config Config

	envs, err := config.loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		Env: envs,
	}
}
