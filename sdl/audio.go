package sdl

import (
	"log"
	"path"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

const chanSpeech = 0

type audio struct {
	id       string
	priority int
	notify   bool
}

type audioSystem struct {
	eng    *spin.Engine
	music  map[string]*mix.Music
	speech map[string]*mix.Chunk
	sound  map[string]*mix.Chunk

	musicPlaying  string
	speechPlaying audio
	soundsPlaying []string

	musicNotify bool
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
	nChan := mix.AllocateChannels(-1)

	s := &audioSystem{
		eng:           eng,
		music:         make(map[string]*mix.Music),
		speech:        make(map[string]*mix.Chunk),
		sound:         make(map[string]*mix.Chunk),
		soundsPlaying: make([]string, nChan),
	}

	mix.ChannelFinished(s.channelFinished)
	mix.HookMusicFinished(s.musicFinished)

	eng.RegisterActionHandler(s)
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
	case spin.StopSound:
		s.stopSound(act)
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
	loops := a.Loops
	if loops == 0 {
		loops = -1
	}
	m.Play(loops)
	s.musicPlaying = a.ID
	s.musicNotify = a.Notify

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
	channel, err := sp.Play(-1, a.Loops)
	if err != nil {
		log.Panic(err)
	}
	s.soundsPlaying[channel] = a.ID
}

func (s *audioSystem) playSpeech(a spin.PlaySpeech) {
	sp, ok := s.speech[a.ID]
	if !ok {
		spin.Warn("%v not found", a.ID)
		return
	}
	if s.speechPlaying.id != "" && s.speechPlaying.priority > a.Priority {
		spin.Info("speech does not have priority: %v", a.ID)
		return
	}
	s.speechPlaying = audio{
		id:       a.ID,
		priority: a.Priority,
		notify:   a.Notify,
	}
	sp.Play(chanSpeech, 0)
}

func (s *audioSystem) stopAudio(a spin.StopAudio) {
	mix.HaltMusic()
	mix.HaltChannel(-1)
	s.musicPlaying = ""
	s.speechPlaying = audio{}
}

func (s *audioSystem) stopMusic(a spin.StopMusic) {
	if a.Any || a.ID == s.musicPlaying {
		mix.HaltMusic()
		s.musicPlaying = ""
	}
}

func (s *audioSystem) stopSound(a spin.StopSound) {
	for channel, id := range s.soundsPlaying {
		if id == a.ID {
			mix.HaltChannel(channel)
			break
		}
	}
}

func (s *audioSystem) stopSpeech(a spin.StopSpeech) {
	if a.ID == "" || a.ID == s.speechPlaying.id {
		mix.HaltChannel(0)
		s.speechPlaying = audio{}
	}
}

func (s *audioSystem) channelFinished(ch int) {
	if ch == chanSpeech {
		s.speechPlaying = audio{}
	}
	if ch == chanSpeech && s.speechPlaying.notify {
		s.eng.Post(spin.SpeechFinishedEvent{})
	}
}

func (s *audioSystem) musicFinished() {
	if s.musicNotify {
		s.eng.Post(spin.MusicFinishedEvent{})
	}
}
