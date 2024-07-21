package presentation

import (
	"paywatcher/infrastructure/userinfra"
	"paywatcher/presentation/userctrl"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppRouter struct {
	app *fiber.App
	db  *gorm.DB
}

func (appRouter *AppRouter) Init() {
	app := appRouter.app
	api := app.Group("/api")

	userDatasource := &userinfra.PostgresUserDatasrc{DB: appRouter.db}
	userRepositoryImpl := userinfra.NewUserRepository(userDatasource)
	userController := userctrl.NewUserController(userRepositoryImpl)

	api.Get("/", userController.Index)
}
