package main

import (
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"

	"hexhello/src/adapters/driven"
	"hexhello/src/app"
)

type model struct {
	input   textinput.Model
	message string
	done    bool
	greet   func(string)
}

func (m model) Init() tea.Cmd { return textinput.Blink }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch v := msg.(type) {
	case tea.KeyMsg:
		if v.Type == tea.KeyEnter {
			name := m.input.Value()
			if name == "" {
				name = "World"
			}
			if m.greet != nil {
				m.greet(name)
			}
			return m, nil
		}
		var cmd tea.Cmd
		m.input, cmd = m.input.Update(msg)
		return m, cmd
	case string:
		m.message = v
		m.done = true
		return m, tea.Quit
	default:
		return m, nil
	}
}

func (m model) View() string {
	if m.done {
		return m.message + "\n"
	}
	return "What's your name?\n" + m.input.View() + "\n"
}

func main() {
	// Channel between presenter and Bubble Tea program
	ch := make(chan string, 1)
	presenter := driven.NewTeaPresenter(ch)
	svc := app.NewGreetUserService(presenter)

	m := model{input: textinput.New()}
	m.input.Placeholder = "World"
	m.input.Focus()
	m.greet = func(name string) { svc.Greet(name) }

	// Start Bubble Tea program
	p := tea.NewProgram(m)
	// Forward greeting messages from presenter into the program
	go func() { p.Send(<-ch) }()
	if _, err := p.Run(); err != nil {
		os.Exit(1)
	}
}
