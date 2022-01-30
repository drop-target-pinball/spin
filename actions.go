package spin

type Action interface {
	action()
}

type AddBall struct {
	N int
}

type AddPlayer struct{}

type AdvanceGame struct{}

type AutoPulseOn struct {
	ID string
}

type AutoPulseOff struct {
	ID string
}

type AwardScore struct {
	Val int
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
	ID  string
	On  int // milliseconds
	Off int // milliseconds
}

type FadeOutMusic struct {
	Time int // milliseconds
}

type FlippersOn struct {
	FlipperIDs []string
}

type FlippersOff struct {
	FlipperIDs []string
}

type MusicVolume struct {
	Set int
	Add int
	Mul float64
}

type PlayMusic struct {
	ID     string
	Vol    int
	Loops  int
	Notify bool
}

type PlayScript struct {
	ID string
}

type PlaySound struct {
	ID     string
	Vol    int
	Loop   bool
	Repeat int
	Notify bool
	Duck   float64
}

type PlaySpeech struct {
	ID       string
	Vol      int
	Priority int
	Notify   bool
	Duck     float64
}

type RegisterAutoPulse struct {
	ID         string
	SwitchAddr interface{}
	CoilAddr   interface{}
	Time       int // millseconds
}

type RegisterCoil struct {
	ID   string
	Addr interface{}
}

type RegisterDisplay struct {
	ID      string
	Display Display
}

type RegisterFlasher struct {
	ID   string
	Addr interface{}
}

type RegisterFlipper struct {
	ID            string
	SwitchAddr    interface{}
	PowerCoilAddr interface{}
	HoldCoilAddr  interface{}
}

type RegisterFont struct {
	ID   string
	Path string
	Size int
}

type RegisterLamp struct {
	ID   string
	Addr interface{}
}

type RegisterMagnet struct {
	ID   string
	Addr interface{}
}

type RegisterMotor struct {
	ID   string
	Addr interface{}
}

type RegisterMusic struct {
	ID   string
	Path string
}

type RegisterScript struct {
	ID     string
	Script ScriptFn
	Group  string
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
	ID   string
	Addr interface{}
	NC   bool
}

type SetScore struct {
	Val int
}

type SetVar struct {
	Vars string
	ID   string
	Val  string
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

type StopScriptGroup struct {
	ID string
}

type StopSound struct {
	ID string
}

type StopSpeech struct {
	ID string
}

func (AddBall) action()           {}
func (AddPlayer) action()         {}
func (AdvanceGame) action()       {}
func (AutoPulseOn) action()       {}
func (AutoPulseOff) action()      {}
func (AwardScore) action()        {}
func (Debug) action()             {}
func (DriverOn) action()          {}
func (DriverOff) action()         {}
func (DriverPulse) action()       {}
func (DriverPWM) action()         {}
func (FadeOutMusic) action()      {}
func (FlippersOn) action()        {}
func (FlippersOff) action()       {}
func (MusicVolume) action()       {}
func (PlayMusic) action()         {}
func (PlayScript) action()        {}
func (PlaySound) action()         {}
func (PlaySpeech) action()        {}
func (RegisterAutoPulse) action() {}
func (RegisterCoil) action()      {}
func (RegisterDisplay) action()   {}
func (RegisterFlasher) action()   {}
func (RegisterFlipper) action()   {}
func (RegisterFont) action()      {}
func (RegisterLamp) action()      {}
func (RegisterMagnet) action()    {}
func (RegisterMotor) action()     {}
func (RegisterMusic) action()     {}
func (RegisterScript) action()    {}
func (RegisterSound) action()     {}
func (RegisterSpeech) action()    {}
func (RegisterSwitch) action()    {}
func (SetScore) action()          {}
func (SetVar) action()            {}
func (StopAudio) action()         {}
func (StopMusic) action()         {}
func (StopScript) action()        {}
func (StopScriptGroup) action()   {}
func (StopSound) action()         {}
func (StopSpeech) action()        {}

func registerActions(e *Engine) {
	e.RegisterAction(AddBall{})
	e.RegisterAction(AddPlayer{})
	e.RegisterAction(AdvanceGame{})
	e.RegisterAction(AwardScore{})
	e.RegisterAction(Debug{})
	e.RegisterAction(DriverOn{})
	e.RegisterAction(DriverOff{})
	e.RegisterAction(DriverPulse{})
	e.RegisterAction(DriverPWM{})
	e.RegisterAction(FadeOutMusic{})
	e.RegisterAction(FlippersOn{})
	e.RegisterAction(FlippersOff{})
	e.RegisterAction(MusicVolume{})
	e.RegisterAction(PlayMusic{})
	e.RegisterAction(PlayScript{})
	e.RegisterAction(PlaySound{})
	e.RegisterAction(PlaySpeech{})
	e.RegisterAction(RegisterMusic{})
	e.RegisterAction(RegisterSound{})
	e.RegisterAction(RegisterSpeech{})
	e.RegisterAction(SetScore{})
	e.RegisterAction(SetVar{})
	e.RegisterAction(StopAudio{})
	e.RegisterAction(StopMusic{})
	e.RegisterAction(StopScript{})
	e.RegisterAction(StopScriptGroup{})
	e.RegisterAction(StopSound{})
	e.RegisterAction(StopSpeech{})
}
