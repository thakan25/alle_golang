package utils

import (
	"regexp"
	"task-manager/common"
)

// ValidateEmail checks if the email is valid
func ValidateEmail(email string) error {
	if email == "" {
		return common.ErrInvalidRequest
	}
	
	// Basic email regex validation
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return common.ErrInvalidRequest
	}
	
	return nil
}

// ValidatePassword checks if the password meets requirements
func ValidatePassword(password string) error {
	if password == "" {
		return common.ErrInvalidRequest
	}
	
	// Password should be at least 8 characters
	if len(password) < 8 {
		return common.ErrInvalidRequest
	}
	
	return nil
}

// ValidateUsername checks if the username is valid
func ValidateUsername(username string) error {
	if username == "" {
		return common.ErrInvalidRequest
	}
	
	// Username should be between 3 and 20 characters
	if len(username) < 3 || len(username) > 20 {
		return common.ErrInvalidRequest
	}
	
	return nil
} 