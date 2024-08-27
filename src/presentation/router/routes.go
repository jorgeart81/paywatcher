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
	authService        services.Authenticator
	authController     *controller.AuthController
	categoryController *controller.CategoryController
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
	categoryController := ar.categoryController
	{
		// Open routes
		auth := api.Group("auth/")
		auth.POST("register", userController.Register)
		auth.POST("login", userController.Login)
		auth.POST("refresh-token", userController.RefreshToken)
	}

	{
		// Auth routes
		authMiddleware := middlewares.NewAuthMiddleware(ar.authService)
		authorized := api.Group("/")
		authorized.Use(authMiddleware.AuthRequired())

		auth := authorized.Group("auth/")
		auth.GET("test-auth", userController.Index)
		auth.PATCH("change-password", userController.ChangePassword)
		auth.GET("logout", userController.Logout)
		auth.PATCH("delete", userController.SoftDeleteUser)

		category := authorized.Group("category/")
		category.POST("", categoryController.Create)
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
