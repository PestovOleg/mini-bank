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

func TestDeactivateUserTableDriven(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewService(mockRepo)

	userID, err := service.CreateAuth("user", "password")
	if err != nil {
		t.Fatal("Cannot create auth record")
	}

	tests := []struct {
		in  uuid.UUID
		out error
	}{
		{uuid.Nil, ErrIDMustBeEntered},
		{uuid.MustParse("0ddb1104-d805-4b1b-b00d-03ca8fa05ea4"), ErrNotFound},
		{userID, nil},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			err := service.DeactivateAuth(i.in)
			if !errors.Is(err, i.out) {
				t.Errorf("got %v, wants %v", err, i.out) // здесь не используется .Error()
			}
		})
	}
}

func TestDeleteUserTableDriven(t *testing.T) {
	mockRepo := &MockRepository{}
	service := NewService(mockRepo)

	userID, err := service.CreateAuth("user", "password")
	if err != nil {
		t.Fatal("Cannot create auth record")
	}

	tests := []struct {
		in  uuid.UUID
		out error
	}{
		{uuid.Nil, ErrIDMustBeEntered},
		{uuid.MustParse("0ddb1104-d805-4b1b-b00d-03ca8fa05ea4"), ErrNotFound},
		{userID, nil},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			err := service.DeleteAuth(i.in)
			if !errors.Is(err, i.out) {
				t.Errorf("got %v, wants %v", err, i.out) // здесь не используется .Error()
			}
		})
	}
}
