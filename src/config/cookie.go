package config

import (
	"net/http"
	"time"
)

func GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     Cookie.CookieName,
		Path:     Cookie.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(Cookie.RefreshExpiry),
		MaxAge:   int(Cookie.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   Cookie.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     Cookie.CookieName,
		Path:     Cookie.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   Cookie.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}
