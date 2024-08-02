package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `gorm:"column:name;unique"`
	Priority  uint      `gorm:"column:priority"`
	Recurrent bool      `gorm:"column:recurrent"`
	Notify    bool      `gorm:"column:notify"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
