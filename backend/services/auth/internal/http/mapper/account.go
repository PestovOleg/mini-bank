package mapper

import (
	"time"

	"github.com/google/uuid"
)

// Account data
type Account struct {
	ID           uuid.UUID `json:"id" example:"fdee7aae-f79f-4653-8a16-9207e6805b93"`
	UserID       uuid.UUID `json:"user_id" example:"fdee7aae-f79f-4653-8a16-9207e6805b93"`
	Account      string    `json:"account" example:"40817810902007654321"`
	Currency     string    `json:"currency" example:"810"`
	Name         string    `json:"name" example:"Удачный"`
	Amount       float64   `json:"amount" example:"99999.99"`
	InterestRate float64   `json:"interest_rate" example:"0.1250"`
	IsActive     bool      `json:"is_active" example:"true"`
	CreatedAt    time.Time `json:"created_at" example:"2023-09-19T10:58:00.000Z"`
	UpdatedAt    time.Time `json:"updated_at" example:"2023-09-19T10:58:00.000Z"`
}
