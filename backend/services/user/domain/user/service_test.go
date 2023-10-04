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
			id         uuid.UUID
			Email      string
			Phone      string
			Name       string
			LastName   string
			Patronymic string
			Birthday   time.Time
		}
		out struct {
			uuid uuid.UUID
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
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
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
				Email:      "vasn@mail.ru",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
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
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrIDMustBeEntered},
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
				Email:      "",
				Phone:      "1234567890",
				Name:       "Vasya",
				LastName:   "Vasilev",
				Patronymic: "Vasilich",
				Birthday:   time.Now(),
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyEmail},
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
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyName},
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
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyLastName},
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
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPatronymic},
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
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyBirthday},
		},
	}

	mockRepo := &MockRepository{}

	service := NewService(mockRepo)

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			id, err := service.CreateUser(
				i.in.id,
				i.in.Email,
				i.in.Phone,
				i.in.Name,
				i.in.LastName,
				i.in.Patronymic,
				i.in.Birthday,
			)
			if id != i.out.uuid && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v", id, err, i.out)
			}
		})
	}
}
