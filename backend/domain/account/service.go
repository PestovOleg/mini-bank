package account

import "github.com/google/uuid"

type Service struct {
	repo Repository
}

// Usecase Constructor
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
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
