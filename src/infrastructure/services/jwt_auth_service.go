package services

import (
	"errors"
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
		"aud":      a.JWTAudience,
		"iss":      a.JWTIssuer,
		"iat":      now.Unix(),
		"type":     "JWT",
		"exp":      now.Add(a.JWTExpiry).Unix(),
	})

	// Create a signed token
	signedAccessToken, err := accessToken.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return services.TokenPairs{}, err
	}

	// Create a refresh token and set claims
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"iat":  now.Unix(),
		"type": "JWT",
		"exp":  now.Add(a.RefreshExpiry).Unix(),
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
		}, nil
	} else {
		return nil, errors.New("invalid token")
	}
}
