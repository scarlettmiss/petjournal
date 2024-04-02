package record

import (
	"errors"
)

var (
	// ErrNotFound is returned when a record is not found
	ErrNotFound         = errors.New("record not found")
	ErrNotValidName     = errors.New("record name not valid")
	ErrNotValidResult   = errors.New("record result not valid")
	ErrNotValidDate     = errors.New("record date not valid")
	ErrNotValidType     = errors.New("record type not valid")
	ErrNotValidVerifier = errors.New("record cannot be validated by this user")
)
