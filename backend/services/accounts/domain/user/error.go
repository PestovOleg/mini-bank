// Сборник однотипных ошибок
package user

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

	// ErrEmptyEmail Email must be filled in
	ErrEmptyEmail = errors.New("email must be filled in")

	// ErrEmptyName Name must be filled in
	ErrEmptyName = errors.New("name must be filled in")

	// ErrEmptyLastName LastName must be filled in
	ErrEmptyLastName = errors.New("lastName must be filled in")

	// ErrEmptyPatronymic Patronymic must be filled in
	ErrEmptyPatronymic = errors.New("patronymic must be filled in")

	// ErrEmptyPassword Password must be filled in
	ErrEmptyPassword = errors.New("password must be filled in")

	// ErrEmptyBirthday Birthday must be filled in
	ErrEmptyBirthday = errors.New("birthday must be filled in")

	// ErrEmptyPassword Password must be filled in
	ErrEmptyPhone = errors.New("phone must be filled in")
)
