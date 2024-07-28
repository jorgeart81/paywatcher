package config

import (
	"log"
	"time"
)

var (
	Database *database
	Server   *server
	JWT      *jwt
	Cookie   *cookie

	logger *Logger
)

func Initialize() {
	e, err := loadEnv()
	if err != nil {
		log.Fatal(err)
		return
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
		Host:            e.APP_HOST,
		Port:            e.APP_PORT,
		GinMode:         e.GIN_MODE,
		CorsAllowOrigin: e.CORS_ALLOW_ORIGIN,
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

	Cookie = &cookie{
		CookieDomain:  e.COOKIE_DOMAIN,
		CookiePath:    "/",
		CookieName:    "refresh_token",
		RefreshExpiry: JWT.RefreshExpiry,
	}
}

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}
