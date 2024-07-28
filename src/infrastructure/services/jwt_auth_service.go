package services

import (
	"errors"
	"net/http"
	"paywatcher/src/domain/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTAuth struct {
	JWTIssuer     string
	JWTAudience   string
	JWTSecret     string
	JWTExpiry     time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtClaims struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"sub"`
	Audience string    `json:"aud"`
	Issuer   string    `json:"iss"`
	IssuedAt int64     `json:"iat"`
	Expires  int64     `json:"exp"`
	jwt.RegisteredClaims
}

func (a *JWTAuth) GenerateTokenPair(user *services.AuthUser) (services.TokenPairs, error) {
	// Create a token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"sub":      user.ID,
		"aud":      a.JWTAudience,
		"iss":      a.JWTIssuer,
		"iat":      time.Now().UTC().Unix(),
		"type":     "JWT",
		"exp":      time.Now().UTC().Add(a.JWTExpiry).Unix(),
	})

	// Create a signed token
	signedAccessToken, err := token.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create a refresh token and set claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"iat":  time.Now().UTC().Unix(),
		"type": "JWT",
		"exp":  time.Now().UTC().Add(a.RefreshExpiry).Unix(),
	})

	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create services.TokenPairs and populate with signed tokens
	var tokenPairs = services.TokenPairs{
		AccessToken:  signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// Return services.TokenPairs
	return tokenPairs, nil
}

// VerifyToken implements services.Authenticator.
func (a *JWTAuth) VerifyToken(token string) (*services.Claims, error) {

	parseToken, err := jwt.ParseWithClaims(token, &jwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(a.JWTSecret), nil
	})

	if err != nil {
		if strings.HasPrefix(err.Error(), "token is expired by") {
			return nil, errors.New("expired token")
		}
		return nil, err
	}

	if jwtClaims, ok := parseToken.Claims.(*jwtClaims); ok && parseToken.Valid {
		return &services.Claims{
			Username: jwtClaims.Username,
			ID:       jwtClaims.ID,
			Audience: jwtClaims.Audience,
			Issuer:   jwtClaims.Issuer,
			IssuedAt: jwtClaims.IssuedAt,
			Expires:  jwtClaims.Expires,
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

// GetRefreshCookie implements services.Authenticator.
func (a *JWTAuth) GetRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:     a.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(a.RefreshExpiry),
		MaxAge:   int(a.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

// GetExpiredRefreshCookie implements services.Authenticator.
func (a *JWTAuth) GetExpiredRefreshCookie() *http.Cookie {
	return &http.Cookie{
		Name:     a.CookieName,
		Path:     a.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   a.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}
