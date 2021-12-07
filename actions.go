package spin

type Action interface {
	action()
}

type AddPlayer struct {
}

type AdvanceGame struct {
}

type Debug struct {
	ID string
}

type DriverOn struct {
	ID string
}

type DriverOff struct {
	ID string
}

type DriverPulse struct {
	ID   string
	Time int // millseconds
}

type DriverPWM struct {
	ID      string
	TimeOn  int // milliseconds
	TimeOff int // milliseconds
}

type FadeOutMusic struct {
	Time int // milliseconds
}

type AwardScore struct {
	Val int
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

type RegisterCoil struct {
	ID      string
	Address interface{}
}

type RegisterDisplay struct {
	ID      string
	Display Display
}

type RegisterFlasher struct {
	ID      string
	Address interface{}
}

type RegisterFont struct {
	ID   string
	Path string
	Size int
}

type RegisterLamp struct {
	ID      string
	Address interface{}
}

type RegisterMagnet struct {
	ID      string
	Address interface{}
}

type RegisterMotor struct {
	ID      string
	Address interface{}
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

type RegisterSwitch struct {
	ID      string
	Address interface{}
	NC      bool
}

type SetScore struct {
	Val int
}

type StopAudio struct {
}

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

func (AddPlayer) action()       {}
func (AdvanceGame) action()     {}
func (AwardScore) action()      {}
func (Debug) action()           {}
func (DriverOn) action()        {}
func (DriverOff) action()       {}
func (DriverPulse) action()     {}
func (DriverPWM) action()       {}
func (FadeOutMusic) action()    {}
func (PlayMusic) action()       {}
func (PlayScript) action()      {}
func (PlaySound) action()       {}
func (PlaySpeech) action()      {}
func (RegisterCoil) action()    {}
func (RegisterDisplay) action() {}
func (RegisterFlasher) action() {}
func (RegisterFont) action()    {}
func (RegisterKey) action()     {}
func (RegisterLamp) action()    {}
func (RegisterMagnet) action()  {}
func (RegisterMotor) action()   {}
func (RegisterMusic) action()   {}
func (RegisterScript) action()  {}
func (RegisterSound) action()   {}
func (RegisterSpeech) action()  {}
func (RegisterSwitch) action()  {}
func (SetScore) action()        {}
func (StopAudio) action()       {}
func (StopMusic) action()       {}
func (StopScript) action()      {}
func (StopSpeech) action()      {}
func (VolumeMusic) action()     {}

func registerActions(e *Engine) {
	e.RegisterAction(AddPlayer{})
	e.RegisterAction(AdvanceGame{})
	e.RegisterAction(AwardScore{})
	e.RegisterAction(Debug{})
	e.RegisterAction(DriverOn{})
	e.RegisterAction(DriverOff{})
	e.RegisterAction(DriverPulse{})
	e.RegisterAction(DriverPWM{})
	e.RegisterAction(FadeOutMusic{})
	e.RegisterAction(PlayMusic{})
	e.RegisterAction(PlayScript{})
	e.RegisterAction(PlaySound{})
	e.RegisterAction(PlaySpeech{})
	e.RegisterAction(RegisterMusic{})
	e.RegisterAction(RegisterSound{})
	e.RegisterAction(RegisterSpeech{})
	e.RegisterAction(RegisterSwitch{})
	e.RegisterAction(SetScore{})
	e.RegisterAction(StopAudio{})
	e.RegisterAction(StopMusic{})
	e.RegisterAction(StopScript{})
	e.RegisterAction(StopSpeech{})
	e.RegisterAction(VolumeMusic{})
}
