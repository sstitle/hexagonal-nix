package app

import (
	"hexhello/src/domain"
)

type GreetUserService struct {
	presenter domain.GreeterPresenter
}

func NewGreetUserService(presenter domain.GreeterPresenter) *GreetUserService {
	return &GreetUserService{presenter: presenter}
}

func (s *GreetUserService) Greet(name string) {
	message := "Hello, " + name + "!"
	s.presenter.PresentGreeting(message)
}
