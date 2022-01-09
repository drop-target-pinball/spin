package spin

type Event interface {
	Key() interface{}
}

type AdvanceEvent struct {
	ID string
}

func (e AdvanceEvent) Key() interface{} {
	return e.ID
}

type BallAddedEvent struct {
	BallsInPlay int
}

func (e BallAddedEvent) Key() interface{} {
	return BallAddedEvent{}
}

type BallDrainEvent struct {
	BallsInPlay int
}

func (e BallDrainEvent) Key() interface{} {
	return BallDrainEvent{}
}

type BallWillDrainEvent struct{}

func (e BallWillDrainEvent) Key() interface{} {
	return BallDrainEvent{}
}

type EndOfBallEvent struct {
	Player int
	Ball   int
}

func (e EndOfBallEvent) Key() interface{} {
	return EndOfBallEvent{}
}

type EndOfGameEvent struct{}

func (e EndOfGameEvent) Key() interface{} {
	return EndOfGameEvent{}
}

type GameOverEvent struct{}

func (e GameOverEvent) Key() interface{} {
	return GameOverEvent{}
}

type Message struct {
	ID string
}

func (e Message) Key() interface{} {
	return Message{ID: e.ID}
}

type MusicFinishedEvent struct{}

func (e MusicFinishedEvent) Key() interface{} {
	return MusicFinishedEvent{}
}

type PlayerAddedEvent struct {
	Player int
}

func (e PlayerAddedEvent) Key() interface{} {
	return PlayerAddedEvent{}
}

type ScriptFinishedEvent struct {
	ID string
}

func (e ScriptFinishedEvent) Key() interface{} {
	return ScriptFinishedEvent{ID: e.ID}
}

type ShotEvent struct {
	ID string
}

func (e ShotEvent) Key() interface{} {
	return ShotEvent{ID: e.ID}
}

type SoundFinishedEvent struct {
	ID string
}

func (e SoundFinishedEvent) Key() interface{} {
	return SoundFinishedEvent{ID: e.ID}
}

type SpeechFinishedEvent struct {
	ID string
}

func (e SpeechFinishedEvent) Key() interface{} {
	return SpeechFinishedEvent{}
}

type StartOfBallEvent struct {
	Player     int
	Ball       int
	ShootAgain bool
}

func (e StartOfBallEvent) Key() interface{} {
	return StartOfBallEvent{}
}

type SwitchEvent struct {
	ID       string
	Released bool
}

func (e SwitchEvent) Key() interface{} {
	return SwitchEvent{ID: e.ID, Released: e.Released}
}

type TimeoutEvent struct {
	ID string
}

func (e TimeoutEvent) Key() interface{} {
	return e.ID
}

type Done struct{}

func (e Done) Key() interface{} {
	return Done{}
}

func registerEvents(e *Engine) {
	e.RegisterEvent(BallDrainEvent{})
	e.RegisterEvent(Message{})
	e.RegisterEvent(ShotEvent{})
	e.RegisterEvent(SwitchEvent{})
}
