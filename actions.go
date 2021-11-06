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
func (RegisterMusic) action()  {}
func (RegisterScript) action() {}
func (RegisterSound) action()  {}
func (RegisterSpeech) action() {}
func (StopAudio) action()      {}
func (StopMusic) action()      {}
func (StopScript) action()     {}
func (StopSpeech) action()     {}
