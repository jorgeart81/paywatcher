package config

import (
	"log"
	"time"
)

var Database *database
var Server *server
var JWT *jwt

func (c *Config) Init() {
	e, err := c.loadEnv()
	if err != nil {
		log.Fatal(err)
	}

	Database = &database{
		Host:           e.DB_HOST,
		Port:           e.DB_PORT,
		User:           e.DB_USER,
		Password:       e.DB_PASSWORD,
		DBName:         e.DB_NAME,
		SSLMode:        e.SSLMODE,
		Timezone:       e.TIMEZONE,
		ConnectTimeout: e.CONNECT_TIMEOUT,
	}

	Server = &server{
		Host:    e.APP_HOST,
		Port:    e.APP_PORT,
		GinMode: e.GIN_MODE,
	}

	JWT = &jwt{
		Issuer:        e.JWT_ISSUER,
		Audience:      e.JWT_AUDIENCE,
		Secret:        e.JWT_SECRET,
		Expiry:        time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  e.COOKIE_DOMAIN,
		CookiePath:    "/",
		CookieName:    "refresh_token",
	}
}
