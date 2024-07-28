package middlewares

import (
	"net/http"
	"paywatcher/src/domain/services"
	"paywatcher/src/presentation/response"
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	authService services.Authenticator
}

func NewAuthMiddleware(authService services.Authenticator) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Response vary
		ctx.Header("Vary", "Authorization")

		// Get Authorization Header
		authHeader := ctx.Request.Header.Get("Authorization")
		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": response.GenericError{
				Message: "missing authorization header",
			}})
			ctx.Abort()
			return
		}

		// Trim if we have the word Bearer
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// Verify the token value
		_, err := m.authService.VerifyToken(tokenString)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": response.GenericError{
				Message: "unauthorized",
			}})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
