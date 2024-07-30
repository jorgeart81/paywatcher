package controller

import (
	"net/http"
	"paywatcher/src/application/usecases/user"
	"paywatcher/src/config"
	"paywatcher/src/presentation/request"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AuthController struct {
	createUC         *user.RegisterUserUseCase
	loginUC          *user.LoginUserUseCase
	refreshTokenUC   *user.RefreshTokenUseCase
	changePasswordUC *user.ChangePasswordUseCase
}

func newAuthController(createUserUC user.RegisterUserUseCase, loginUserUC user.LoginUserUseCase, refreshTokenUC user.RefreshTokenUseCase, changePasswordUC user.ChangePasswordUseCase) *AuthController {
	return &AuthController{
		createUC:         &createUserUC,
		loginUC:          &loginUserUC,
		refreshTokenUC:   &refreshTokenUC,
		changePasswordUC: &changePasswordUC,
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
func (c AuthController) Register(ctx *gin.Context) {
	var req request.RegisterUser

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	if err := req.ValidatePassword(); err != nil {
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

	refreshCookie := config.GetRefreshCookie(tokenPairs.RefreshToken)
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

	refreshCookie := config.GetRefreshCookie(tokenPairs.RefreshToken)
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

	refreshCookie := config.GetRefreshCookie(tokenPairs.RefreshToken)
	http.SetCookie(ctx.Writer, refreshCookie)

	refreshTokenResponse := response.NewRefreshTokenResponse(tokenPairs.AccessToken)
	response.SendSuccess(ctx, http.StatusOK, refreshTokenResponse)

}

// @Summary Change Password
// @Description Change a new password
// @Tags User
// @Accept json
// @Produce json
// @Param request body request.ChangePassword true "Request body"
// @Success 200 {object} response.UpdateUserResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /change-password [post]
func (c AuthController) ChangePassword(ctx *gin.Context) {
	var req request.ChangePassword

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	if err := req.ValidatePassword(); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	id, ok := ctx.Value("ID").(uuid.UUID)
	if !ok {
		response.SendError(ctx, http.StatusInternalServerError, &response.GenericError{
			Message: "id not found",
		})
		return
	}

	user, err := c.changePasswordUC.Execute(id, req.CurrentPassword, req.NewPassword)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	updateUserResponse := response.NewUpdateUserResponse(user)
	response.SendSuccess(ctx, http.StatusOK, updateUserResponse)
}

// TODO: only for demo
func (c AuthController) Index(ctx *gin.Context) {
	ctx.String(200, "Hello, World!!")
}
