package driven

import "fmt"

type ConsolePresenter struct{}

func NewConsolePresenter() *ConsolePresenter { return &ConsolePresenter{} }

func (p *ConsolePresenter) PresentGreeting(message string) {
	fmt.Println(message)
}
