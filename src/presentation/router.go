package presentation

import (
	"paywatcher/src/application/usecases"
	"paywatcher/src/infrastructure/userinfra"
	"paywatcher/src/presentation/userctrl"

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
	api.Post("/user", userController.Create)
}

func initUserController(db *gorm.DB) *userctrl.UserController {
	// Create datasource, repository and use case
	userDatasource := &userinfra.PostgresUserDatasrc{DB: db}
	userRepository := userinfra.NewUserRepository(userDatasource)
	createUserUC := usecases.NewCreateUserUseCase(userRepository)

	// Create and return the controller
	return userctrl.NewUserController(*createUserUC)
}
