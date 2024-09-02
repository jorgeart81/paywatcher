package controller

import (
	"paywatcher/src/application/usecases/category"
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/infrastructure/datasource"
	"paywatcher/src/infrastructure/repositories"
	"paywatcher/src/infrastructure/services"

	"gorm.io/gorm"
)

var (
	authController     *AuthController
	categoryController *CategoryController
)

type Controller struct {
	Auth     *AuthController
	Category *CategoryController
}

func InitializeController(db *gorm.DB) {
	authController = initAuthController(db)
	categoryController = initCategoryController(db)
}

func GetControllers() *Controller {
	return &Controller{
		Auth:     authController,
		Category: categoryController,
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

func initCategoryController(db *gorm.DB) *CategoryController {
	categoryDatasource := &datasource.PostgresCategoryDatasrc{DB: db}
	categoryRepositoryImpl := repositories.NewcCategoryRepository(categoryDatasource)

	createCategoryUC := category.NewCreateCategoryUseCase(categoryRepositoryImpl)
	userCategoriesUC := category.NewUserCategoriesUseCase(categoryRepositoryImpl)

	return newCategoryController(createCategoryUC, userCategoriesUC)
}
