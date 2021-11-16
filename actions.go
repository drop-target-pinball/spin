package spin

import (
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

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

type RegisterDisplaySDL struct {
	ID      string
	Display Display
	Surface *sdl.Surface
	Mutex   *sync.Mutex
}

type RegisterFont struct {
	ID   string
	Path string
	Size int
}

type RegisterKeySDL struct {
	Key       sdl.Keycode
	Mod       uint16
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

func (PlayMusic) action()          {}
func (PlayScript) action()         {}
func (PlaySound) action()          {}
func (PlaySpeech) action()         {}
func (RegisterDisplaySDL) action() {}
func (RegisterFont) action()       {}
func (RegisterKeySDL) action()     {}
func (RegisterMusic) action()      {}
func (RegisterScript) action()     {}
func (RegisterSound) action()      {}
func (RegisterSpeech) action()     {}
func (StopAudio) action()          {}
func (StopMusic) action()          {}
func (StopScript) action()         {}
func (StopSpeech) action()         {}
func (VolumeMusic) action()        {}

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
	e.RegisterAction(VolumeMusic{})
}
