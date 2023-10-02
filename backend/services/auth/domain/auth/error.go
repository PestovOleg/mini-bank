// Сборник однотипных ошибок
package auth

import "errors"

var (
	// ErrMustBeFilledIn Must be filled in
	ErrMustBeFilledIn = errors.New("one or more fields must be filled in")

	// ErrNotFound User not found
	ErrNotFound = errors.New("user not found")

	// ErrIDMustBeEntered ID must be entered
	ErrIDMustBeEntered = errors.New("ID must be entered")

	// ErrEmptyUsername Username must be filled in
	ErrEmptyUsername = errors.New("username must be filled in")

	// ErrEmptyPassword Password must be filled in
	ErrEmptyPassword = errors.New("password must be filled in")

	// ErrEmptyBirthday Birthday must be filled in
	ErrEmptyBirthday = errors.New("birthday must be filled in")

	// ErrEmptyPassword Password must be filled in
	ErrEmptyPhone = errors.New("phone must be filled in")

	// ErrAuthentication Authentication failed
	ErrAuthentication = errors.New("authentication failed")

	// ErrUnauthorized Unauthorized
	ErrUnauthorized = errors.New("unauthorized ")
)
