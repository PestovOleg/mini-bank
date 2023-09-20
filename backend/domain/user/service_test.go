package user

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreateUserTableDriven(t *testing.T) {
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
			uuid uuid.UUID
			err  error
		}
	}{
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"vasyan", "vasyan@mail.ru", "Vasya", "Vasilev", "Vasilich", "vasyaqwerty"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"petyan", "petyan@mail.ru", "Petya", "Petrov", "Petrovich", "petya123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"", "nil@nil.ru", "Nil", "Nilov", "Nilovich", "nil123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyUsername},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "", "Nil", "Nilov", "Nilovich", "nil123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyEmail},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "", "Nilov", "Nilovich", "nil123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyName},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "", "Nilovich", "nil123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyLastName},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "Nilov", "", "nil123"},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPatronymic},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"nill", "nil@nil.ru", "Nil", "Nilov", "Nilovich", ""},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPassword},
		},
		{
			in: struct {
				username, email, name, lastName, patronymic, password string
			}{
				"", "", "", "", "", ""},
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
			id, err := service.CreateUser(i.in.username, i.in.email, i.in.name, i.in.lastName, i.in.patronymic, i.in.password)
			if id != i.out.uuid && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v", id, err, i.out)
			}
		})
	}
}
