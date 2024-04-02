package pet

import (
	"errors"
)

var (
	// ErrNotFound is returned when a pet is not found
	ErrNotFound         = errors.New("pet not found")
	ErrNoValidName      = errors.New("a valid name should be provided")
	ErrNoValidBreedname = errors.New("a valid breed should be provided")
	ErrNoValidBirthDate = errors.New("a valid birthdate should be provided")
)
