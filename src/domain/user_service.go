package domain

import (
	"crypto/sha256"
	"fmt"

	"hexagonal-nix/entities"
	"hexagonal-nix/ports"
)

// UserService implements the business logic for user management
type UserService struct {
	userRepo ports.UserRepository
}

// NewUserService creates a new UserService instance
func NewUserService(userRepo ports.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Register creates a new user account with validation
func (s *UserService) Register(username, password string) (*entities.User, error) {
	// Validate input
	if err := ValidateUsername(username); err != nil {
		return nil, err
	}

	if err := ValidatePassword(password); err != nil {
		return nil, err
	}

	// Check if user already exists
	exists, err := s.userRepo.Exists(username)
	if err != nil {
		return nil, fmt.Errorf("failed to check user existence: %w", err)
	}

	if exists {
		return nil, ErrUserAlreadyExists
	}

	// Hash password
	passwordHash := hashPassword(password)

	// Create user
	user := entities.NewUser(username, passwordHash)

	// Save to repository
	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

// Authenticate verifies user credentials
func (s *UserService) Authenticate(username, password string) (*entities.User, error) {
	// Validate input
	if err := ValidateUsername(username); err != nil {
		return nil, err
	}

	if err := ValidatePassword(password); err != nil {
		return nil, err
	}

	// Get user by username
	user, err := s.userRepo.GetByUsername(username)
	if err != nil {
		return nil, ErrInvalidCredentials
	}

	// Verify password
	if !verifyPassword(password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

// GetProfile retrieves a user's profile by ID
func (s *UserService) GetProfile(userID string) (*entities.User, error) {
	if userID == "" {
		return nil, ValidationError{Field: "user_id", Message: "user ID is required"}
	}

	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

// UpdateProfile updates a user's profile message
func (s *UserService) UpdateProfile(userID, message string) error {
	if userID == "" {
		return ValidationError{Field: "user_id", Message: "user ID is required"}
	}

	if err := ValidateProfileMessage(message); err != nil {
		return err
	}

	// Get existing user
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Update profile
	user.UpdateProfile(message)

	// Save changes
	if err := s.userRepo.Update(user); err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

// ListUsers returns all users (for browsing profiles)
func (s *UserService) ListUsers() ([]*entities.User, error) {
	users, err := s.userRepo.List()
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}

	return users, nil
}

// DeleteAccount removes a user account
func (s *UserService) DeleteAccount(userID string) error {
	if userID == "" {
		return ValidationError{Field: "user_id", Message: "user ID is required"}
	}

	// Check if user exists
	_, err := s.userRepo.GetByID(userID)
	if err != nil {
		return ErrUserNotFound
	}

	// Delete user
	if err := s.userRepo.Delete(userID); err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}

// hashPassword creates a simple hash of the password
// In a real application, use bcrypt or similar
func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return fmt.Sprintf("%x", hash)
}

// verifyPassword verifies a password against its hash
func verifyPassword(password, hash string) bool {
	expectedHash := hashPassword(password)
	return expectedHash == hash
}
