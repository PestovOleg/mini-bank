package auth

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestCreateAuthTableDriven(t *testing.T) {
	tests := []struct {
		in struct {
			Username string
			Password string
		}
		out struct {
			uuid uuid.UUID
			err  error
		}
	}{
		{
			in: struct {
				Username string
				Password string
			}{
				Username: "vasyan",
				Password: "vasyaqwerty",
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				Username string
				Password string
			}{
				Username: "Petyan",
				Password: "vasyaqwerty",
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.New(), err: nil},
		},
		{
			in: struct {
				Username string
				Password string
			}{
				Username: "",
				Password: "vasyaqwerty",
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyUsername},
		},
		{
			in: struct {
				Username string
				Password string
			}{
				Username: "vasyan",
				Password: "",
			},
			out: struct {
				uuid uuid.UUID
				err  error
			}{uuid: uuid.Nil, err: ErrEmptyPassword},
		},
	}

	mockRepo := &MockRepository{}

	service := NewService(mockRepo)

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			id, err := service.CreateAuth(
				i.in.Username,
				i.in.Password,
			)
			if id != i.out.uuid && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v", id, err, i.out)
			}
		})
	}
}
