package domain

import (
	"strings"
	"testing"
)

func TestValidateUsername(t *testing.T) {
	tests := []struct {
		name     string
		username string
		wantErr  bool
	}{
		{"valid username", "testuser", false},
		{"valid username with numbers", "test123", false},
		{"valid username with underscore", "test_user", false},
		{"empty username", "", true},
		{"whitespace only", "   ", true},
		{"too short", "ab", true},
		{"too long", strings.Repeat("a", MaxUsernameLength+1), true},
		{"invalid characters", "test-user", true},
		{"invalid characters", "test@user", true},
		{"invalid characters", "test user", true},
		{"starts with number", "123test", false}, // This should be valid
		{"only numbers", "123456", false},        // This should be valid
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateUsername(tt.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateUsername() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"valid password", "password123", false},
		{"minimum length", "123456", false},
		{"empty password", "", true},
		{"too short", "12345", true},
		{"special characters", "pass@word!", false},
		{"unicode characters", "пароль123", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateProfileMessage(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wantErr bool
	}{
		{"empty message", "", false},
		{"valid message", "Hello, world!", false},
		{"long message", strings.Repeat("a", MaxProfileMsgLength), false},
		{"too long message", strings.Repeat("a", MaxProfileMsgLength+1), true},
		{"unicode message", "Привет, мир! 🌍", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProfileMessage(tt.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateProfileMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
