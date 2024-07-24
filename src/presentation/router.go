package presentation

import (
	"paywatcher/src/application/auth"
	"paywatcher/src/application/usecases"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/services"
	"paywatcher/src/infrastructure/userinfra"
	"paywatcher/src/presentation/userctrl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AppRouter struct {
	app *gin.Engine
	db  *gorm.DB
}

func NewAppRouter(app *gin.Engine, db *gorm.DB) *AppRouter {
	return &AppRouter{app: app, db: db}
}

func (appRouter *AppRouter) Init() {
	r := appRouter.app
	api := r.Group("/api")

	userController := initUserController(appRouter.db)

	{
		api.GET("/", userController.Index)
		api.POST("/user/create", userController.Create)
		api.POST("/user/login", userController.Login)
	}
}

func initUserController(db *gorm.DB) *userctrl.UserController {
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
	return userctrl.NewUserController(createUserUC, loginUserUC)
}
