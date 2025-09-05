package entities

import (
	"time"
)

// User represents a user in the system
type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	PasswordHash string    `json:"password_hash"`
	ProfileMsg   string    `json:"profile_msg"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// NewUser creates a new user with the given username and password hash
func NewUser(username, passwordHash string) *User {
	now := time.Now()
	return &User{
		ID:           generateID(),
		Username:     username,
		PasswordHash: passwordHash,
		ProfileMsg:   "",
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

// UpdateProfile updates the user's profile message
func (u *User) UpdateProfile(message string) {
	u.ProfileMsg = message
	u.UpdatedAt = time.Now()
}

// generateID generates a simple ID for the user
// In a real application, this would use a proper UUID generator
func generateID() string {
	return time.Now().Format("20060102150405.000000")
}
