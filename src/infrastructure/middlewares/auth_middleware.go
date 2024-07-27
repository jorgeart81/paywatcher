package middlewares

import (
	"net/http"
	"paywatcher/src/domain/services"
	"paywatcher/src/presentation/response"

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
		_, err := m.authService.GetTokenFromHeaderAndVerify(ctx.Writer, ctx.Request)
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
