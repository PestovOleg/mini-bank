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

// GetUser Поиск пользователя по id
func (s *Service) GetUser(id uuid.UUID) (*User, error) {
	u, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	if u.ID == uuid.Nil {
		return nil, ErrNotFound
	}

	return u, nil
}

// CreateUser Создать пользователя TODO: переделать на возврат User?
func (s *Service) CreateUser(
	id uuid.UUID,
	email,
	phone,
	name,
	lastName,
	patronymic string,
	birthday time.Time,
) (uuid.UUID, error) {
	u, err := NewUser(id, email, phone, name, lastName, patronymic, birthday)

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

	u, err := s.Get(id)
	if err != nil {
		return err
	}
	// включаем проверку на "неактивность"(isActive) пользователя
	if u.ID == uuid.Nil {
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
