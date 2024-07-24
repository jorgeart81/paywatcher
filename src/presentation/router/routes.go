package router

import (
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"
)

type appRoutes struct {
	userController *controller.UserController
}

func (routes *appRoutes) initializeRoutes(router *gin.Engine) {
	api := router.Group("/api")

	{
		userController := routes.userController
		api.GET("/", userController.Index)
		api.POST("/user/create", userController.Create)
		api.POST("/user/login", userController.Login)
	}
}
