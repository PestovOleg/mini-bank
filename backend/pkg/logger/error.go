package logger

import "errors"

var (
	// ErrCouldNotInit could not initialize rootLogger
	ErrCouldNotInit = errors.New("could not initialize rootLogger")
)
