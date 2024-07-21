package userctrl

import (
	"paywatcher/domain/userdomain"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserController struct {
	Repo userdomain.UserRepository
}

func NewUserController(repo userdomain.UserRepository) *UserController {
	return &UserController{
		Repo: repo,
	}
}

func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
	idString := ctx.Params("id")
	id, err := uuid.Parse(idString)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid id"})
	}

	user, err := c.Repo.GetUserById(id)
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "invalid id"})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "data": user})
}

func (c UserController) Index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!!")
}
