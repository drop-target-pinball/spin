//go:build sdl

package sdl

import (
	"fmt"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioHandler struct {
	id     string
	config spin.AudioDevice
	queue  *spin.QueueClient
	sounds map[string]*mix.Chunk
}

func (a AudioHandler) FormatID() string {
	if a.id == "" {
		return AudioHandlerName
	}
	return fmt.Sprintf("%v(%v)", AudioHandlerName, a.id)
}

func NewAudioDevice(conf any) (spin.Device, bool) {
	c, ok := conf.(spin.AudioDevice)
	if !ok {
		return nil, false
	}
	if c.Handler != AudioHandlerName {
		return nil, false
	}
	return &AudioHandler{
		id:     c.ID,
		config: c,
		sounds: make(map[string]*mix.Chunk),
	}, true
}

func (h *AudioHandler) Init(e *spin.Engine) bool {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		e.Warn("unable to initialize audio: %v", err)
		return false
	}
	if err := mix.Init(mix.INIT_OGG); err != nil {
		e.Warn("unable to initialize audio: %v", err)
		return false
	}

	if err := mix.OpenAudio(
		mix.DEFAULT_FREQUENCY,
		mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS,
		mix.DEFAULT_CHUNKSIZE,
	); err != nil {
		e.Warn("unable to initialize audio: %v", err)
		return false
	}

	// mix.ReserveChannels(1)
	// nChan := mix.AllocateChannels(-1)
	// h.load(e)

	h.queue = e.NewQueueClient()
	h.queue.Reset()

	return true
}

func (h *AudioHandler) Process(e *spin.Engine) bool {
	msg, err := h.queue.Read()
	if err != nil {
		e.Error(err)
	}
	switch m := msg.(type) {
	case spin.Load:
		h.load(e, m)
	case spin.Play:
		h.play(m)
	}
	return true
}

func (h *AudioHandler) load(e *spin.Engine, load spin.Load) {
	for _, audio := range e.Config.Audio {
		if audio.Device != h.id {
			continue
		}
		if audio.Module != load.ID {
			continue
		}
		e.Debug("%v: loading %v %v: %v", h.FormatID(), audio.Type, audio.ID, audio.File)
		fullPath := e.PathTo(audio.File)
		snd, err := mix.LoadWAV(fullPath)
		if err != nil {
			e.Warn("unable to load %v: %v", fullPath, err)
			continue
		}
		h.sounds[audio.ID] = snd
	}
}

func (h *AudioHandler) play(msg spin.Play) {
	snd, ok := h.sounds[msg.ID]
	if !ok {
		return
	}
	snd.Play(-1, 0)
}
