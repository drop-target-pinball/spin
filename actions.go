package spin

type Action interface {
	action()
}

type PlayMusic struct {
	ID string
}

type PlayScript struct {
	ID string
}

type PlaySound struct {
	ID string
}

type PlaySpeech struct {
	ID string
}

type RegisterAction struct {
	Action Action
}

type RegisterEvent struct {
	Event Event
}

type RegisterMusic struct {
	ID   string
	Path string
}

type RegisterScript struct {
	ID     string
	Script Script
}

type RegisterSound struct {
	ID   string
	Path string
}

type RegisterSpeech struct {
	ID   string
	Path string
}

type StopAudio struct{}

type StopMusic struct{}

type StopScript struct {
	ID string
}

type StopSpeech struct{}

func (PlayMusic) action()      {}
func (PlayScript) action()     {}
func (PlaySound) action()      {}
func (PlaySpeech) action()     {}
func (RegisterAction) action() {}
func (RegisterEvent) action()  {}
func (RegisterMusic) action()  {}
func (RegisterScript) action() {}
func (RegisterSound) action()  {}
func (RegisterSpeech) action() {}
func (StopAudio) action()      {}
func (StopMusic) action()      {}
func (StopScript) action()     {}
func (StopSpeech) action()     {}

func registerActions(e *Engine) {
	e.RegisterAction(PlayMusic{})
	e.RegisterAction(PlayScript{})
	e.RegisterAction(PlaySound{})
	e.RegisterAction(PlaySpeech{})
	e.RegisterAction(RegisterMusic{})
	e.RegisterAction(RegisterSound{})
	e.RegisterAction(RegisterSpeech{})
	e.RegisterAction(StopAudio{})
	e.RegisterAction(StopMusic{})
	e.RegisterAction(StopScript{})
	e.RegisterAction(StopSpeech{})
}
