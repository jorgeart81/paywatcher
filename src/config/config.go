package config

import (
	"log"
)

type Config struct {
	env *environmentVariables
}

func (c *Config) Load() {
	envs, err := c.loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	c.env = envs
}

func (c *Config) GetEnvs() *environmentVariables {
	return c.env
}
