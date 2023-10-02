package auth

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//nolint:gochecknoglobals
var mockUser = &Auth{
	ID:        uuid.New(),
	Username:  "vasyan",
	IsActive:  true,
	Password:  "",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func TestNewAuthTableDriven(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("vasyaqwerty"), 10)
	if err != nil {
		t.Fatalf("cannot generate a hash %v", err)
	}

	mockUser.Password = string(hash)

	tests := []struct {
		in struct {
			Username string
			Password string
		}
		out struct {
			user *Auth
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
				user *Auth
				err  error
			}{
				user: mockUser,
				err:  nil,
			},
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
				user *Auth
				err  error
			}{
				user: mockUser,
				err:  ErrEmptyUsername,
			},
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
				user *Auth
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
			u, err := NewAuth(
				i.in.Username,
				i.in.Password,
			)
			if u != nil && (u.Username != mockUser.Username &&
				u.Password != mockUser.Password &&
				u.CreatedAt.IsZero() &&
				u.UpdatedAt.IsZero() &&
				u.ID == uuid.Nil && err == nil) || (u == nil && !errors.Is(err, i.out.err)) {
				t.Errorf("got %v and %v, wants %v", u, err, i.out)
			}
		})
	}
}

func TestValidateAuthTableDriven(t *testing.T) {
	tests := []struct {
		in  *Auth
		out error
	}{
		{mockUser, nil},
		{&Auth{
			ID:        uuid.New(),
			Username:  "",
			IsActive:  true,
			Password:  "samplePassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, ErrEmptyUsername},
		{&Auth{
			ID:        uuid.New(),
			Username:  "vasyan",
			IsActive:  true,
			Password:  "",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, ErrEmptyPassword},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			err := i.in.ValidateAuth()
			if !errors.Is(err, i.out) {
				t.Errorf("got %v, wants %v", i.in, i.out)
			}
		})
	}
}
