package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Authentication entity
type Auth struct {
	ID        uuid.UUID
	Username  string
	Password  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Constructor
func NewAuth(username, password string) (*Auth, error) {
	u := &Auth{
		ID:        uuid.New(),
		Username:  username,
		IsActive:  true,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := u.ValidateAuth()
	if err != nil {
		return nil, err
	}

	hash, err := generateHash(u.Password)
	if err != nil {
		return nil, err
	}

	u.Password = hash

	return u, nil
}

func (u *Auth) ValidateAuth() error {
	if u.Username == "" {
		return ErrEmptyUsername
	}

	if u.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

// генерация хеша, cost=10
func generateHash(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return "", fmt.Errorf("cannot generate a hash %w", err)
	}

	return string(hash), nil
}

// проверка хеша пароля
func (u *Auth) VerifyPassword(pwd string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pwd))
	if err != nil {
		return fmt.Errorf("password is incorrect %w", err)
	}

	return nil
}
