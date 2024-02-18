package utils

import (
	"errors"
	"regexp"
	"strings"
)

func TextIsEmpty(text string) bool {
	return strings.TrimSpace(text) == ""
}

func IsEmailValid(email string) bool {
	if TextIsEmpty(email) {
		return false
	}
	// Regular expression pattern for email validation
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern into a regular expression object
	regExp := regexp.MustCompile(pattern)

	// Match the email against the regular expression
	return regExp.MatchString(email)
}

var (
	// ErrNotFound is returned when a user is not found
	ErrPasswordLength      = errors.New("password should be of 8 characters long")
	ErrPasswordLowerCase   = errors.New("password should contain at least one lower case character")
	ErrPasswordUpperCase   = errors.New("password should contain at least one upper case character")
	ErrPasswordDigit       = errors.New("password should contain atleast one digit")
	ErrPasswordSpecialChar = errors.New("password should contain at least one special character")
)

// IsPasswordValid
// Password should be of 8 characters long
// Password should contain atleast one lower case character
// Password should contain at least one upper case character
// Password should contain atleast one digit
// Password should contain at least one special character
func IsPasswordValid(password string) error {
	if len(password) < 8 {
		return ErrPasswordLength
	}
	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return ErrPasswordLowerCase
	}
	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return ErrPasswordUpperCase
	}
	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}
	if !done {
		return ErrPasswordDigit
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}
	if !done {
		return ErrPasswordSpecialChar
	}
	return nil
}

func ErrorResponse(err error) map[string]any {
	return map[string]any{
		"error": err.Error(),
	}
}
