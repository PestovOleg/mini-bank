package user

import (
	"github.com/google/uuid"
)

type MockRepository struct {
	errToReturn  error
	userToReturn *User
	createdUser  *User
}

func (m *MockRepository) Create(u *User) (uuid.UUID, error) {
	m.createdUser = u
	if m.errToReturn != nil {
		return uuid.Nil, m.errToReturn
	}

	return uuid.New(), nil
}

func (m *MockRepository) Update(u *User) error {
	return nil
}

func (m *MockRepository) Delete(u *User) error {
	return nil
}

func (m *MockRepository) GetByID(id uuid.UUID) (*User, error) {
	return nil, nil
}
func (m *MockRepository) GetByUName(username string) (*User, error) {
	return nil, nil
}
func (m *MockRepository) List() ([]*User, error) {
	return nil, nil
}
