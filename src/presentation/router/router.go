package router

import (
	"fmt"
	"log"
	"paywatcher/src/config"
	"paywatcher/src/presentation/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var logger *config.Logger

func Initialize(port int, host string, ginMode string, db *gorm.DB) {
	logger = config.GetLogger("router")

	if len(ginMode) > 0 {
		gin.SetMode(ginMode)
	}

	router := gin.Default()
	logger.Info("router created")

	controller.InitializeController(db)
	controllers := controller.GetControllers()

	routes := &appRoutes{
		userController: controllers.User,
	}
	routes.initializeRoutes(router)

	addr := fmt.Sprintf("%s:%d", host, port)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
