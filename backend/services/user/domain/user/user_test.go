package user

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

//nolint:gochecknoglobals
var mockUser = &User{
	ID:         uuid.New(),
	Email:      "vasyan@mail.ru",
	Phone:      "1234567890",
	Name:       "Vasya",
	LastName:   "Vasilev",
	Patronymic: "Vasilich",
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
	Birthday:   time.Now(),
}

func TestNewUserTableDriven(t *testing.T) {
	tests := []struct {
		in struct {
			id         uuid.UUID
			Email      string
			Phone      string
			Name       string
			LastName   string
			Patronymic string
			Birthday   time.Time
		}
		out struct {
			user *User
			err  error
		}
	}{
		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  nil,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.Nil,
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyUsername,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyEmail,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyPhone,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyName,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyLastName,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyPatronymic,
			},
		},

		{
			in: struct {
				id         uuid.UUID
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Birthday   time.Time
			}{
				id:         uuid.New(),
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				user *User
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyPassword,
			},
		},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			u, err := NewUser(
				i.in.id,
				i.in.Email,
				i.in.Phone,
				i.in.Name,
				i.in.LastName,
				i.in.Patronymic,
				i.in.Birthday,
			)
			if u != nil && (u.Email != mockUser.Email &&
				u.Phone != mockUser.Phone &&
				u.Name != mockUser.Name &&
				u.LastName != mockUser.LastName &&
				u.Patronymic != mockUser.Patronymic &&
				u.Birthday != mockUser.Birthday &&
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
			Phone:      "1234567890",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Birthday:   time.Now(),
		}, ErrEmptyUsername},
		{&User{
			ID:         uuid.New(),
			Username:   "vasyan",
			Email:      "",
			Phone:      "1234567890",
			Name:       "Vasya",
			LastName:   "Vasilev",
			IsActive:   true,
			Patronymic: "Vasilich",
			Password:   "samplePassword",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Birthday:   time.Now(),
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
			Phone:      "1234567890",
			Birthday:   time.Now(),
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
			Phone:      "1234567890",
			Birthday:   time.Now(),
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
			Phone:      "1234567890",
			Birthday:   time.Now(),
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
			Phone:      "1234567890",
			Birthday:   time.Now(),
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
