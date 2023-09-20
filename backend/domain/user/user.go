package user

import (
	"fmt"
	"time"

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
		Password:   password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := u.ValidateUser()
	if err != nil {
		return nil, err
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
	if u.Username == "" {
		return ErrEmptyUsername
	}

	if u.Email == "" {
		return ErrEmptyEmail
	}

	if u.Name == "" {
		return ErrEmptyName
	}

	if u.LastName == "" {
		return ErrEmptyLastName
	}

	if u.Patronymic == "" {
		return ErrEmptyPatronymic
	}

	if u.Password == "" {
		return ErrEmptyPassword
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
func (u *User) VerifyPassword(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return fmt.Errorf("password is incorrect %w", err)
	}

	return nil
}

// Reader
// Get - информация о записи клиента
// List - вернуть все записи
type Reader interface {
	GetByID(id uuid.UUID) (*User, error)
	GetByUName(username string) (*User, error)
	List() ([]*User, error)
}

// Writer
// Create - создать запись с пользователем
// Update - обновить пользователя новыми значениями
// Delete - удалить пользователя
type Writer interface {
	Create(u *User) (uuid.UUID, error)
	Update(u *User) error
	Delete(u *User) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecase интерфейс
// GetUserByID - поиск пользователя по ID
// GetUserByUName - поиск пользователя по username
// ListUsers - список пользователей
// CreateUser - создание пользователя
// UpdateUser - обновление данных пользователя
// DeleteUser - удаление пользователя

type UseCase interface {
	GetUserByID(id uuid.UUID) (*User, error)
	GetUserByUName(username string) (*User, error)
	ListUsers() ([]*User, error)
	CreateUser(username, email, name, lastName, patronymic, password string) (uuid.UUID, error)
	UpdateUser(id uuid.UUID, email, name, lastName, patronymic string) error
	DeleteUser(id uuid.UUID) error
}
