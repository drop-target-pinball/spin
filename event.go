package spin

type Event interface {
	event()
}

type Message struct {
	ID string
}

type SwitchEvent struct {
	ID string
}

func (Message) event()     {}
func (SwitchEvent) event() {}
