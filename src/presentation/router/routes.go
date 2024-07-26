package router

import (
	"paywatcher/docs"
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type appRoutes struct {
	userController *controller.UserController
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

	{
		userController := ar.userController
		api.GET("/", userController.Index)
		api.POST("/user/register", userController.Create)
		api.POST("/user/login", userController.Login)
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
