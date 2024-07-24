package presentation

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	Port    int
	Host    string
	GinMode string
	DB      *gorm.DB
}

func (s *Server) Start() {

	addr := fmt.Sprintf("%s:%d", s.Host, s.Port)
	if len(s.GinMode) > 0 {
		gin.SetMode(s.GinMode)
	}
	app := gin.Default()

	router := NewAppRouter(app, s.DB)
	router.Init()

	if err := app.Run(addr); err != nil {
		log.Fatal(err)
	}
}
