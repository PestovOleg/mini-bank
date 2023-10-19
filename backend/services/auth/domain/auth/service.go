// Бизнес-логика
package auth

import (
	"encoding/base64"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

// конструктор юзкейса
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// GetUserByID Поиск авторизации по id
func (s *Service) GetAuthByID(id uuid.UUID) (*Auth, error) {
	if id == uuid.Nil {
		return nil, ErrIDMustBeEntered
	}

	u, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	// вкючаем проверку на "неактивность"(isActive) пользователя
	if u.ID == uuid.Nil || !u.IsActive {
		return nil, ErrNotFound
	}

	return u, nil
}

// GetUserByUName Поиск авторизации по username
func (s *Service) GetUserByUName(username string) (*Auth, error) {
	if username == "" {
		return nil, ErrMustBeFilledIn
	}

	u, err := s.repo.GetByUName(username)
	if err != nil {
		return nil, err
	}
	// вкючаем проверку на "неактивность"(isActive) авторизацию
	if u.ID == uuid.Nil || !u.IsActive {
		return nil, ErrNotFound
	}

	return u, nil
}

// CreateAuth Создать авторизацию
func (s *Service) CreateAuth(
	username,
	password string,
) (uuid.UUID, error) {
	u, err := NewAuth(username, password)

	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(u)
}

// DeactivateAuth Удалить авторизацию (установить признак isActive в false)
func (s *Service) DeactivateAuth(id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrIDMustBeEntered
	}

	u, err := s.GetAuthByID(id)
	if err != nil {
		return err
	}
	// вкючаем проверку на "неактивность"(isActive) авторизацию
	if u.ID == uuid.Nil || !u.IsActive {
		return ErrNotFound
	}

	u.IsActive = false
	u.UpdatedAt = time.Now()

	return s.repo.Update(u)
}

// DeleteAuth Удалить авторизацию (физически,для межсервисного взаимодействия)
func (s *Service) DeleteAuth(id uuid.UUID) error {
	if id == uuid.Nil {
		return ErrIDMustBeEntered
	}

	return s.repo.Delete(id)
}

func (s *Service) AuthenticateUser(username, password string) (uuid.UUID, error) {
	if username == "" {
		return uuid.Nil, ErrEmptyUsername
	}

	if password == "" {
		return uuid.Nil, ErrEmptyPassword
	}

	a, err := s.repo.GetByUName(username)
	if err != nil {
		return uuid.Nil, err
	}

	if !a.IsActive {
		return uuid.Nil, ErrNotFound
	}

	err = a.VerifyPassword(password)
	if err != nil {
		return uuid.Nil, ErrAuthentication
	}

	return a.ID, nil
}

func (s *Service) AuthorizeUser(token string) error {
	// Should be in the format "Basic base64credentials"
	parts := strings.SplitN(token, " ", 2)
	if len(parts) != 2 || parts[0] != "Basic" {
		return ErrUnauthorized
	}

	decodedBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return ErrUnauthorized
	}

	credentials := strings.SplitN(string(decodedBytes), ":", 2)
	if len(credentials) != 2 {
		return ErrUnauthorized
	}

	username := credentials[0]
	password := credentials[1]

	a, err := s.repo.GetByUName(username)
	if err != nil {
		return ErrUnauthorized
	}

	if username == "" {
		return ErrEmptyUsername
	}

	if password == "" {
		return ErrEmptyPassword
	}

	if !a.IsActive {
		return ErrNotFound
	}

	err = a.VerifyPassword(password)
	if err != nil {
		return ErrUnauthorized
	}

	return nil
}
