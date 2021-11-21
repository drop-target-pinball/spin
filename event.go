package spin

type Event interface {
	Key() interface{}
}

type Message struct {
	ID string
}

func (e Message) Key() interface{} {
	return e.ID
}

type SwitchEvent struct {
	ID       string
	Released bool
}

func (e SwitchEvent) Key() interface{} {
	return struct {
		ID       string
		Released bool
	}{
		e.ID,
		e.Released,
	}
}

func registerEvents(e *Engine) {
	e.RegisterEvent(Message{})
	e.RegisterEvent(SwitchEvent{})
}
