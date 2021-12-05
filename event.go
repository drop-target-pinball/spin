package spin

import "github.com/drop-target-pinball/spin/coroutine"

type Event interface {
	coroutine.Selector
}

type EndOfBallEvent struct {
	Player int
	Ball   int
}

func (e EndOfBallEvent) Key() interface{} {
	return e
}

type EndOfGameEvent struct{}

func (e EndOfGameEvent) Key() interface{} {
	return e
}

type GameOverEvent struct{}

func (e GameOverEvent) Key() interface{} {
	return e
}

type Message struct {
	ID string
}

func (e Message) Key() interface{} {
	return e.ID
}

type PlayerAddedEvent struct {
	Player int
}

func (e PlayerAddedEvent) Key() interface{} {
	return "PlayerAddedEvent"
}

type ScriptStartedEvent struct {
	ID string
}

func (e ScriptStartedEvent) Key() interface{} {
	return e
}

type ScriptStoppedEvent struct {
	ID string
}

func (e ScriptStoppedEvent) Key() interface{} {
	return e
}

type StartOfBallEvent struct {
	Player     int
	Ball       int
	ShootAgain bool
}

func (e StartOfBallEvent) Key() interface{} {
	return e
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
