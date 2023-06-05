package utils

import (
	"errors"
	"regexp"
)

func TextIsEmpty(text string) bool {
	return text == ""
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

// IsPasswordValid
// Password should be of 8 characters long
// Password should contain atleast one lower case character
// Password should contain at least one upper case character
// Password should contain atleast one digit
// Password should contain at least one special character
func IsPasswordValid(password string) error {
	if len(password) < 8 {
		return errors.New("password should be of 8 characters long")
	}
	done, err := regexp.MatchString("([a-z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return errors.New("password should contain atleast one lower case character")
	}
	done, err = regexp.MatchString("([A-Z])+", password)
	if err != nil {
		return err
	}
	if !done {
		return errors.New("password should contain at least one upper case character")
	}
	done, err = regexp.MatchString("([0-9])+", password)
	if err != nil {
		return err
	}
	if !done {
		return errors.New("password should contain atleast one digit")
	}

	done, err = regexp.MatchString("([!@#$%^&*.?-])+", password)
	if err != nil {
		return err
	}
	if !done {
		return errors.New("password should contain at least one special character")
	}
	return nil
}
