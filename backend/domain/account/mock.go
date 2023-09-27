package account

import (
	"time"

	"github.com/google/uuid"
)

type MockAccountRepository struct {
	accounts map[uuid.UUID]*Account
}

func NewMockAccountRepository() *MockAccountRepository {
	return &MockAccountRepository{
		accounts: make(map[uuid.UUID]*Account),
	}
}

func (m *MockAccountRepository) Create(a *Account) error {
	id := uuid.New()
	a.ID = id
	m.accounts[id] = a

	return nil
}

func (m *MockAccountRepository) Update(a *Account) error {
	if _, exists := m.accounts[a.ID]; !exists {
		return ErrNotFound
	}

	a.UpdatedAt = time.Now()
	m.accounts[a.ID] = a

	return nil
}

func (m *MockAccountRepository) Delete(id uuid.UUID) error {
	if _, exists := m.accounts[id]; !exists {
		return ErrNotFound
	}

	delete(m.accounts, id)

	return nil
}

func (m *MockAccountRepository) GetByID(id uuid.UUID) (*Account, error) {
	account, exists := m.accounts[id]
	if !exists {
		return nil, ErrNotFound
	}

	return account, nil
}

func (m *MockAccountRepository) GetByNumber(acc string) (*Account, error) {
	for _, account := range m.accounts {
		if account.Account == acc {
			return account, nil
		}
	}

	return nil, ErrNotFound
}

func (m *MockAccountRepository) List(userID uuid.UUID) ([]*Account, error) {
	var userAccounts []*Account

	for _, account := range m.accounts {
		if account.UserID == userID {
			userAccounts = append(userAccounts, account)
		}
	}

	return userAccounts, nil
}

func (m *MockAccountRepository) GetLastOpenedAccount(currency string) (string, error) {
	var lastOpenedAccount string

	var latestTime time.Time

	for _, account := range m.accounts {
		if account.Currency == currency && account.CreatedAt.After(latestTime) {
			latestTime = account.CreatedAt
			lastOpenedAccount = account.Account
		}
	}

	if lastOpenedAccount == "" {
		return "", nil
	}

	return lastOpenedAccount, nil
}
