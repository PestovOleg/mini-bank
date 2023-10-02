package mapper

import (
	"github.com/google/uuid"
)

// User data
type User struct {
	ID         uuid.UUID `json:"id" example:"fdee7aae-f79f-4653-8a16-9207e6805b93"`
	Username   string    `json:"username" example:"Ivanych"`
	Email      string    `json:"email" example:"Ivanych@gmail.com"`
	Phone      string    `json:"phone" example:"+7(495)999-99-99"`
	Birthday   string    `json:"birthday" example:"01.01.1999"`
	Name       string    `json:"name" example:"Ivan"`
	LastName   string    `json:"last_name" example:"Ivanov"`
	Patronymic string    `json:"patronymic" example:"Ivanych"`
	IsActive   bool      `json:"is_active" example:"true"`
	CreatedAt  string    `json:"created_at" example:"01.01.1999"`
	UpdatedAt  string    `json:"updated_at" example:"01.01.1999"`
}
