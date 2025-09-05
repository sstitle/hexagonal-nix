package app

import "testing"

type stubPresenter struct{ got string }

func (s *stubPresenter) PresentGreeting(m string) { s.got = m }

func TestGreet(t *testing.T) {
	sp := &stubPresenter{}
	svc := NewGreetUserService(sp)
	svc.Greet("Alice")
	if sp.got != "Hello, Alice!" {
		t.Fatalf("got %q", sp.got)
	}
}
