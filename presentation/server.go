package presentation

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	Port int
	Host string
}

func (s *Server) Start() {

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	app := fiber.New()

	router := AppRouter{app: app}
	router.Init()

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
