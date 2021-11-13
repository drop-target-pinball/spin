package system

import (
	"os"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type key struct {
	key sdl.Keycode
	mod uint16
}

type keyHandler struct {
	eventDown spin.Event
}

type InputSDL struct {
	eng  *spin.Engine
	keys map[key]keyHandler
}

func RegisterInputSDL(eng *spin.Engine) {
	s := &InputSDL{
		eng:  eng,
		keys: make(map[key]keyHandler),
	}
	eng.RegisterServer(s)
	eng.RegisterActionHandler(s)
}

func (s *InputSDL) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterKeySDL:
		s.registerKey(act)
	}
}

func (s *InputSDL) registerKey(act spin.RegisterKeySDL) {
	k := key{key: act.Key, mod: act.Mod}
	v := keyHandler{eventDown: act.EventDown}
	s.keys[k] = v
}

func (s *InputSDL) Service() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.KeyboardEvent:
			s.handleKey(event)
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

func (s *InputSDL) handleKey(kbe *sdl.KeyboardEvent) {
	if kbe.Repeat != 0 {
		return
	}
	key := key{key: kbe.Keysym.Sym, mod: kbe.Keysym.Mod}
	handlers, ok := s.keys[key]
	if !ok {
		return
	}
	if kbe.Type == sdl.KEYDOWN && handlers.eventDown != nil {
		s.eng.Post(handlers.eventDown)
	}
}
