// Бизнес-логика
package user

import (
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

// GetUserByID Поиск пользователя по id
func (s *Service) GetUserByID(id uuid.UUID) (*User, error) {
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

// GetUserByUName Поиск пользователя по username
func (s *Service) GetUserByUName(username string) (*User, error) {
	u, err := s.repo.GetByUName(username)
	if err != nil {
		return nil, err
	}
	// вкючаем проверку на "неактивность"(isActive) пользователя
	if u.ID == uuid.Nil || !u.IsActive {
		return nil, ErrNotFound
	}

	return u, nil
}

// CreateUser Создать пользователя TODO: переделать на возврат User?
func (s *Service) CreateUser(
	username,
	email,
	phone,
	name,
	lastName,
	patronymic,
	password string,
	birthday time.Time,
) (uuid.UUID, error) {
	u, err := NewUser(username, email, phone, name, lastName, patronymic, password, birthday)

	if err != nil {
		return uuid.Nil, err
	}

	return s.repo.Create(u)
}

// UpdateUser Обновить пользователя
func (s *Service) UpdateUser(id uuid.UUID, email, phone string) error {
	if id == uuid.Nil {
		return ErrIDMustBeEntered
	}

	if email == "" && phone == "" {
		return ErrMustBeFilledIn
	}

	u, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	// вкючаем проверку на "неактивность"(isActive) пользователя
	if u.ID == uuid.Nil || !u.IsActive {
		return ErrNotFound
	}

	if email != "" {
		u.Email = email
	}

	if phone != "" {
		u.Phone = phone
	}

	u.UpdatedAt = time.Now()

	return s.repo.Update(u)
}

// DeleteUser Удалить пользователя (установить признак isActive в false)
func (s *Service) DeleteUser(id uuid.UUID) error {
	u, err := s.GetUserByID(id)
	if err != nil {
		return err
	}
	// вкючаем проверку на "неактивность"(isActive) пользователя
	if u.ID == uuid.Nil || !u.IsActive {
		return ErrNotFound
	}

	u.IsActive = false
	u.UpdatedAt = time.Now()

	return s.repo.Delete(u)
}
