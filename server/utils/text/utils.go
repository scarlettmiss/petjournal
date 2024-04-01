package text

import (
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
