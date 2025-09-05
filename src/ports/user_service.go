package ports

import "hexagonal-nix/entities"

// UserService defines the interface for user business operations
type UserService interface {
	// Register creates a new user account
	Register(username, password string) (*entities.User, error)

	// Authenticate verifies user credentials
	Authenticate(username, password string) (*entities.User, error)

	// GetProfile retrieves a user's profile by ID
	GetProfile(userID string) (*entities.User, error)

	// UpdateProfile updates a user's profile message
	UpdateProfile(userID, message string) error

	// ListUsers returns all users (for browsing profiles)
	ListUsers() ([]*entities.User, error)

	// DeleteAccount removes a user account
	DeleteAccount(userID string) error
}
