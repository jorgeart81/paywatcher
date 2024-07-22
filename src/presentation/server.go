package presentation

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	Port int
	Host string
	DB   *gorm.DB
}

func (s *Server) Start() {

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	app := fiber.New()

	router := NewAppRouter(app, s.DB)
	router.Init()

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}
}
