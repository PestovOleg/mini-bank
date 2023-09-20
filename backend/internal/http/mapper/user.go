package mapper

import (
	"time"

	"github.com/google/uuid"
)

// User data
type User struct {
	ID         uuid.UUID `json:"id" example:"fdee7aae-f79f-4653-8a16-9207e6805b93"`
	Username   string    `json:"username" example:"Ivanych"`
	Email      string    `json:"email" example:"Ivanych@gmail.com"`
	Name       string    `json:"name" example:"Ivan"`
	LastName   string    `json:"last_name" example:"Ivanov"`
	Patronymic string    `json:"patronymic" example:"Ivanych"`
	IsActive   bool      `json:"is_active" example:"true"`
	CreatedAt  time.Time `json:"created_at" example:"2023-09-19T10:58:00.000Z"`
	UpdatedAt  time.Time `json:"updated_at" example:"2023-09-19T10:58:00.000Z"`
}
