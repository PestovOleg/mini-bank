// Интерфейсы:
// Reader - интерфейс для чтения
// Writer - интерфейс записи
// Repository - репозиторий
// Usecase - юхкейс,сервис
package user

import (
	"github.com/PestovOleg/mini-bank/entity"
	"github.com/google/uuid"
)

// Reader
// Get - информация о записи клиента
// List - вернуть все записи
type Reader interface {
	Get(id uuid.UUID) (*entity.User, error)
	List() ([]*entity.User, error)
}

// Writer
// Create - создать запись с пользователем
// Update - обновить пользователя новыми значениями
// Delete - установить признак удаления (пользователь не может быть удален)
type Writer interface {
	Create(u *entity.User) (uuid.UUID, error)
	Update(u *entity.User) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecase интерфейс
// GetUser - поиск пользователя
// ListUsers - список пользователей
// CreateUser - создание пользователя
// UpdateUser - обновление данных пользователя

type UseCase interface {
	GetUser(id uuid.UUID) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	CreateUser(username, email, name, lastName, patronymic, password string) (uuid.UUID, error)
	UpdateUser(u *entity.User) error
	DeleteUser(uuid.UUID) error
}
