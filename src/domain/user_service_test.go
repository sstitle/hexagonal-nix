package domain

import (
	"testing"

	"hexagonal-nix/domain/test"
)

func TestUserService_Register(t *testing.T) {
	repo := test.NewMockUserRepository()
	service := NewUserService(repo)

	t.Run("successful registration", func(t *testing.T) {
		user, err := service.Register("testuser", "password123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if user.Username != "testuser" {
			t.Errorf("Expected username testuser, got %s", user.Username)
		}

		if user.PasswordHash == "" {
			t.Error("Expected password to be hashed")
		}

		// Verify user was saved
		exists, err := repo.Exists("testuser")
		if err != nil {
			t.Fatalf("Expected no error checking existence, got %v", err)
		}
		if !exists {
			t.Error("Expected user to exist in repository")
		}
	})

	t.Run("duplicate username", func(t *testing.T) {
		// Register first user
		_, err := service.Register("duplicate", "password123")
		if err != nil {
			t.Fatalf("Expected no error for first registration, got %v", err)
		}

		// Try to register same username again
		_, err = service.Register("duplicate", "password456")
		if err != ErrUserAlreadyExists {
			t.Errorf("Expected ErrUserAlreadyExists, got %v", err)
		}
	})

	t.Run("invalid username", func(t *testing.T) {
		_, err := service.Register("ab", "password123")
		if err == nil {
			t.Error("Expected validation error for short username")
		}
	})

	t.Run("invalid password", func(t *testing.T) {
		_, err := service.Register("testuser2", "12345")
		if err == nil {
			t.Error("Expected validation error for short password")
		}
	})
}

func TestUserService_Authenticate(t *testing.T) {
	repo := test.NewMockUserRepository()
	service := NewUserService(repo)

	// Register a user first
	_, err := service.Register("testuser", "password123")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	t.Run("successful authentication", func(t *testing.T) {
		user, err := service.Authenticate("testuser", "password123")
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if user.Username != "testuser" {
			t.Errorf("Expected username testuser, got %s", user.Username)
		}
	})

	t.Run("wrong password", func(t *testing.T) {
		_, err := service.Authenticate("testuser", "wrongpassword")
		if err != ErrInvalidCredentials {
			t.Errorf("Expected ErrInvalidCredentials, got %v", err)
		}
	})

	t.Run("non-existent user", func(t *testing.T) {
		_, err := service.Authenticate("nonexistent", "password123")
		if err != ErrInvalidCredentials {
			t.Errorf("Expected ErrInvalidCredentials, got %v", err)
		}
	})

	t.Run("invalid username", func(t *testing.T) {
		_, err := service.Authenticate("ab", "password123")
		if err == nil {
			t.Error("Expected validation error for short username")
		}
	})
}

func TestUserService_UpdateProfile(t *testing.T) {
	repo := test.NewMockUserRepository()
	service := NewUserService(repo)

	// Register a user first
	user, err := service.Register("testuser", "password123")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	t.Run("successful profile update", func(t *testing.T) {
		profileMsg := "Hello, world!"
		err := service.UpdateProfile(user.ID, profileMsg)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify update
		updatedUser, err := service.GetProfile(user.ID)
		if err != nil {
			t.Fatalf("Expected no error getting profile, got %v", err)
		}

		if updatedUser.ProfileMsg != profileMsg {
			t.Errorf("Expected profile message %s, got %s", profileMsg, updatedUser.ProfileMsg)
		}
	})

	t.Run("update non-existent user", func(t *testing.T) {
		err := service.UpdateProfile("nonexistent", "message")
		if err != ErrUserNotFound {
			t.Errorf("Expected ErrUserNotFound, got %v", err)
		}
	})

	t.Run("invalid profile message", func(t *testing.T) {
		longMessage := string(make([]byte, MaxProfileMsgLength+1))
		err := service.UpdateProfile(user.ID, longMessage)
		if err == nil {
			t.Error("Expected validation error for long profile message")
		}
	})
}

func TestUserService_ListUsers(t *testing.T) {
	t.Run("list all users", func(t *testing.T) {
		repo := test.NewMockUserRepository()
		service := NewUserService(repo)

		// Register some users
		_, err := service.Register("user1", "password123")
		if err != nil {
			t.Fatalf("Failed to register user1: %v", err)
		}

		_, err = service.Register("user2", "password456")
		if err != nil {
			t.Fatalf("Failed to register user2: %v", err)
		}

		users, err := service.ListUsers()
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		if len(users) != 2 {
			t.Errorf("Expected 2 users, got %d. Users: %+v", len(users), users)
		}
	})
}

func TestUserService_DeleteAccount(t *testing.T) {
	repo := test.NewMockUserRepository()
	service := NewUserService(repo)

	// Register a user first
	user, err := service.Register("testuser", "password123")
	if err != nil {
		t.Fatalf("Failed to register user: %v", err)
	}

	t.Run("successful deletion", func(t *testing.T) {
		err := service.DeleteAccount(user.ID)
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}

		// Verify user is deleted
		_, err = service.GetProfile(user.ID)
		if err != ErrUserNotFound {
			t.Errorf("Expected ErrUserNotFound, got %v", err)
		}
	})

	t.Run("delete non-existent user", func(t *testing.T) {
		err := service.DeleteAccount("nonexistent")
		if err != ErrUserNotFound {
			t.Errorf("Expected ErrUserNotFound, got %v", err)
		}
	})
}
