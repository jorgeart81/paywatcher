package controller

import (
	"net/http"
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/config"
	"paywatcher/src/domain/services"
	"paywatcher/src/presentation/request"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService    services.Authenticator
	createUC       *user.RegisterUserUseCase
	loginUC        *user.LoginUserUseCase
	refreshTokenUC *user.RefreshTokenUseCase
}

func newAuthController(authService services.Authenticator, createUserUC user.RegisterUserUseCase, loginUserUC user.LoginUserUseCase, refreshTokenUC user.RefreshTokenUseCase) *AuthController {
	return &AuthController{
		authService:    authService,
		createUC:       &createUserUC,
		loginUC:        &loginUserUC,
		refreshTokenUC: &refreshTokenUC,
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
// @Router /register [post]
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
// @Router /login [post]
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

// @Summary Refresh Token
// @Description Create a new refresh token
// @Tags User
// @Produce json
// @Success 200 {object} response.RefreshTokenResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 401 {object} response.ErrorResponse
// @Router /refresh-token [get]
func (c AuthController) RefreshToken(ctx *gin.Context) {
	cookieName := config.JWT.CookieName
	refreshToken, err := ctx.Cookie(cookieName)
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	tokenPairs, err := c.refreshTokenUC.Execute(refreshToken)
	if err != nil {
		response.SendError(ctx, http.StatusUnauthorized, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	refreshCookie := c.authService.GetRefreshCookie(tokenPairs.RefreshToken)
	http.SetCookie(ctx.Writer, refreshCookie)

	refreshTokenResponse := response.NewRefreshTokenResponse(tokenPairs.AccessToken)
	response.SendSuccess(ctx, http.StatusOK, refreshTokenResponse)

}

// TODO: only for demo
func (c AuthController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, World!!")
}
