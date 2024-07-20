package presentation

import (
	"paywatcher/database"
	"paywatcher/infrastructure/userinfra"
	"paywatcher/presentation/userctrl"

	"github.com/gofiber/fiber/v2"
)

type AppRouter struct {
	app *fiber.App
}

func (appRouter *AppRouter) Init() {
	app := appRouter.app
	api := app.Group("/api")

	userDatasource := &userinfra.PostgresUserDatasrc{DB: database.PotsgresDB}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)
	userController := userctrl.NewUserController(userRepositoryImpl)

	api.Get("/", userController.Index)
}
