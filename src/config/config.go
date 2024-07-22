package config

import (
	"log"
)

type Config struct{}

var Envs *environmentVariables

func (c *Config) Load() {
	envs, err := c.loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	Envs = envs
}
