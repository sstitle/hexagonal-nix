package secondary

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"hexagonal-nix/entities"
	"hexagonal-nix/ports"
)

// JSONUserRepository implements UserRepository using JSON file storage
type JSONUserRepository struct {
	filePath string
	users    map[string]*entities.User
	mutex    sync.RWMutex
}

// NewJSONUserRepository creates a new JSON-based user repository
func NewJSONUserRepository(filePath string) *JSONUserRepository {
	return &JSONUserRepository{
		filePath: filePath,
		users:    make(map[string]*entities.User),
	}
}

// Load loads users from the JSON file
func (r *JSONUserRepository) Load() error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Create directory if it doesn't exist
	dir := filepath.Dir(r.filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// Check if file exists
	if _, err := os.Stat(r.filePath); os.IsNotExist(err) {
		// File doesn't exist, start with empty map
		return nil
	}

	// Read and parse JSON file
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	if len(data) == 0 {
		// Empty file, start with empty map
		return nil
	}

	if err := json.Unmarshal(data, &r.users); err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	return nil
}

// Save saves users to the JSON file
func (r *JSONUserRepository) Save() error {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	data, err := json.MarshalIndent(r.users, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(r.filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// Create saves a new user to the repository
func (r *JSONUserRepository) Create(user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.Load(); err != nil {
		return err
	}

	r.users[user.ID] = user
	return r.Save()
}

// GetByID retrieves a user by their ID
func (r *JSONUserRepository) GetByID(id string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if err := r.Load(); err != nil {
		return nil, err
	}

	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// GetByUsername retrieves a user by their username
func (r *JSONUserRepository) GetByUsername(username string) (*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if err := r.Load(); err != nil {
		return nil, err
	}

	for _, user := range r.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

// Update updates an existing user
func (r *JSONUserRepository) Update(user *entities.User) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.Load(); err != nil {
		return err
	}

	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user not found")
	}

	r.users[user.ID] = user
	return r.Save()
}

// Delete removes a user from the repository
func (r *JSONUserRepository) Delete(id string) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.Load(); err != nil {
		return err
	}

	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	delete(r.users, id)
	return r.Save()
}

// List returns all users in the repository
func (r *JSONUserRepository) List() ([]*entities.User, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if err := r.Load(); err != nil {
		return nil, err
	}

	users := make([]*entities.User, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user)
	}

	return users, nil
}

// Exists checks if a user exists with the given username
func (r *JSONUserRepository) Exists(username string) (bool, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	if err := r.Load(); err != nil {
		return false, err
	}

	for _, user := range r.users {
		if user.Username == username {
			return true, nil
		}
	}

	return false, nil
}

// Ensure JSONUserRepository implements UserRepository
var _ ports.UserRepository = (*JSONUserRepository)(nil)
