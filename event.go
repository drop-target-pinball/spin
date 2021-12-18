package spin

import "github.com/drop-target-pinball/spin/coroutine"

type Event interface {
	coroutine.Selector
}

type BallDrainEvent struct{}

func (e BallDrainEvent) Key() interface{} {
	return "BallDrainEvent"
}

type BallWillDrainEvent struct{}

func (e BallWillDrainEvent) Key() interface{} {
	return "BallWillDrainEvent"
}

type EndOfBallEvent struct {
	Player int
	Ball   int
}

func (e EndOfBallEvent) Key() interface{} {
	return "EndOfBallEvent"
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

type ModeFinishedEvent struct {
	ID string
}

func (e ModeFinishedEvent) Key() interface{} {
	return e.ID
}

type PlayerAddedEvent struct {
	Player int
}

func (e PlayerAddedEvent) Key() interface{} {
	return "PlayerAddedEvent"
}

type ShotEvent struct {
	ID string
}

func (e ShotEvent) Key() interface{} {
	return e.ID
}

type StartOfBallEvent struct {
	Player     int
	Ball       int
	ShootAgain bool
}

func (e StartOfBallEvent) Key() interface{} {
	return "StartOfBallEvent"
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

type Done struct{}

func (e Done) Key() interface{} {
	return "Done"
}

func registerEvents(e *Engine) {
	e.RegisterEvent(Message{})
	e.RegisterEvent(ShotEvent{})
	e.RegisterEvent(SwitchEvent{})
}
