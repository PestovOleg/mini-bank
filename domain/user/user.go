package user

import (
	"fmt"
	"time"

	"github.com/PestovOleg/mini-bank/domain"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// User entity
type User struct {
	ID         uuid.UUID
	Username   string
	Email      string
	Name       string
	LastName   string
	Patronymic string
	Password   string
	IsActive   bool
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Конструктор
func NewUser(userName, email, name, lastName, patronymic, password string) (*User, error) {
	u := &User{
		ID:         uuid.New(),
		Username:   userName,
		Email:      email,
		Name:       name,
		LastName:   lastName,
		IsActive:   true,
		Patronymic: patronymic,
		CreatedAt:  time.Now(),
	}

	err := u.ValidateUser()
	if err != nil {
		return nil, domain.ErrMustBeFilledIn
	}

	hash, err := generateHash(u.Password)
	if err != nil {
		return nil, err
	}

	u.Password = hash

	return u, nil
}

// TODO: сделать проверки дополнительные
func (u *User) ValidateUser() error {
	if u.Username == "" || u.Email == "" || u.Name == "" || u.LastName == "" || u.Patronymic == "" {
		return domain.ErrMustBeFilledIn
	}

	return nil
}

// генерация хеша к паролю
func generateHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", fmt.Errorf("cannot generate a hash %w", err)
	}

	return string(hash), nil
}

// проверка пароля
func (u *User) CheckPassword(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return fmt.Errorf("password is oncorrect %w", err)
	}

	return nil
}

// Reader
// Get - информация о записи клиента
// List - вернуть все записи
type Reader interface {
	Get(id uuid.UUID) (*User, error)
	List() ([]*User, error)
}

// Writer
// Create - создать запись с пользователем
// Update - обновить пользователя новыми значениями
// Delete - установить признак удаления (пользователь не может быть удален)
type Writer interface {
	Create(u *User) (uuid.UUID, error)
	Update(u *User) error
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
	GetUser(id uuid.UUID) (*User, error)
	ListUsers() ([]*User, error)
	CreateUser(username, email, name, lastName, patronymic, password string) (uuid.UUID, error)
	UpdateUser(u *User) error
	DeleteUser(uuid.UUID) error
}
