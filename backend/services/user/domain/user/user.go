package user

import (
	"time"

	"github.com/google/uuid"
)

// User entity
type User struct {
	ID         uuid.UUID
	Email      string
	Phone      string
	Birthday   time.Time
	Name       string
	LastName   string
	Patronymic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Конструктор
func NewUser(id uuid.UUID, email, phone, name, lastName, patronymic string, birthday time.Time) (*User, error) {
	u := &User{
		ID:         id,
		Email:      email,
		Phone:      phone,
		Birthday:   birthday,
		Name:       name,
		LastName:   lastName,
		Patronymic: patronymic,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err := u.ValidateUser()
	if err != nil {
		return nil, err
	}

	return u, nil
}

// TODO: сделать проверки дополнительные
func (u *User) ValidateUser() error {
	if u.ID == uuid.Nil {
		return ErrIDMustBeEntered
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

	if u.Birthday.IsZero() {
		return ErrEmptyBirthday
	}

	if u.Phone == "" {
		return ErrEmptyPhone
	}

	return nil
}
