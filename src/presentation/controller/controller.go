package controller

import (
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/domain/services"
	"paywatcher/src/infrastructure/userinfra"

	"gorm.io/gorm"
)

var (
	userController *AuthController
	authService    services.Authenticator
	hashService    services.HashService
)

type Controller struct {
	User *AuthController
}

func InitializeController(db *gorm.DB, authServ services.Authenticator, hashServ services.HashService) {
	authService = authServ
	hashService = hashServ
	userController = initAuthController(db)
}

func GetControllers() *Controller {
	return &Controller{
		User: userController,
	}
}

func initAuthController(db *gorm.DB) *AuthController {
	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)

	createUserUC := user.NewRegisterUserUseCase(userRepositoryImpl, authService, hashService)
	loginUserUC := user.NewLoginUserUseCase(userRepositoryImpl, authService, hashService)
	refreshTokenUC := user.NewRefreshTokenUseCase(userRepositoryImpl, authService)

	// Create and return the controller
	return newAuthController(authService, createUserUC, loginUserUC, refreshTokenUC)
}
