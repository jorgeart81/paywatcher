package router

import (
	"paywatcher/docs"
	"paywatcher/src/domain/services"
	"paywatcher/src/infrastructure/middlewares"
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type appRoutes struct {
	authService    services.Authenticator
	authController *controller.AuthController
}

func (ar *appRoutes) initializeRoutes(router *gin.Engine) {
	basePath := "/api"
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Paywatcher"
	docs.SwaggerInfo.Description = "This is a sample Paywatcher server."
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "petstore.swagger.io"
	docs.SwaggerInfo.BasePath = basePath
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logger.Info("routes initialized")
	api := router.Group(basePath)

	userController := ar.authController
	{
		api.POST("/register", userController.Register)
		api.POST("/login", userController.Login)
	}

	{
		authMiddleware := middlewares.NewAuthMiddleware(ar.authService)
		authorized := api.Group("/")
		authorized.Use(authMiddleware.AuthRequired())

		authorized.GET("/test-auth", userController.Index)
		authorized.GET("/refresh-token", userController.RefreshToken)
		authorized.PUT("/change-password", userController.ChangePassword)
		authorized.GET("/logout", userController.Logout)
		authorized.PUT("/disable-user", userController.DisableUser)
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
