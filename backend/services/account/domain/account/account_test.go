package account

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestNewAccountTableDriven(t *testing.T) {
	tests := []struct {
		in struct {
			userID   uuid.UUID
			currency string
			account  string
			name     string
		}
		out struct {
			account string
			err     error
		}
	}{
		{
			in: struct {
				userID   uuid.UUID
				currency string
				account  string
				name     string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")),
				currency: "810",
				account:  "40817810902000000000",
				name:     "так себе счет"},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000001", err: nil},
		},
		{
			in: struct {
				userID   uuid.UUID
				currency string
				account  string
				name     string
			}{userID: uuid.Must(uuid.Parse("99371a5f-ad18-4d75-8e4f-ce1e4cbd6dac")),
				currency: "810",
				account:  "",
				name:     "так себе счет"},
			out: struct {
				account string
				err     error
			}{account: "40817810902000000001", err: nil},
		},
	}

	for _, i := range tests {
		testname := fmt.Sprintf("input in: %v wants out: %v", i.in, i.out)
		t.Run(testname, func(t *testing.T) {
			a, err := NewAccount(i.in.userID, i.in.currency, i.in.account, i.in.name)
			if a.Account != i.out.account || err != nil {
				t.Errorf("got %v and %v, wants %v", a, err, i.out)
			}
		})
	}
}

func TestNextAccount(t *testing.T) {
	// Тест случая, когда account пустой
	currency := "810"
	account := ""
	expected := "40817810902000000001"
	result, err := NextAccount(currency, account)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}

	// Тест случая, когда account не пустой
	currency = "810"
	account = "40817810920000000001"
	expected = "40817810902000000002"
	result, err = NextAccount(currency, account)

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestValidate(t *testing.T) {
	// Тест случая с правильными данными
	a := &Account{
		Currency: "810",
		Amount:   100.0,
		UserID:   uuid.New(),
		Name:     "AccountName",
	}
	err := a.Validate()

	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Тест случая с некорректной валютой
	a.Currency = "invalid"
	err = a.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Тест случая с пустой валютой
	a.Currency = ""
	err = a.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Тест случая с отрицательным балансом
	a.Currency = "810"
	a.Amount = -50.0

	err = a.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Тест случая с пустым id
	a.ID = uuid.Nil
	err = a.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	// Тест случая с пустым name
	a.Name = ""
	err = a.Validate()

	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}

func TestGenerateInterestRate(t *testing.T) {
	// Тест случая генерации процентной ставки
	rate := generateInterestRate()
	if rate < 0 || rate >= 10 {
		t.Errorf("Expected rate between 0 and 10, but got %f", rate)
	}
}
