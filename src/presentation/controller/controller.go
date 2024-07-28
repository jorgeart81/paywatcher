package controller

import (
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/domain/services"
	iServices "paywatcher/src/infrastructure/services"
	"paywatcher/src/infrastructure/userinfra"

	"gorm.io/gorm"
)

var (
	userController *AuthController
	authService    services.Authenticator
)

type Controller struct {
	User *AuthController
}

func InitializeController(db *gorm.DB, authServ services.Authenticator) {
	authService = authServ
	userController = initAuthController(db)
}

func GetControllers() *Controller {
	return &Controller{
		User: userController,
	}
}

func initAuthController(db *gorm.DB) *AuthController {
	hashService := iServices.NewBcryptService()
	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)

	createUserUC := user.NewRegisterUserUseCase(userRepositoryImpl, authService, hashService)
	loginUserUC := user.NewLoginUserUseCase(userRepositoryImpl, authService, hashService)
	refreshTokenUC := user.NewRefreshTokenUseCase(userRepositoryImpl, authService)

	// Create and return the controller
	return newAuthController(createUserUC, loginUserUC, refreshTokenUC)
}
