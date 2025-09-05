package main

import (
	"os"

	"hexhello/src/adapters/driven"
	"hexhello/src/app"
)

func main() {
	presenter := driven.NewConsolePresenter()
	usecase := app.NewGreetUserService(presenter)

	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	usecase.Greet(name)
}
