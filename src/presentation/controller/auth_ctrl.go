package controller

import (
	"net/http"
	"paywatcher/src/application/usecases"
	"paywatcher/src/domain/services"
	"paywatcher/src/presentation/request"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService services.Authenticator
	createUC    *usecases.CreateUserUseCase
	loginUC     *usecases.LoginUserUseCase
}

func newAuthController(authService services.Authenticator, createUserUC usecases.CreateUserUseCase, loginUserUC usecases.LoginUserUseCase) *AuthController {
	return &AuthController{
		authService: authService,
		createUC:    &createUserUC,
		loginUC:     &loginUserUC,
	}
}

// @BasePath /api

// @Summary Register user
// @Description Register a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.RegisterUser true "Request body"
// @Success 201 {object} response.AuthResponse
// @Failure 400 {object} response.ErrorResponse
// @Router /user/register [post]
func (c AuthController) Create(ctx *gin.Context) {
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

	newUser, tokenPairs, err := c.createUC.Execute(req.ToUserEntity())
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	refreshCookie := c.authService.GetRefreshCookie(tokenPairs.RefreshToken)
	http.SetCookie(ctx.Writer, refreshCookie)

	authResponse := response.NewAuthResponse(newUser, tokenPairs.AccessToken)
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
func (c AuthController) Login(ctx *gin.Context) {
	var req *request.LoginUser

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	user, tokenPairs, err := c.loginUC.Execute(req.Email, req.Password)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	refreshCookie := c.authService.GetRefreshCookie(tokenPairs.RefreshToken)
	http.SetCookie(ctx.Writer, refreshCookie)

	authResponse := response.NewAuthResponse(user, tokenPairs.AccessToken)
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

func (c AuthController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, World!!")
}
