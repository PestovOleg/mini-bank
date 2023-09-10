package entity

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
		CreatedAt:  time.Now(),
	}

	err := u.ValidateUser()
	if err != nil {
		return nil, ErrMustBeFilledIn
	}

	hash, err := generateHash(u.Password)
	if err != nil {
		return nil, err
	}

	u.Password = hash

	return u, nil
}

// TODO: сделать проверки дополнительные(regex)
func (u *User) ValidateUser() error {
	if u.Username == "" || u.Email == "" || u.Name == "" || u.LastName == "" || u.Patronymic == "" {
		return ErrMustBeFilledIn
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
