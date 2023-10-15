package account

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccountTableDriven(t *testing.T) {
	tests := []struct {
		userID   uuid.UUID
		currency string
		name     string
		out      struct {
			a   *Account
			err error
		}
	}{
		{
			userID:   uuid.MustParse("35ec1e95-87d3-42fe-8158-bf59c78b9e26"),
			currency: "810",
			name:     "Удачный",
			out: struct {
				a   *Account
				err error
			}{
				a: &Account{
					UserID:       uuid.MustParse("35ec1e95-87d3-42fe-8158-bf59c78b9e26"),
					Account:      "ACC123456",
					Currency:     "810",
					Amount:       1000.00,
					InterestRate: 0.02,
					IsActive:     true,
				},
				err: nil,
			},
		},
	}

	mockRepo := NewMockAccountRepository()

	service := NewService(mockRepo)

	for _, i := range tests {
		testname := fmt.Sprintf("input userID: %v, currency: %s wants out: %v, %v", i.userID, i.currency, i.out.a, i.out.err)
		t.Run(testname, func(t *testing.T) {
			acc, err := service.CreateAccount(i.userID, i.currency, i.name)
			if acc.ID != i.out.a.ID && !errors.Is(err, i.out.err) {
				t.Errorf("got %v and %v, wants %v and %v", acc.ID, err, i.out.a.ID, i.out.err)
			}
		})
	}
}

func TestGetAccountByID(t *testing.T) {
	repo := NewMockAccountRepository()
	service := NewService(repo)

	// Создаем тестовый аккаунт и добавляем его в мок репозитория
	account, _ := service.CreateAccount(uuid.New(), "810", "test")

	// Вызываем функцию GetAccountByID
	resultAccount, err := service.GetAccountByID(account.ID)

	// Проверяем, что функция вернула ожидаемый аккаунт и нет ошибок
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if resultAccount == nil {
		t.Fatalf("Expected account not to be nil")
	}

	if resultAccount != account {
		t.Fatalf("Expected account ID to match, but got %v", resultAccount.ID)
	}
}

func TestGetAccountByIDAndUserID(t *testing.T) {
	// Создаем мок репозитория
	repo := NewMockAccountRepository()
	service := NewService(repo)

	// Создаем тестовый аккаунт и добавляем его в мок репозитория
	userID := uuid.New()
	account, _ := service.CreateAccount(userID, "810", "test")

	// Вызываем функцию GetAccountByIDAndUserID с верным userID
	resultAccount, err := service.GetAccountByIDAndUserID(account.ID, userID)

	// Проверяем, что функция вернула ожидаемый аккаунт и нет ошибок
	assert.NoError(t, err)
	assert.Equal(t, account, resultAccount)

	// Вызываем функцию GetAccountByIDAndUserID с неверным userID
	otherUserID := uuid.New()
	_, err = service.GetAccountByIDAndUserID(account.ID, otherUserID)

	// Проверяем, что функция вернула ошибку ErrNotFound
	assert.EqualError(t, err, ErrNotFound.Error())
}

func TestListAccount(t *testing.T) {
	// Создаем мок репозитория
	repo := NewMockAccountRepository()
	service := NewService(repo)
	userID := uuid.New()

	account1, _ := service.CreateAccount(userID, "810", "test1")
	account2, _ := service.CreateAccount(userID, "840", "test2")
	account3, _ := service.CreateAccount(userID, "810", "test3")

	// Вызываем функцию ListAccount
	resultAccounts, err := service.ListAccount(userID)
	assert.NoError(t, err)
	assert.ElementsMatch(t, []*Account{account1, account2, account3}, resultAccounts)
}

func TestTopUpAccount(t *testing.T) {
	repo := NewMockAccountRepository()
	service := NewService(repo)
	resultAmount := 100.34

	// Создаем тестовый аккаунт и добавляем его в мок репозитория
	account, _ := service.CreateAccount(uuid.New(), "810", "test")

	// Вызываем функцию TopUp
	resultAccount, err := service.TopUp(account.ID, resultAmount)
	if err != nil {
		t.Fatalf("Error occurred: %v", err)
	}

	if resultAmount != resultAccount {
		t.Fatalf("Expected amount to match, but got %v", account.Amount)
	}
}
