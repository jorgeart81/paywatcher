package schemas

import (
	"paywatcher/src/domain/entity"
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"column:name;not null"`
	Priority  uint      `gorm:"column:priority"`
	Recurrent bool      `gorm:"column:recurrent"`
	Notify    bool      `gorm:"column:notify"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	UserID uuid.UUID `gorm:"type:uuid"`
}

func ToCategorySchema(category *entity.CategoryEnt) *Category {
	return &Category{
		ID:        category.ID,
		Name:      category.Name,
		Priority:  category.Priority,
		Recurrent: category.Recurrent,
		Notify:    category.Notify,
	}
}

func (e *Category) ToDomain() *entity.CategoryEnt {
	return &entity.CategoryEnt{
		ID:        e.ID,
		Name:      e.Name,
		Priority:  e.Priority,
		Recurrent: e.Recurrent,
		Notify:    e.Notify,
	}
}
