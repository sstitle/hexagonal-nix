package domain

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// Username validation constants
const (
	MinUsernameLength   = 3
	MaxUsernameLength   = 20
	MinPasswordLength   = 6
	MaxProfileMsgLength = 500
)

// UsernameRegex defines valid username characters (alphanumeric and underscore)
var UsernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)

// ValidateUsername validates a username according to business rules
func ValidateUsername(username string) error {
	if username == "" {
		return ValidationError{Field: "username", Message: "username is required"}
	}

	username = strings.TrimSpace(username)

	if len(username) < MinUsernameLength {
		return ValidationError{Field: "username", Message: "username must be at least 3 characters long"}
	}

	if len(username) > MaxUsernameLength {
		return ValidationError{Field: "username", Message: "username must be no more than 20 characters long"}
	}

	if !UsernameRegex.MatchString(username) {
		return ValidationError{Field: "username", Message: "username can only contain letters, numbers, and underscores"}
	}

	return nil
}

// ValidatePassword validates a password according to business rules
func ValidatePassword(password string) error {
	if password == "" {
		return ValidationError{Field: "password", Message: "password is required"}
	}

	if len(password) < MinPasswordLength {
		return ValidationError{Field: "password", Message: "password must be at least 6 characters long"}
	}

	return nil
}

// ValidateProfileMessage validates a profile message
func ValidateProfileMessage(message string) error {
	if message == "" {
		return nil // Profile message is optional
	}

	if utf8.RuneCountInString(message) > MaxProfileMsgLength {
		return ValidationError{Field: "profile_message", Message: "profile message must be no more than 500 characters"}
	}

	return nil
}
