package spin

type Action interface {
	action()
}

type PlayMusic struct {
	ID string
}

func (a PlayMusic) String() string {
	return "play-music: " + a.ID
}

type PlayScript struct {
	ID string
}

func (a PlayScript) String() string {
	return "play-script: " + a.ID
}

type PlaySound struct {
	ID string
}

func (a PlaySound) String() string {
	return "play-sound: " + a.ID
}

type PlaySpeech struct {
	ID string
}

func (a PlaySpeech) String() string {
	return "play-speech: " + a.ID
}

type RegisterMusic struct {
	ID   string
	Path string
}

func (a RegisterMusic) String() string {
	return "register-music: " + a.ID
}

type RegisterScript struct {
	ID     string
	Script Script
}

func (a RegisterScript) String() string {
	return "register-script: " + a.ID
}

type RegisterSound struct {
	ID   string
	Path string
}

func (a RegisterSound) String() string {
	return "register-sound: " + a.ID
}

type RegisterSpeech struct {
	ID   string
	Path string
}

func (a RegisterSpeech) String() string {
	return "register-speech: " + a.ID
}

type StopAudio struct{}

func (a StopAudio) String() string {
	return "stop-audio"
}

type StopMusic struct{}

func (a StopMusic) String() string {
	return "stop-music"
}

type StopScript struct {
	ID string
}

func (a StopScript) String() string {
	return "stop-script: " + a.ID
}

type StopSpeech struct{}

func (a StopSpeech) String() string {
	return "stop-speech"
}

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
