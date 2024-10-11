package sdl

import (
	"time"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type Audio struct {
	pin   *spin.Engine
	music map[string]*mix.Music
}

func NewAudio(pin *spin.Engine) *Audio {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		pin.Abort("unable to initialize audio: %v", err)
	}
	if err := mix.Init(mix.INIT_OGG); err != nil {
		pin.Abort("unable to initialize mixer: %v", err)
	}

	if err := mix.OpenAudio(
		mix.DEFAULT_FREQUENCY,
		mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS,
		mix.DEFAULT_CHUNKSIZE,
	); err != nil {
		pin.Abort("unable to initialize audio: %v", err)
	}

	// mix.ReserveChannels(1)
	// nChan := mix.AllocateChannels(-1)

	return &Audio{
		pin:   pin,
		music: make(map[string]*mix.Music),
	}
}

func (a *Audio) Service(t time.Time) {}

func (a *Audio) Handle(m spin.Message) {

}
