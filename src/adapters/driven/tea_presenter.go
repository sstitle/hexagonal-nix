package driven

// TeaPresenter implements GreeterPresenter by sending the greeting message
// on a channel to be consumed by a Bubble Tea program.
type TeaPresenter struct {
	ch chan<- string
}

func NewTeaPresenter(ch chan<- string) *TeaPresenter {
	return &TeaPresenter{ch: ch}
}

func (p *TeaPresenter) PresentGreeting(message string) {
	// Non-blocking send would drop messages; block to ensure delivery is fine here.
	p.ch <- message
}
