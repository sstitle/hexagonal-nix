package ports

import "hexagonal-nix/entities"

// UserRepository defines the interface for user persistence operations
type UserRepository interface {
	// Create saves a new user to the repository
	Create(user *entities.User) error

	// GetByID retrieves a user by their ID
	GetByID(id string) (*entities.User, error)

	// GetByUsername retrieves a user by their username
	GetByUsername(username string) (*entities.User, error)

	// Update updates an existing user
	Update(user *entities.User) error

	// Delete removes a user from the repository
	Delete(id string) error

	// List returns all users in the repository
	List() ([]*entities.User, error)

	// Exists checks if a user exists with the given username
	Exists(username string) (bool, error)
}
