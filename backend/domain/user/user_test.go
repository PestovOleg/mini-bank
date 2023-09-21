package user

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var mockUser = &User{
	ID:         uuid.New(),
	Username:   "vasyan",
	Email:      "vasyan@mail.ru",
	Name:       "Vasya",
	LastName:   "Vasilev",
	IsActive:   true,
	Patronymic: "Vasilich",
	Password:   "",
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
}

func TestNewUserTableDriven(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("vasyaqwerty"), 10)
	if err != nil {
		t.Fatalf("cannot generate a hash %v", err)
	}
	mockUser.Password = string(hash)

	tests := []struct {
		in struct {
			username   string
			email      string
			name       string
			lastName   string
			patronymic string
			password   string
		}
		out struct {
			user *User
			err  error
		}
	}{
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"vasyan", "vasyan@mail.ru", "Vasya", "Vasilev", "Vasilich", "vasyaqwerty"},
			out: struct {
				user *User
				err  error
			}{user: mockUser, err: nil},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"", "nil@nil.ru", "Nil", "Nilov", "Nilovich", "nil123"},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyUsername},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "", "Nil", "Nilov", "Nilovich", "nil123"},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyEmail},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "", "Nilov", "Nilovich", "nil123"},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyName},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "", "Nilovich", "nil123"},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyLastName},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "Nilov", "", "nil123"},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyPatronymic},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "Nilov", "Nilovich", ""},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyPassword},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"", "", "", "", "", ""},
			out: struct {
				user *User
				err  error
			}{user: nil, err: ErrEmptyUsername},
		},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			u, err := NewUser(i.in.username, i.in.email, i.in.name, i.in.lastName, i.in.patronymic, i.in.password)
			if u != nil && (u.Username != mockUser.Username &&
				u.Email != mockUser.Email &&
				u.Name != mockUser.Name &&
				u.LastName != mockUser.LastName &&
				u.Patronymic != mockUser.Patronymic &&
				u.Password != mockUser.Password &&
				u.CreatedAt.IsZero() &&
				u.UpdatedAt.IsZero() &&
				u.ID == uuid.Nil && err == nil) || (u == nil && !errors.Is(err, i.out.err)) {
				t.Errorf("got %v and %v, wants %v", u, err, i.out)
			}
		})
	}
}

func TestValidateUserTableDriven(t *testing.T) {
	tests := []struct {
		in  *User
		out error
	}{
		{mockUser, nil},
		{&User{
			ID:         uuid.New(),
			Username:   "",
			Email:      "vasyan@mail.ru",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyUsername},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyEmail},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "vasyan@mail.ru",
			Name:       "",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyName},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "vasyan@mail.ru",
			Name:       "Vasya",
			LastName:   "",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyLastName},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "vasyan@mail.ru",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyPatronymic},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "vasyan@mail.ru",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}, ErrEmptyPassword},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			err := i.in.ValidateUser()
			if !errors.Is(err, i.out) {
				t.Errorf("got %v, wants %v", i.in, i.out)
			}
		})
	}
}
