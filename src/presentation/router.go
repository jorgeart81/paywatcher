package presentation

import (
	"paywatcher/src/application/auth"
	"paywatcher/src/application/usecases"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/userinfra"
	"paywatcher/src/presentation/userctrl"
	"time"

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
	env := config.Envs

	authService := &auth.Auth{
		JWTIssuer:     env.JWT_ISSUER,
		JWTAudience:   env.JWT_AUDIENCE,
		JWTSecret:     env.JWT_SECRET,
		JWTExpiry:     time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  env.COOKIE_DOMAIN,
		CookiePath:    "/",
		CookieName:    "refresh_token",
	}

	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepository := userinfra.NewUserRepository(userDatasource)

	createUserUC := usecases.NewCreateUserUseCase(userRepository)
	loginUserUC := usecases.NewLoginUserUseCase(userRepository, authService)

	// Create and return the controller
	return userctrl.NewUserController(createUserUC, loginUserUC)
}
