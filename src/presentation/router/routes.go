package router

import (
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"
)

type appRoutes struct {
	userController *controller.UserController
}

func (ar *appRoutes) initializeRoutes(router *gin.Engine) {
	logger.Info("routes initialized")

	api := router.Group("/api")

	{
		userController := ar.userController
		api.GET("/", userController.Index)
		api.POST("/user/create", userController.Create)
		api.POST("/user/login", userController.Login)
	}
}
