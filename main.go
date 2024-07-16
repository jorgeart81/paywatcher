package main

import (
	"fmt"
	"log"
	"paywatcher/config"
	"paywatcher/database"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config := config.GetConfig()
	env := config.Env

	database.Connect()

	// Start server
	addr := fmt.Sprintf("%s:%d", env.APP_HOST, env.APP_PORT)
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	if err := app.Listen(addr); err != nil {
		log.Fatal(err)
	}

}
