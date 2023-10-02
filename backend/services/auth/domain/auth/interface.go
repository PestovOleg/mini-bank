package auth

import "github.com/google/uuid"

// Reader
// Get - информация о записи клиента
// List - вернуть все записи
type Reader interface {
	GetByID(id uuid.UUID) (*Auth, error)
	GetByUName(username string) (*Auth, error)
}

// Writer
// Create - создать запись
// Update - обновить запись новыми значениями
// Delete - удалить запись
type Writer interface {
	Create(u *Auth) (uuid.UUID, error)
	Delete(u *Auth) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecase интерфейс
// GetAuthByID - поиск пользователя по ID
// GetAuthByUName - поиск пользователя по username
// CreateAuth - создание записи
// DeleteAuth - удаление записи(деактивация)
// AuthenticateUser - аутентификация пользователя по username,password
// AuthorizeUser - авторизация пользователя для доступа к сервису

type UseCase interface {
	GetAuthByID(id uuid.UUID) (*Auth, error)
	GetAuthByUName(username string) (*Auth, error)
	AuthenticateUser(username, password string) (uuid.UUID, error)
	AuthorizeUser(token string) error
	CreateAuth(username, password string) (uuid.UUID, error)
	DeleteAuth(id uuid.UUID) error
}
