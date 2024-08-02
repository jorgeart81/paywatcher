package entity

import (
	"github.com/google/uuid"
)

type CategoryEnt struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Priority  uint      `json:"priority"`
	Recurrent bool      `json:"recurrent"`
	Notify    bool      `json:"notify"`
}

func (c *CategoryEnt) NewCategoryEnt() *CategoryEnt {
	c.ID = uuid.New()
	return c
}
