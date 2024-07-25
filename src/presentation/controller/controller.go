package controller

import (
	"paywatcher/src/application/auth"
	"paywatcher/src/application/usecases"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/services"
	"paywatcher/src/infrastructure/userinfra"

	"gorm.io/gorm"
)

var (
	userController *UserController
)

type Controller struct {
	User *UserController
}

func InitializeController(db *gorm.DB) {
	userController = initUserController(db)
}

func GetControllers() *Controller {
	return &Controller{
		User: userController,
	}
}

func initUserController(db *gorm.DB) *UserController {
	jwtConf := config.JWT

	hashService := services.NewBcryptService()
	authService := &auth.Auth{
		JWTIssuer:     jwtConf.Issuer,
		JWTAudience:   jwtConf.Audience,
		JWTSecret:     jwtConf.Secret,
		JWTExpiry:     jwtConf.Expiry,
		RefreshExpiry: jwtConf.RefreshExpiry,
		CookieDomain:  jwtConf.CookieDomain,
		CookiePath:    jwtConf.CookiePath,
		CookieName:    jwtConf.CookieName,
	}

	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)

	createUserUC := usecases.NewCreateUserUseCase(userRepositoryImpl, hashService)
	loginUserUC := usecases.NewLoginUserUseCase(userRepositoryImpl, authService, hashService)

	// Create and return the controller
	return newUserController(createUserUC, loginUserUC)
}
