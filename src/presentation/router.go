package presentation

import (
	"paywatcher/src/application/auth"
	"paywatcher/src/application/usecases"
	"paywatcher/src/config"
	"paywatcher/src/infrastructure/userinfra"
	"paywatcher/src/presentation/userctrl"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppRouter struct {
	app *fiber.App
	db  *gorm.DB
}

func NewAppRouter(app *fiber.App, db *gorm.DB) *AppRouter {
	return &AppRouter{app: app, db: db}
}

func (appRouter *AppRouter) Init() {
	app := appRouter.app
	api := app.Group("/api")

	userController := initUserController(appRouter.db)

	api.Get("/", userController.Index)
	api.Post("/user/create", userController.Create)
	api.Post("/user/login", userController.Login)
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
	return userctrl.NewUserController(*createUserUC, *loginUserUC)
}
