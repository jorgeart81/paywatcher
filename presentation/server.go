package presentation

import (
	"fmt"
	"log"
	"paywatcher/database"
	userinfrastructure "paywatcher/infrastructure/user-infrastructure"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port int
	Host string
}

func (s *Server) Start() {

	userDatasource := &userinfrastructure.UserDatasource{DB: database.PotsgresDB}
	userinfrastructure.NewUserRepository(userDatasource)

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
