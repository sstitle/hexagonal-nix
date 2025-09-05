package domain

// GreetUser is the input port representing the use case offered by the application.
type GreetUser interface {
	Greet(name string)
}

// GreeterPresenter is the output port the application depends on to present the greeting.
type GreeterPresenter interface {
	PresentGreeting(message string)
}
