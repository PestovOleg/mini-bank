package user

import (
	"time"

	"github.com/google/uuid"
)

// Reader
// Get - информация о записи клиента
type Reader interface {
	Get(id uuid.UUID) (*User, error)
}

// Writer
// Create - создать запись с пользователем
// Update - обновить пользователя новыми значениями
// Delete - удалить пользователя
type Writer interface {
	Create(u *User) (uuid.UUID, error)
	Update(u *User) error
	Delete(id uuid.UUID) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecase интерфейс
// GetUser - поиск пользователя по ID
// CreateUser - создание пользователя
// UpdateUser - обновление данных пользователя
// DeleteUser - удаление пользователя

type UseCase interface {
	GetUser(id uuid.UUID) (*User, error)
	CreateUser(id, email, phone, name, lastName, patronymic string, birthday time.Time) (uuid.UUID, error)
	UpdateUser(id uuid.UUID, email, phone string) error
	DeleteUser(id uuid.UUID) error
}
