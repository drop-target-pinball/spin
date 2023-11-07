package sdl

import (
	"github.com/drop-target-pinball/spin/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	AudioHandlerName = "sdl"
)

type AudioHandler struct {
	id     string
	config spin.AudioDevice
	queue  *spin.QueueClient
	sounds map[string]*mix.Chunk
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

func init() {
	spin.AddNewDeviceFunc(AudioHandlerName, NewAudioDevice)
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
	h.load(e)

	h.queue = e.NewQueueClient()
	h.queue.Reset()

	return true
}

func (h *AudioHandler) Process(e *spin.Engine) {
	msg, err := h.queue.Read()
	if err != nil {
		e.Error(err)
	}
	switch m := msg.(type) {
	case spin.Play:
		h.play(m)
	}
}

func (h *AudioHandler) load(e *spin.Engine) {
	for _, audio := range e.Config.Audio {
		if audio.Device != h.id {
			continue
		}
		if audio.Module != e.Module {
			continue
		}

		e.Log("loading %v", audio.File)
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
