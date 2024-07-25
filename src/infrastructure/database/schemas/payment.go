package schemas

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID          uuid.UUID `json:"id" gorm:"primaryKey;type:uuid"`
	Name        string    `json:"name"`
	CategoryID  uuid.UUID `json:"categoryId"`
	NetAmount   float64   `json:"netAmount"`
	GrossAmount float64   `json:"grossAmount"`
	Deductible  float64   `json:"deductible"`
	ChargeDate  string    `json:"chargeDate"`
	Recurrent   bool      `json:"recurrent"`
	PaymentType string    `json:"paymentType"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
}
