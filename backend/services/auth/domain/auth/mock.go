package auth

import (
	"github.com/google/uuid"
)

type MockRepository struct {
	errToReturn error
	createdAuth *Auth
}

func (m *MockRepository) Create(u *Auth) (uuid.UUID, error) {
	m.createdAuth = u
	if m.errToReturn != nil {
		return uuid.Nil, m.errToReturn
	}

	return uuid.New(), nil
}

func (m *MockRepository) Update(u *Auth) error {
	return nil
}

func (m *MockRepository) Delete(u *Auth) error {
	return nil
}

func (m *MockRepository) GetByID(id uuid.UUID) (*Auth, error) {
	return nil, nil
}
func (m *MockRepository) GetByUName(username string) (*Auth, error) {
	return nil, nil
}
