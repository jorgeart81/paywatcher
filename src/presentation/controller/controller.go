package controller

import (
	"paywatcher/src/application/usecases"
	"paywatcher/src/domain/services"
	"paywatcher/src/infrastructure/userinfra"

	"gorm.io/gorm"
)

var (
	userController *UserController
	authService    services.Authenticator
	hashService    services.HashService
)

type Controller struct {
	User *UserController
}

func InitializeController(db *gorm.DB, authServ services.Authenticator, hashServ services.HashService) {
	authService = authServ
	hashService = hashServ
	userController = initUserController(db)
}

func GetControllers() *Controller {
	return &Controller{
		User: userController,
	}
}

func initUserController(db *gorm.DB) *UserController {
	// jwtConf := config.JWT

	// hashService := services.NewBcryptService()
	// authService := &services.JWTAuth{
	// 	JWTIssuer:     jwtConf.Issuer,
	// 	JWTAudience:   jwtConf.Audience,
	// 	JWTSecret:     jwtConf.Secret,
	// 	JWTExpiry:     jwtConf.Expiry,
	// 	RefreshExpiry: jwtConf.RefreshExpiry,
	// 	CookieDomain:  jwtConf.CookieDomain,
	// 	CookiePath:    jwtConf.CookiePath,
	// 	CookieName:    jwtConf.CookieName,
	// }

	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)

	createUserUC := usecases.NewCreateUserUseCase(userRepositoryImpl, authService, hashService)
	loginUserUC := usecases.NewLoginUserUseCase(userRepositoryImpl, authService, hashService)

	// Create and return the controller
	return newUserController(authService, createUserUC, loginUserUC)
}
