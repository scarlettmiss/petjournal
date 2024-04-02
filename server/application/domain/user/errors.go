package user

import (
	"errors"
)

var (
	// ErrNotFound is returned when a user is not found
	ErrNotFound            = errors.New("user not found")
	ErrMailExists          = errors.New("mail in use")
	ErrNoValidMail         = errors.New("a valid mail should be provided")
	ErrNoValidName         = errors.New("a valid name should be provided")
	ErrNoValidSurname      = errors.New("a valid surname should be provided")
	ErrAuthentication      = errors.New("wrong credentials")
	ErrUserDeleted         = errors.New("user has been deleted")
	ErrNoValidType         = errors.New("a valid userType should be provided")
	ErrPasswordLength      = errors.New("password should be of 8 characters long")
	ErrPasswordLowerCase   = errors.New("password should contain at least one lower case character")
	ErrPasswordUpperCase   = errors.New("password should contain at least one upper case character")
	ErrPasswordDigit       = errors.New("password should contain atleast one digit")
	ErrPasswordSpecialChar = errors.New("password should contain at least one special character")
)
