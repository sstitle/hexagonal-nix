package main

import (
	"fmt"
	"log"
	"os"

	"hexagonal-nix/adapters/secondary"
)

func main() {
	fmt.Println("🚀 Hexagonal Architecture User Profile System")
	fmt.Println("Phase 1: TUI + JSON Storage")
	fmt.Println()

	// For now, just test that our basic structure compiles
	// This will be replaced with the actual TUI implementation
	repo := secondary.NewJSONUserRepository("./data/users.json")

	if err := repo.Load(); err != nil {
		log.Printf("Failed to load repository: %v", err)
		os.Exit(1)
	}

	fmt.Println("✅ Repository loaded successfully")
	fmt.Println("✅ Basic hexagonal architecture structure is working")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("- Implement TUI interface")
	fmt.Println("- Add user service implementation")
	fmt.Println("- Add authentication logic")
}
