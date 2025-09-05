package test

import (
	"fmt"
	"sync"

	"hexagonal-nix/entities"
	"hexagonal-nix/ports"
)

// MockUserRepository is a mock implementation of UserRepository for testing
type MockUserRepository struct {
	users map[string]*entities.User
	mutex sync.RWMutex
}

// NewMockUserRepository creates a new mock repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[string]*entities.User),
	}
}

// Create saves a new user to the repository
func (m *MockUserRepository) Create(user *entities.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.users[user.ID] = user
	return nil
}

// GetByID retrieves a user by their ID
func (m *MockUserRepository) GetByID(id string) (*entities.User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	user, exists := m.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

// GetByUsername retrieves a user by their username
func (m *MockUserRepository) GetByUsername(username string) (*entities.User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

// Update updates an existing user
func (m *MockUserRepository) Update(user *entities.User) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.users[user.ID]; !exists {
		return fmt.Errorf("user not found")
	}

	m.users[user.ID] = user
	return nil
}

// Delete removes a user from the repository
func (m *MockUserRepository) Delete(id string) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.users[id]; !exists {
		return fmt.Errorf("user not found")
	}

	delete(m.users, id)
	return nil
}

// List returns all users in the repository
func (m *MockUserRepository) List() ([]*entities.User, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	users := make([]*entities.User, 0, len(m.users))
	for _, user := range m.users {
		users = append(users, user)
	}

	return users, nil
}

// Exists checks if a user exists with the given username
func (m *MockUserRepository) Exists(username string) (bool, error) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for _, user := range m.users {
		if user.Username == username {
			return true, nil
		}
	}

	return false, nil
}

// Ensure MockUserRepository implements UserRepository
var _ ports.UserRepository = (*MockUserRepository)(nil)
