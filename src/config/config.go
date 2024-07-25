package config

import "time"

type database struct {
	Host           string
	Port           int
	User           string
	Password       string
	DBName         string
	SSLMode        string
	Timezone       string
	ConnectTimeout int
}

type server struct {
	Host    string
	Port    int
	GinMode string
}

type jwt struct {
	Issuer        string
	Audience      string
	Secret        string
	Expiry        time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}
