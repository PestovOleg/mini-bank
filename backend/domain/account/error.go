package account

import "errors"

var (
	// ErrCurrencyMustBeEntered Currency must be entered
	ErrCurrencyMustBeEntered = errors.New("currency must be entered")

	// ErrAccountLength Account must be 20 symbols
	ErrAccountLength = errors.New("account must be 20 symbols")

	// ErrAccountMustBeDigits Account must contains only digits
	ErrAccountMustBeDigits = errors.New("account must contains only digits")

	// ErrCurrencyMustBeDigits Currency must contains only digits
	ErrCurrencyMustBeDigits = errors.New("currency must contains only digits")

	// ErrUserIDMustBeEntered User ID must be entered
	ErrUserIDMustBeEntered = errors.New("user ID must be entered")

	// ErrAccountNameMustBeEntered User ID must be entered
	ErrAccountNameMustBeEntered = errors.New("account name must be entered")

	// ErrMustBePositiveOrZero Sum must be positive or more then 0
	ErrMustBePositiveOrZero = errors.New("sum must be positive or more than 0")

	// ErrMustBePositive Sum must be more then 0
	ErrMustBePositive = errors.New("sum must be more than 0")

	// ErrNotEnoughMoney Not enough money
	ErrNotEnoughMoney = errors.New("вам нужно больше зарабатывать,чтобы столько снимать")

	// ErrNotFound Account not found
	ErrNotFound = errors.New("account not found")
)
