package response

import (
	"paywatcher/src/domain/entity"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Priority  uint      `json:"priority"`
	Recurrent bool      `json:"recurrent"`
	Notify    bool      `json:"notify"`
}

func NewCategoryResponse(category *entity.CategoryEnt) CategoryResponse {
	return CategoryResponse{
		ID:        category.ID,
		Name:      category.Name,
		Priority:  category.Priority,
		Recurrent: category.Recurrent,
		Notify:    category.Notify,
	}
}
