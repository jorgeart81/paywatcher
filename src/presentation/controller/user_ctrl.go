package controller

import (
	"net/http"
	"paywatcher/src/application/usecases"
	"paywatcher/src/presentation/request"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUC *usecases.CreateUserUseCase
	loginUC  *usecases.LoginUserUseCase
}

func newUserController(createUserUC usecases.CreateUserUseCase, loginUserUC usecases.LoginUserUseCase) *UserController {
	return &UserController{
		createUC: &createUserUC,
		loginUC:  &loginUserUC,
	}
}

// @BasePath /api

// @Summary Create user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.RegisterUser true "Request body"
// @Success 201 {object} response.AuthResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/create [post]
func (c UserController) Create(ctx *gin.Context) {
	var req request.RegisterUser

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	if err := req.ValidateRoles(); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	newUser, token, err := c.createUC.Execute(req.ToUserEntity())
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	authResponse := response.NewAuthResponse(newUser, token)
	response.SendSuccess(ctx, http.StatusCreated, authResponse)
}

// @Summary Login user
// @Description User login with email and password
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.LoginUser true "Request body"
// @Success 200 {object} response.AuthResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/login [post]
func (c UserController) Login(ctx *gin.Context) {
	var req *request.LoginUser

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	user, token, err := c.loginUC.Execute(req.Email, req.Password)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	authResponse := response.NewAuthResponse(user, token)
	response.SendSuccess(ctx, http.StatusOK, authResponse)
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
