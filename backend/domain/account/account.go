package account

import (
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Account   string
	Currency  string
	IsActive  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAccount(userID uuid.UUID, currency, account string) (*Account, error) {
	a := &Account{
		ID:        uuid.New(),
		Account:   account,
		Currency:  currency,
		UserID:    userID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	err := a.Validate()
	if err != nil {
		return nil, err
	}

	a.Account, err = NextAccount(a.Currency, a.Account)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func NextAccount(currency, account string) (string, error) {
	if len(account) != 20 {
		return "", ErrAccountLength
	}

	if account != "" {
		matched, err := regexp.MatchString(`0123456789`, account)
		if err != nil {
			return "", err
		}

		if !matched {
			return "", ErrAccountMustBeDigits
		}
	}

	if currency != "" {
		matched, err := regexp.MatchString(`0123456789`, currency)
		if err != nil {
			return "", err
		}

		if !matched {
			return "", ErrCurrencyMustBeDigits
		}
	}

	// Контрольный ключ счета
	kk := "9"

	var toDigits int

	if account == "" {
		toDigits = 0
	} else {
		toDigits, _ = strconv.Atoi(account[13:])
	}
	toDigits++

	mapSuffix := "0000000"
	newSuffix := mapSuffix[:7-len(strconv.Itoa(toDigits))] + strconv.Itoa(toDigits)
	nextAccount := account[:5] + currency + kk + "0200" + newSuffix

	return nextAccount, nil
}

func (a *Account) Validate() error {
	if a.Currency == "" {
		return ErrCurrencyMustBeEntered
	}

	if a.UserID == uuid.Nil {
		return ErrUserIDMustBeEntered
	}

	return nil
}

// Reader
// GetAccountByID - get record by ID
// GetAccountByNumber - get record by account' number
// List - get all records
// GetLastOpenedAccount -get last opened account (for new account)
type Reader interface {
	GetByID(id uuid.UUID) (*Account, error)
	GetByNumber(acc string) (*Account, error)
	List(userID uuid.UUID) ([]*Account, error)
	GetLastOpenedAccount(currency string) (string, error)
}

// Writer
// Create - create record
// Update - update record
// Delete - delete record (deactivate)
type Writer interface {
	Create(a *Account) (uuid.UUID, error)
	Delete(id uuid.UUID) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecase интерфейс
// GetAccountByID - get account by ID
// GetAccountByNumber - get account by account' number
// ListAccount - get all clients' accounts
// CreateAccount - create account
// DeleteAccount - deactivate account

type UseCase interface {
	GetAccountByID(id uuid.UUID) (*Account, error)
	GetAccountByNumber(account string) (*Account, error)
	ListAccount() ([]*Account, error)
	CreateAccount(userID uuid.UUID, currency string) (*Account, error)
	DeleteAccount(id uuid.UUID) error
}
