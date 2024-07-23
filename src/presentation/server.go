package presentation

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Port int
	Host string
	DB   *gorm.DB
}

func (s *Server) Start() {

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	app := gin.Default()

	router := NewAppRouter(app, s.DB)
	router.Init()

	if err := app.Run(addr); err != nil {
		log.Fatal(err)
	}
}
