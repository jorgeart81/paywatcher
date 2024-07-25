package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name      string    `json:"name"`
	Priority  uint      `json:"priority"`
	Recurrent bool      `json:"recurrent"`
	Notify    bool      `json:"notify"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
