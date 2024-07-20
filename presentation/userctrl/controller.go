package userctrl

import (
	"paywatcher/domain/userdomain"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Repo userdomain.UserRepository
}

func NewUserController(repo userdomain.UserRepository) *UserController {
	return &UserController{
		Repo: repo,
	}
}

// func (c *UserController) GetUser(ctx *fiber.Ctx) error {
// 	id := ctx.Params("id")

// }

func (c UserController) Index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!!")
}
