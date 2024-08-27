package request

import (
	"paywatcher/src/domain/entity"
)

type CreateCategoryReq struct {
	Name      string `form:"name" json:"name" binding:"required"`
	Priority  uint   `form:"priority" json:"priority"`
	Recurrent bool   `form:"recurrent" json:"recurrent"`
	Notify    bool   `form:"notify" json:"notify"`
}

func (c *CreateCategoryReq) ToCategoryEntity() *entity.CategoryEnt {
	return &entity.CategoryEnt{
		Name:      c.Name,
		Priority:  c.Priority,
		Recurrent: c.Recurrent,
		Notify:    c.Notify,
	}
}
