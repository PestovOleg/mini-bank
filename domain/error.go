// Сборник однотипных ошибок
package domain

import "errors"

var (
	// ErrMustBeFilledIn Must be filled in
	ErrMustBeFilledIn = errors.New("one or more fields must be filled in")

	// NotFound not found
	ErrNotFound = errors.New("not found")
)
