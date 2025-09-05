package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"hexhello/src/adapters/driven"
	"hexhello/src/app"
)

func main() {
	presenter := driven.NewConsolePresenter()
	usecase := app.NewGreetUserService(presenter)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	input, _ := reader.ReadString('\n')
	name := strings.TrimSpace(input)
	if name == "" {
		name = "World"
	}
	usecase.Greet(name)
}
