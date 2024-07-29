package services

import (
	"errors"
	"paywatcher/src/config"
	"paywatcher/src/domain/services"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTAuth struct {
	jwtIssuer     string
	jwtAudience   string
	jwtSecret     string
	jwtExpiry     time.Duration
	refreshExpiry time.Duration
}

func JWTAuthService() *JWTAuth {
	jwtConf := config.JWT
	return &JWTAuth{
		jwtIssuer:     jwtConf.Issuer,
		jwtAudience:   jwtConf.Audience,
		jwtSecret:     jwtConf.Secret,
		jwtExpiry:     jwtConf.Expiry,
		refreshExpiry: jwtConf.RefreshExpiry,
	}
}

type jwtClaims struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"sub"`
	jwt.RegisteredClaims
}

func (a *JWTAuth) GenerateTokenPair(user *services.AuthUser) (services.TokenPairs, error) {
	now := time.Now().UTC()
	// Create a token
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"sub":      user.ID,
		"aud":      a.jwtAudience,
		"iss":      a.jwtIssuer,
		"iat":      now.Unix(),
		"type":     "JWT",
		"exp":      now.Add(a.jwtExpiry).Unix(),
	})

	// Create a signed token
	signedAccessToken, err := accessToken.SignedString([]byte(a.jwtSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create a refresh token and set claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"iat":  now.Unix(),
		"type": "JWT",
		"exp":  now.Add(a.refreshExpiry).Unix(),
	})

	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(a.jwtSecret))
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
		return []byte(a.jwtSecret), nil
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
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
