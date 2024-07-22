package userctrl

import (
	"paywatcher/src/application/usecases"
	"paywatcher/src/domain/userdomain"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	createUserUC *usecases.CreateUserUseCase
}

func NewUserController(createUserUC usecases.CreateUserUseCase) *UserController {
	return &UserController{
		createUserUC: &createUserUC,
	}
}

func (c UserController) Create(ctx *fiber.Ctx) error {
	var user userdomain.User
	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid request"})
	}

	newUser, err := c.createUserUC.Execute(user)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": err.Error()})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "ok", "data": newUser})
}

// func (c *UserController) GetUserById(ctx *fiber.Ctx) error {
// 	idString := ctx.Params("id")
// 	id, err := uuid.Parse(idString)

// 	if err != nil {
// 		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "invalid id"})
// 	}

// 	user, err := c.Repo.GetUserById(id)
// 	if err != nil {
// 		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "invalid id"})
// 	}

// 	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{"status": "ok", "data": user})
// }

func (c UserController) Index(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!!")
}
