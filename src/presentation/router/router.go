package router

import (
	"fmt"
	"log"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/middlewares"
	"paywatcher/src/infrastructure/services"
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var logger *config.Logger

func Initialize(port int, host string, ginMode string, db *gorm.DB, corsAllowOrigin string) {
	logger = config.GetLogger("router")

	if len(ginMode) > 0 {
		gin.SetMode(ginMode)
	}

	router := gin.Default()
	logger.Info("router created")
	router.Use(middlewares.EnableCORS(corsAllowOrigin))
	logger.Info("cors enabled")

	controller.InitializeController(db)
	controllers := controller.GetControllers()

	routes := &appRoutes{
		authController: controllers.Auth,
		authService:    services.JWTAuthService(),
	}
	routes.initializeRoutes(router)

	addr := fmt.Sprintf("%s:%d", host, port)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
