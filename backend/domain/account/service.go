package account

import (
	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

// Usecase Constructor
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) GetAccountByID(id uuid.UUID) (*Account, error) {
	a, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Service) GetAccountByNumber(account string) (*Account, error) {
	a, err := s.repo.GetByNumber(account)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Service) ListAccount(userID uuid.UUID) ([]*Account, error) {
	acc, err := s.repo.List(userID)
	if err != nil {
		return nil, err
	}

	return acc, nil
}

func (s *Service) CreateAccount(userID uuid.UUID, currency string) (*Account, error) {
	acc, err := s.repo.GetLastOpenedAccount(currency)
	if err != nil {
		return nil, err
	}

	a, err := NewAccount(userID, currency, acc)
	if err != nil {
		return nil, err
	}
	_, err = s.repo.Create(a)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (s *Service) DeleteAccount(id uuid.UUID) error {
	return s.repo.Delete(id)
}

func (s *Service) TopUp(id uuid.UUID, money float64) (float64, error) {
	if money <= 0 {
		return 0, ErrMustBePositiveOrZero
	}

	a, err := s.repo.GetByID(id)
	if err != nil {
		return 0, err
	}

	a.Amount += money

	err = s.repo.Update(a)
	if err != nil {
		return 0, err
	}

	return a.Amount, nil
}

func (s *Service) WithDraw(id uuid.UUID, money float64) (float64, error) {
	if money <= 0 {
		return 0, ErrMustBePositiveOrZero
	}

	a, err := s.repo.GetByID(id)
	if err != nil {
		return 0, err
	}

	a.Amount -= money
	if a.Amount < 0 {
		return 0, ErrNotEnoughMoney
	}

	err = s.repo.Update(a)
	if err != nil {
		return 0, err
	}

	return a.Amount, nil
}
