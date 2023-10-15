package account

import (
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Account      string
	Currency     string
	Name         string
	Amount       float64
	InterestRate float64
	IsActive     bool
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func NewAccount(userID uuid.UUID, currency, account, name string) (*Account, error) {
	a := &Account{
		ID:           uuid.New(),
		Account:      account,
		Currency:     currency,
		Name:         name,
		Amount:       0,
		InterestRate: generateInterestRate(),
		UserID:       userID,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		IsActive:     true,
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
	if len(account) != 20 && len(account) != 0 {
		return "", ErrAccountLength
	}

	if account != "" {
		matched, err := regexp.MatchString(`^[0-9]+$`, account)
		if err != nil {
			return "", err
		}

		if !matched {
			return "", ErrAccountMustBeDigits
		}
	}

	if currency != "" {
		matched, err := regexp.MatchString(`^[0-9]+$`, currency)
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
		account = "40817"
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

	if a.Currency != "810" && a.Currency != "840" {
		return ErrCurrencyMustBeEntered
	}

	if a.Amount < 0 {
		return ErrMustBePositiveOrZero
	}

	if a.UserID == uuid.Nil {
		return ErrUserIDMustBeEntered
	}

	if a.Name == "" {
		return ErrAccountNameMustBeEntered
	}

	return nil
}

// generate random rate as a decimal(5,4)
func generateInterestRate() float64 {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	intPart := r.Intn(10)
	decimalPart := r.Float64()
	rate := float64(intPart) + decimalPart

	// Округляем до 4х знаков после запятой
	return float64(int(rate*10000)) / 10000
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
	Create(a *Account) error
	Update(a *Account) error
	Delete(id uuid.UUID) error
}

// Repository -композиция интерфейсов Writer и Reader
type Repository interface {
	Reader
	Writer
}

// Usecases
// GetAccountByID - get account by ID
// GetAccountByNumber - get account by account' number
// GetAccountByIDAndUserID - get account by ID (safe)
// ListAccount - get all clients' accounts
// CreateAccount - create account
// DeleteAccount - deactivate account
// TopUp - top up account balance
// Withdraw - withdraw money

type UseCase interface {
	GetAccountByID(id uuid.UUID) (*Account, error)
	GetAccountByNumber(account string) (*Account, error)
	GetAccountByIDAndUserID(id, userID uuid.UUID) (*Account, error)
	ListAccount(userID uuid.UUID) ([]*Account, error)
	CreateAccount(userID uuid.UUID, currency string) (*Account, error)
	TopUp(id uuid.UUID, money float64)
	WithDraw(id uuid.UUID, money float64)
	DeleteAccount(id uuid.UUID) error
	UpdateAccount(id uuid.UUID, name string, rate float64) error
}
