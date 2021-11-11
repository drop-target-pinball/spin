package system

import (
	"log"
	"path"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioSDL struct {
	eng    *spin.Engine
	music  map[string]*mix.Music
	speech map[string]*mix.Chunk
	sound  map[string]*mix.Chunk
}

func NewAudioSDL(eng *spin.Engine) *AudioSDL {
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
	sys := &AudioSDL{
		eng:    eng,
		music:  make(map[string]*mix.Music),
		speech: make(map[string]*mix.Chunk),
		sound:  make(map[string]*mix.Chunk),
	}
	eng.RegisterActionHandler(sys)
	return sys
}

func (s *AudioSDL) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterMusic:
		s.registerMusic(act)
	case spin.RegisterSound:
		s.registerSound(act)
	case spin.RegisterSpeech:
		s.registerSpeech(act)
	case spin.PlayMusic:
		s.playMusic(act)
	case spin.PlaySpeech:
		s.playSpeech(act)
	case spin.PlaySound:
		s.playSound(act)
	case spin.StopAudio:
		s.stopAudio(act)
	case spin.StopMusic:
		s.stopMusic(act)
	case spin.StopSpeech:
		s.stopSpeech(act)
	case spin.VolumeMusic:
		s.volumeMusic(act)
	}
}

func (s *AudioSDL) registerMusic(a spin.RegisterMusic) {
	m, err := mix.LoadMUS(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.music[a.ID] = m
}

func (s *AudioSDL) registerSound(a spin.RegisterSound) {
	snd, err := mix.LoadWAV(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.sound[a.ID] = snd
}

func (s *AudioSDL) registerSpeech(a spin.RegisterSpeech) {
	sp, err := mix.LoadWAV(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.speech[a.ID] = sp
}

func (s *AudioSDL) playMusic(a spin.PlayMusic) {
	m, ok := s.music[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	m.Play(1)
	if a.Vol == 0 {
		mix.VolumeMusic(mix.MAX_VOLUME)
	} else {
		mix.VolumeMusic(a.Vol)
	}
}

func (s *AudioSDL) playSound(a spin.PlaySound) {
	sp, ok := s.sound[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	sp.Play(-1, 0)
}

func (s *AudioSDL) playSpeech(a spin.PlaySpeech) {
	sp, ok := s.speech[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	sp.Play(0, 0)
}

func (s *AudioSDL) stopAudio(a spin.StopAudio) {
	mix.HaltMusic()
	mix.HaltChannel(-1)
}

func (s *AudioSDL) stopMusic(a spin.StopMusic) {
	mix.HaltMusic()
}

func (s *AudioSDL) stopSpeech(a spin.StopSpeech) {
	mix.HaltChannel(0)
}

func (s *AudioSDL) volumeMusic(a spin.VolumeMusic) {
	prev := mix.VolumeMusic(-1)
	vol := prev
	if a.Set == 0 && a.Add == 0 && a.Mul == 0 {
		vol = 0
	}
	if a.Set != 0 {
		vol = a.Set
	}
	if a.Mul != 0 {
		vol = int(float64(vol) * a.Mul)
	}
	if a.Add != 0 {
		vol += a.Add
	}
	if vol < 0 {
		vol = 0

	}
	mix.VolumeMusic(vol)
}
