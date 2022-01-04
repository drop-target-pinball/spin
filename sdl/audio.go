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
	origVol  int
	stopped  bool
}

type audioSystem struct {
	eng    *spin.Engine
	music  map[string]*mix.Music
	speech map[string]*mix.Chunk
	sound  map[string]*mix.Chunk

	musicPlaying  string
	speechPlaying audio
	soundsPlaying []audio

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
		soundsPlaying: make([]audio, nChan),
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
	} else if a.Set != 0 {
		vol = a.Set
	} else if a.Mul != 0 {
		vol = int(float64(vol) * a.Mul)
	} else if a.Add != 0 {
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
	loops := 0
	if a.Loop {
		loops = -1
	}
	if a.Repeat > 0 {
		loops = a.Repeat
	}
	channel, err := sp.Play(-1, loops)
	if err != nil {
		log.Panic(err)
	}
	if a.Duck < 0 || a.Duck > 1 {
		spin.Error("invalid duck factor: %v", a.Duck)
		return
	}
	prev := 0
	if a.Duck > 0 {
		prev = mix.VolumeMusic(-1)
		ducked := int(a.Duck * float64(prev))
		mix.VolumeMusic(ducked)
	}
	if a.Vol > 0 {
		sp.Volume(a.Vol)
	}
	s.soundsPlaying[channel] = audio{id: a.ID, notify: a.Notify, origVol: prev}
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
	if a.Duck < 0 || a.Duck > 1 {
		spin.Error("invalid duck factor: %v", a.Duck)
	}
	prev := 0
	if a.Duck > 0 {
		prev = mix.VolumeMusic(-1)
		ducked := int(a.Duck * float64(prev))
		mix.VolumeMusic(ducked)
	}
	s.speechPlaying = audio{
		id:       a.ID,
		priority: a.Priority,
		notify:   a.Notify,
		origVol:  prev,
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
	for channel, playing := range s.soundsPlaying {
		if playing.id == a.ID {
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
	if ch == chanSpeech && s.speechPlaying.notify {
		s.eng.Post(spin.SpeechFinishedEvent{})
		if s.speechPlaying.origVol > 0 {
			mix.VolumeMusic(s.speechPlaying.origVol)
		}
	} else {
		playing := s.soundsPlaying[ch]
		if playing.notify {
			s.eng.Post(spin.SoundFinishedEvent{ID: playing.id})
		}
		if playing.origVol > 0 {
			mix.VolumeMusic(playing.origVol)
		}
	}
	if ch == chanSpeech {
		s.speechPlaying = audio{}
	} else {
		s.soundsPlaying[ch] = audio{}
	}
}

func (s *audioSystem) musicFinished() {
	if s.musicNotify {
		s.eng.Post(spin.MusicFinishedEvent{})
	}
}
