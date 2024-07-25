package router

import (
	"fmt"
	"log"
	"paywatcher/src/application/auth"
	"paywatcher/src/application/usecases"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/services"
	"paywatcher/src/infrastructure/userinfra"
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

	routes := &appRoutes{
		userController: initUserController(db),
	}
	routes.initializeRoutes(router)
	logger.Info("routes initialized")

	addr := fmt.Sprintf("%s:%d", host, port)
	if err := router.Run(addr); err != nil {
		log.Fatal(err)
	}
}

func initUserController(db *gorm.DB) *controller.UserController {
	jwt := config.JWT

	hashService := services.NewBcryptService()
	authService := &auth.Auth{
		JWTIssuer:     jwt.Issuer,
		JWTAudience:   jwt.Audience,
		JWTSecret:     jwt.Secret,
		JWTExpiry:     jwt.Expiry,
		RefreshExpiry: jwt.RefreshExpiry,
		CookieDomain:  jwt.CookieDomain,
		CookiePath:    jwt.CookiePath,
		CookieName:    jwt.CookieName,
	}

	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepository := userinfra.NewUserRepository(userDatasource)

	createUserUC := usecases.NewCreateUserUseCase(userRepository, hashService)
	loginUserUC := usecases.NewLoginUserUseCase(userRepository, authService, hashService)

	// Create and return the controller
	return controller.NewUserController(createUserUC, loginUserUC)
}
