package sdl

import (
	"log"
	"path"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type audioSystem struct {
	eng    *spin.Engine
	music  map[string]*mix.Music
	speech map[string]*mix.Chunk
	sound  map[string]*mix.Chunk

	musicPlaying  string
	speechPlaying string
}

func RegisterAudioSystem(eng *spin.Engine) {
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
	sys := &audioSystem{
		eng:    eng,
		music:  make(map[string]*mix.Music),
		speech: make(map[string]*mix.Chunk),
		sound:  make(map[string]*mix.Chunk),
	}
	eng.RegisterActionHandler(sys)
}

func (s *audioSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.FadeOutMusic:
		s.fadeOutMusic(act)
	case spin.MusicVolume:
		s.musicVolume(act)
	case spin.PlayMusic:
		s.playMusic(act)
	case spin.PlaySpeech:
		s.playSpeech(act)
	case spin.PlaySound:
		s.playSound(act)
	case spin.RegisterMusic:
		s.registerMusic(act)
	case spin.RegisterSound:
		s.registerSound(act)
	case spin.RegisterSpeech:
		s.registerSpeech(act)
	case spin.StopAudio:
		s.stopAudio(act)
	case spin.StopMusic:
		s.stopMusic(act)
	case spin.StopSpeech:
		s.stopSpeech(act)
	}
}

func (s *audioSystem) fadeOutMusic(act spin.FadeOutMusic) {
	mix.FadeOutMusic(act.Time)
}

func (s *audioSystem) registerMusic(a spin.RegisterMusic) {
	m, err := mix.LoadMUS(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.music[a.ID] = m
}

func (s *audioSystem) registerSound(a spin.RegisterSound) {
	snd, err := mix.LoadWAV(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.sound[a.ID] = snd
}

func (s *audioSystem) registerSpeech(a spin.RegisterSpeech) {
	sp, err := mix.LoadWAV(path.Join(spin.AssetDir, a.Path))
	if err != nil {
		spin.Warn("%v: %v", a.ID, err)
		return
	}
	s.speech[a.ID] = sp
}

func (s *audioSystem) musicVolume(a spin.MusicVolume) {
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

func (s *audioSystem) playMusic(a spin.PlayMusic) {
	m, ok := s.music[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	m.Play(-1)
	s.musicPlaying = a.ID
	if a.Vol == 0 {
		mix.VolumeMusic(mix.MAX_VOLUME)
	} else {
		mix.VolumeMusic(a.Vol)
	}
}

func (s *audioSystem) playSound(a spin.PlaySound) {
	sp, ok := s.sound[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	sp.Play(-1, 0)
}

func (s *audioSystem) playSpeech(a spin.PlaySpeech) {
	sp, ok := s.speech[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	s.speechPlaying = a.ID
	sp.Play(0, 0)
}

func (s *audioSystem) stopAudio(a spin.StopAudio) {
	mix.HaltMusic()
	mix.HaltChannel(-1)
	s.musicPlaying = ""
	s.speechPlaying = ""
}

func (s *audioSystem) stopMusic(a spin.StopMusic) {
	if a.Any || a.ID == s.musicPlaying {
		mix.HaltMusic()
		s.musicPlaying = ""
	}
}

func (s *audioSystem) stopSpeech(a spin.StopSpeech) {
	if a.Any || a.ID == s.speechPlaying {
		mix.HaltChannel(0)
		s.speechPlaying = ""
	}
}
