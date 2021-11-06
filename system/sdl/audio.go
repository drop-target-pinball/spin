package sdl

import (
	"log"
	"os"
	"path"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var assetDir = os.Getenv("SPIN_ASSET_DIR")

type AudioSystem struct {
	eng    *spin.Engine
	music  map[string]*mix.Music
	speech map[string]*mix.Chunk
	sound  map[string]*mix.Chunk
}

func NewAudioSystem(eng *spin.Engine) *AudioSystem {
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		log.Fatalf("unable to initialize audio: %v", err)
	}
	if err := mix.Init(mix.INIT_OGG); err != nil {
		log.Fatalf("unable to initialize mixer: %v", err)
	}
	mix.OpenAudio(
		mix.DEFAULT_FREQUENCY,
		mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS,
		mix.DEFAULT_CHUNKSIZE,
	)
	mix.ReserveChannels(1)
	sys := &AudioSystem{
		eng:    eng,
		music:  make(map[string]*mix.Music),
		speech: make(map[string]*mix.Chunk),
		sound:  make(map[string]*mix.Chunk),
	}
	eng.RegisterActionHandler(sys)
	return sys
}

func (s *AudioSystem) HandleAction(a spin.Action) {
	switch action := a.(type) {
	case spin.RegisterMusic:
		s.registerMusic(action)
	case spin.RegisterSound:
		s.registerSound(action)
	case spin.RegisterSpeech:
		s.registerSpeech(action)
	case spin.PlayMusic:
		s.playMusic(action)
	case spin.PlaySpeech:
		s.playSpeech(action)
	case spin.PlaySound:
		s.playSound(action)
	case spin.StopAudio:
		s.stopAudio(action)
	case spin.StopMusic:
		s.stopMusic(action)
	case spin.StopSpeech:
		s.stopSpeech(action)
	}
}

func (s *AudioSystem) registerMusic(a spin.RegisterMusic) {
	m, err := mix.LoadMUS(path.Join(assetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.music[a.ID] = m
}

func (s *AudioSystem) registerSound(a spin.RegisterSound) {
	snd, err := mix.LoadWAV(path.Join(assetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.sound[a.ID] = snd
}

func (s *AudioSystem) registerSpeech(a spin.RegisterSpeech) {
	sp, err := mix.LoadWAV(path.Join(assetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.speech[a.ID] = sp
}

func (s *AudioSystem) playMusic(a spin.PlayMusic) {
	m, ok := s.music[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	m.Play(1)
}

func (s *AudioSystem) playSound(a spin.PlaySound) {
	sp, ok := s.sound[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	sp.Play(-1, 0)
}

func (s *AudioSystem) playSpeech(a spin.PlaySpeech) {
	sp, ok := s.speech[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	sp.Play(0, 0)
}

func (s *AudioSystem) stopAudio(a spin.StopAudio) {
	mix.HaltMusic()
	mix.HaltChannel(-1)
}

func (s *AudioSystem) stopMusic(a spin.StopMusic) {
	mix.HaltMusic()
}

func (s *AudioSystem) stopSpeech(a spin.StopSpeech) {
	mix.HaltChannel(0)
}
