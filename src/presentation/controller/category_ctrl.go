package controller

import (
	"net/http"
	"paywatcher/src/application/usecases/category"
	"paywatcher/src/infrastructure/middlewares"
	"paywatcher/src/presentation/request"
	"paywatcher/src/presentation/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CategoryController struct {
	createUC         *category.CreateCategoryUseCase
	userCategoriesUC *category.UserCategoriesUseCase
}

func newCategoryController(createUC category.CreateCategoryUseCase, userCategoriesUC category.UserCategoriesUseCase) *CategoryController {
	return &CategoryController{
		createUC:         &createUC,
		userCategoriesUC: &userCategoriesUC,
	}
}

func (c CategoryController) Create(ctx *gin.Context) {
	var req request.CreateCategoryReq

	if err := ctx.ShouldBind(&req); err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	id, ok := ctx.Value(middlewares.UserIDKey).(uuid.UUID)
	if !ok {
		response.SendError(ctx, http.StatusInternalServerError, &response.GenericError{
			Message: "user not found",
		})
		return
	}

	newCategory, err := c.createUC.Execute(req.ToCategoryEntity(), id)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
		return
	}

	reponse := response.NewCategoryResponse(newCategory)
	response.SendSuccess(ctx, http.StatusCreated, reponse)
}

func (c CategoryController) GetUserCategories(ctx *gin.Context) {
	id, ok := ctx.Value(middlewares.UserIDKey).(uuid.UUID)
	if !ok {
		response.SendError(ctx, http.StatusInternalServerError, &response.GenericError{
			Message: "user not found",
		})
		return
	}

	categories, err := c.userCategoriesUC.Execute(id)
	if err != nil {
		response.SendError(ctx, http.StatusBadRequest, &response.GenericError{
			Message: err.Error(),
		})
	}

	var categoriesReponse []response.CategoryResponse

	for _, category := range *categories {
		categoriesReponse = append(categoriesReponse, response.NewCategoryResponse(&category))
	}

	response.SendSuccess(ctx, http.StatusOK, categoriesReponse)
}
