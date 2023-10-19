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

	return m.createdAuth.ID, nil
}

func (m *MockRepository) Update(u *Auth) error {
	m.createdAuth = u

	return nil
}

func (m *MockRepository) Delete(id uuid.UUID) error {
	if m.createdAuth.ID != id {
		return ErrNotFound
	}

	m.createdAuth = nil

	return nil
}

func (m *MockRepository) GetByID(id uuid.UUID) (*Auth, error) {
	if id == uuid.Nil {
		return nil, ErrIDMustBeEntered
	}

	if m.createdAuth.ID != id {
		return nil, ErrNotFound
	}

	return m.createdAuth, nil
}

func (m *MockRepository) GetByUName(username string) (*Auth, error) {
	if username == "" {
		return nil, ErrEmptyUsername
	}

	if m.createdAuth.Username != username {
		return nil, ErrNotFound
	}

	return m.createdAuth, nil
}
