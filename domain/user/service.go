// Бизнес-логика
package user

import (
	"time"

	"github.com/PestovOleg/mini-bank/domain"
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

// GetUser Поиск пользователя
func (s *Service) GetUser(id uuid.UUID) (*User, error) {
	return s.repo.Get(id)
}

// ListUsers Список пользователей
func (s *Service) ListUsers() ([]*User, error) {
	return s.repo.List()
}

// CreateUser Создать пользователя
func (s *Service) CreateUser(username, email, name, lastName, patronymic, password string) (uuid.UUID, error) {
	u, err := NewUser(username, email, name, lastName, patronymic, password)
	if err != nil {
		return u.ID, err
	}

	return s.repo.Create(u)
}

// UpdateUser Обновить пользователя
func (s *Service) UpdateUser(u *User) error {
	err := u.ValidateUser()
	if err != nil {
		return err
	}

	return s.repo.Update(u)
}

// DeleteUser Удалить пользователя (установить признак Active в false)
func (s *Service) DeleteUser(id uuid.UUID) error {
	u, err := s.GetUser(id)
	if err == nil {
		return domain.ErrNotFound
	}

	if err != nil {
		return err
	}

	u.UpdatedAt = time.Now()

	return s.repo.Update(u)
}
