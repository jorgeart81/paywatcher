package userctrl

import (
	"net/http"
	"paywatcher/src/application/usecases"
	"paywatcher/src/domain/userdomain"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
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

func (c UserController) Create(ctx *gin.Context) {
	var user userdomain.User
	var gResp response.Generic

	if err := ctx.Bind(&user); err != nil {
		gResp.Message = "invalid request"
		ctx.JSON(http.StatusBadRequest, gResp)
		return
	}

	newUser, err := c.createUC.Execute(user)
	if err != nil {
		gResp.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, gResp.Err())
		return
	}

	gResp.User = response.NewUserResponse(newUser)
	ctx.JSON(http.StatusCreated, gResp.Ok())
}

func (c UserController) Login(ctx *gin.Context) {
	var user userdomain.User
	var gResp response.Generic

	if err := ctx.Bind(&user); err != nil {
		gResp.Message = "invalid request"
		ctx.JSON(http.StatusBadRequest, gResp.Err())
		return
	}

	u, token, err := c.loginUC.Execute(user.Email, user.Password)

	if err != nil {
		gResp.Message = err.Error()
		ctx.JSON(http.StatusBadRequest, gResp.Err())
		return
	}

	gResp.User = response.NewUserResponse(u)
	gResp.Token = token
	ctx.JSON(http.StatusOK, gResp.Ok())
}

// func (c *UserController) GetUserById(ctx *gin.Context) error {
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

func (c UserController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, World!!")
}
