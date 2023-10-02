package user

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateUserTableDriven(t *testing.T) {
	tests := []struct {
		in struct {
			Username   string
			Email      string
			Phone      string
			Name       string
			LastName   string
			Patronymic string
			Password   string
			Birthday   time.Time
		}
		out struct {
			uuid uuid.UUID
			err  error
		}
	}{
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "Petyan",
				Email:      "vasn@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyUsername},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyEmail},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyName},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyLastName},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPatronymic},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "vasyan",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPassword},
		},
		{
			in: struct {
				Username   string
				Email      string
				Phone      string
				Name       string
				LastName   string
				Patronymic string
				Password   string
				Birthday   time.Time
			}{
				Username:   "",
				Email:      "vasyan@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Password:   "vasyaqwerty",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyUsername},
		},
	}

	mockRepo := &MockRepository{}

	service := NewService(mockRepo)

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			id, err := service.CreateUser(
				i.in.Username,
				i.in.Email,
				i.in.Phone,
				i.in.Name,
				i.in.LastName,
				i.in.Patronymic,
				i.in.Password,
				i.in.Birthday,
			)
			if id != i.out.uuid && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v", id, err, i.out)
			}
		})
	}
}
