package userctrl

import (
	"paywatcher/src/application/usecases"
	"paywatcher/src/domain/userdomain"
	"paywatcher/src/presentation/response"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	createUC *usecases.CreateUserUseCase
	loginUC  *usecases.LoginUserUseCase
}

func NewUserController(createUserUC usecases.CreateUserUseCase, loginUserUC usecases.LoginUserUseCase) *UserController {
	return &UserController{
		createUC: &createUserUC,
		loginUC:  &loginUserUC,
	}
}

func (c UserController) Create(ctx *fiber.Ctx) error {
	var user userdomain.User
	var resp response.Generic

	if err := ctx.BodyParser(&user); err != nil {
		resp.Message = "invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(resp.Err())
	}

	newUser, err := c.createUC.Execute(user)
	if err != nil {
		resp.Message = err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(resp.Err())
	}

	resp.Data = newUser
	return ctx.Status(fiber.StatusCreated).JSON(resp.Ok())
}

func (c UserController) Login(ctx *fiber.Ctx) error {
	var user userdomain.User
	var resp response.Generic

	if err := ctx.BodyParser(&user); err != nil {
		resp.Message = "invalid request"
		return ctx.Status(fiber.StatusBadRequest).JSON(resp.Err())
	}

	u, token, err := c.loginUC.Execute(user.Email, user.Password)
	if err != nil {
		resp.Message = err.Error()
		return ctx.Status(fiber.StatusBadRequest).JSON(resp.Err())
	}

	resp.Data = u
	resp.Token = token
	return ctx.Status(fiber.StatusOK).JSON(resp.Ok())
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
