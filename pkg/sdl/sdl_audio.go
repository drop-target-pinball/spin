//go:build sdl

package sdl

import (
	"fmt"
	"reflect"

	"github.com/drop-target-pinball/spin/v2"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioFactory struct{}

func (a AudioFactory) ID() string {
	return "sdl_mixer"
}

func (a AudioFactory) NewDevice(conf any) (spin.Device, error) {
	c, ok := conf.(spin.AudioDevice)
	if !ok {
		return nil, fmt.Errorf("invalid config for %v, got %v", a.ID(), reflect.TypeOf(conf).Name())
	}
	return &AudioDevice{
		id:     c.ID,
		name:   spin.DeviceName(a.ID(), c.ID),
		config: c,
		sounds: make(map[string]*mix.Chunk),
	}, nil
}

type AudioDevice struct {
	handler AudioFactory
	id      string
	name    string
	config  spin.AudioDevice
	queue   *spin.QueueClient
	sounds  map[string]*mix.Chunk
}

func (d *AudioDevice) ID() string {
	return d.id
}

func (d *AudioDevice) Name() string {
	return d.name
}

func (d *AudioDevice) Init(e *spin.Engine) bool {
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

	d.queue = e.NewQueueClient()
	d.queue.Reset()

	return true
}

func (d *AudioDevice) Process(e *spin.Engine) bool {
	msg, err := d.queue.Read()
	if err != nil {
		e.Error(err)
	}
	switch m := msg.(type) {
	case spin.Load:
		d.load(e, m)
	case spin.Play:
		d.play(m)
	}
	return true
}

func (d *AudioDevice) load(e *spin.Engine, load spin.Load) {
	for _, audio := range e.Config.Audio {
		if audio.Device != d.id {
			continue
		}
		if audio.Module != load.ID {
			continue
		}
		e.Debug("%v: loading %v %v: %v", d.name, audio.Type, audio.ID, audio.File)
		fullPath := e.PathTo(audio.File)
		snd, err := mix.LoadWAV(fullPath)
		if err != nil {
			e.Warn("unable to load %v: %v", fullPath, err)
			continue
		}
		d.sounds[audio.ID] = snd
	}
}

func (d *AudioDevice) play(msg spin.Play) {
	sound, ok := d.sounds[msg.ID]
	if !ok {
		return
	}
	sound.Play(-1, 0)
}
