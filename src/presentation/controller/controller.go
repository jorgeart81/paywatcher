package controller

import (
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/infrastructure/datasource"
	"paywatcher/src/infrastructure/repositories"
	"paywatcher/src/infrastructure/services"

	"gorm.io/gorm"
)

var (
	userController *AuthController
)

type Controller struct {
	User *AuthController
}

func InitializeController(db *gorm.DB) {
	userController = initAuthController(db)
}

func GetControllers() *Controller {
	return &Controller{
		User: userController,
	}
}

func initAuthController(db *gorm.DB) *AuthController {
	authService := services.JWTAuthService()
	hashService := services.NewBcryptService()
	// Create datasource, repository and use case
	userDatasource := &datasource.PostgresUserDatasrc{DB: db}
	userRepositoryImpl := repositories.NewUserRepository(userDatasource)

	createUserUC := user.NewRegisterUserUseCase(userRepositoryImpl, authService, hashService)
	loginUserUC := user.NewLoginUserUseCase(userRepositoryImpl, authService, hashService)
	refreshTokenUC := user.NewRefreshTokenUseCase(userRepositoryImpl, authService)
	changePasswordUC := user.NewChangePasswordUseCase(userRepositoryImpl, authService, hashService)
	disableUserUC := user.NewSoftDeleteUserUseCase(userRepositoryImpl, hashService)

	// Create and return the controller
	return newAuthController(createUserUC, loginUserUC, refreshTokenUC, changePasswordUC, disableUserUC)
}
