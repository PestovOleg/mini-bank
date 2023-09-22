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
)
