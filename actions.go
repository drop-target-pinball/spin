package spin

type Action interface {
	action()
}

type PlayMusic struct {
	ID  string
	Vol int
}

type PlayScript struct {
	ID string
}

type PlaySound struct {
	ID string
}

type PlaySpeech struct {
	ID  string
	Vol int
}

type RegisterDisplay struct {
	ID      string
	Display Display
}

type RegisterFont struct {
	ID   string
	Path string
	Size int
}

type RegisterKey struct {
	Key       string
	Mod       string
	EventDown Event
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

type Score struct {
	Add int
}

type StopAudio struct{}

type StopMusic struct {
	ID  string
	Any bool
}

type StopScript struct {
	ID string
}

type StopSpeech struct {
	ID  string
	Any bool
}

type VolumeMusic struct {
	Set int
	Add int
	Mul float64
}

func (PlayMusic) action()       {}
func (PlayScript) action()      {}
func (PlaySound) action()       {}
func (PlaySpeech) action()      {}
func (RegisterDisplay) action() {}
func (RegisterFont) action()    {}
func (RegisterKey) action()     {}
func (RegisterMusic) action()   {}
func (RegisterScript) action()  {}
func (RegisterSound) action()   {}
func (RegisterSpeech) action()  {}
func (Score) action()           {}
func (StopAudio) action()       {}
func (StopMusic) action()       {}
func (StopScript) action()      {}
func (StopSpeech) action()      {}
func (VolumeMusic) action()     {}

func registerActions(e *Engine) {
	e.RegisterAction(PlayMusic{})
	e.RegisterAction(PlayScript{})
	e.RegisterAction(PlaySound{})
	e.RegisterAction(PlaySpeech{})
	e.RegisterAction(Score{})
	e.RegisterAction(RegisterMusic{})
	e.RegisterAction(RegisterSound{})
	e.RegisterAction(RegisterSpeech{})
	e.RegisterAction(Score{})
	e.RegisterAction(StopAudio{})
	e.RegisterAction(StopMusic{})
	e.RegisterAction(StopScript{})
	e.RegisterAction(StopSpeech{})
	e.RegisterAction(VolumeMusic{})
}
