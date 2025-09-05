package domain

import (
	"testing"
	"time"

	"hexagonal-nix/entities"
)

func TestNewUser(t *testing.T) {
	username := "testuser"
	passwordHash := "hashedpassword"

	user := entities.NewUser(username, passwordHash)

	if user.Username != username {
		t.Errorf("Expected username %s, got %s", username, user.Username)
	}

	if user.PasswordHash != passwordHash {
		t.Errorf("Expected password hash %s, got %s", passwordHash, user.PasswordHash)
	}

	if user.ID == "" {
		t.Error("Expected user ID to be generated")
	}

	if user.ProfileMsg != "" {
		t.Error("Expected profile message to be empty")
	}

	if user.CreatedAt.IsZero() {
		t.Error("Expected CreatedAt to be set")
	}

	if user.UpdatedAt.IsZero() {
		t.Error("Expected UpdatedAt to be set")
	}
}

func TestUpdateProfile(t *testing.T) {
	user := entities.NewUser("testuser", "hash")
	originalUpdatedAt := user.UpdatedAt

	// Wait a bit to ensure time difference
	time.Sleep(1 * time.Millisecond)

	profileMsg := "Hello, world!"
	user.UpdateProfile(profileMsg)

	if user.ProfileMsg != profileMsg {
		t.Errorf("Expected profile message %s, got %s", profileMsg, user.ProfileMsg)
	}

	if !user.UpdatedAt.After(originalUpdatedAt) {
		t.Error("Expected UpdatedAt to be updated")
	}
}
