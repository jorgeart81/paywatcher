package router

import (
	"fmt"
	"log"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/services"
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

	jwtConf := config.JWT
	authService := &services.JWTAuth{
		JWTIssuer:     jwtConf.Issuer,
		JWTAudience:   jwtConf.Audience,
		JWTSecret:     jwtConf.Secret,
		JWTExpiry:     jwtConf.Expiry,
		RefreshExpiry: jwtConf.RefreshExpiry,
		CookieDomain:  jwtConf.CookieDomain,
		CookiePath:    jwtConf.CookiePath,
		CookieName:    jwtConf.CookieName,
	}
	hashService := services.NewBcryptService()

	controller.InitializeController(db, authService, hashService)
	controllers := controller.GetControllers()

	routes := &appRoutes{
		userController: controllers.User,
		authService:    authService,
	}
	routes.initializeRoutes(router)

	addr := fmt.Sprintf("%s:%d", host, port)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}
