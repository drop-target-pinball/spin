package sdl

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

type inputSystem struct {
	eng  *spin.Engine
	keys map[key]keyHandler
}

func RegisterInputSystem(eng *spin.Engine) {
	s := &inputSystem{
		eng:  eng,
		keys: make(map[key]keyHandler),
	}
	eng.RegisterServer(s)
	eng.RegisterActionHandler(s)
}

func (s *inputSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterKey:
		s.registerKey(act)
	}
}

func (s *inputSystem) registerKey(act spin.RegisterKey) {
	keycode := sdl.GetKeyFromName(act.Key)
	modcode, ok := uint16(0), false
	if keycode == 0 {
		spin.Warn("unrecognized key: %v", act.Key)
		return
	}
	if act.Mod != "" {
		modcode, ok = sdlModFromName[act.Mod]
		if !ok {
			spin.Warn("unrecognized key modifier: %v", act.Mod)
			return
		}
	}
	k := key{key: keycode, mod: modcode}
	v := keyHandler{eventDown: act.EventDown}
	s.keys[k] = v
}

func (s *inputSystem) Service() {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.KeyboardEvent:
			s.handleKey(event)
		case *sdl.QuitEvent:
			os.Exit(0)
		}
	}
}

func (s *inputSystem) handleKey(kbe *sdl.KeyboardEvent) {
	if kbe.Repeat != 0 {
		return
	}
	key := key{key: kbe.Keysym.Sym, mod: getSingleMod(kbe.Keysym.Mod)}
	handlers, ok := s.keys[key]
	if !ok {
		return
	}
	if kbe.Type == sdl.KEYDOWN && handlers.eventDown != nil {
		s.eng.Post(handlers.eventDown)
	}
}

var sdlModFromName = map[string]uint16{
	"control": sdl.KMOD_CTRL,
	"shift":   sdl.KMOD_SHIFT,
	"alt":     sdl.KMOD_ALT,
}

func getSingleMod(mod uint16) uint16 {
	switch {
	case mod&sdl.KMOD_SHIFT != 0:
		return sdl.KMOD_SHIFT
	case mod&sdl.KMOD_CTRL != 0:
		return sdl.KMOD_CTRL
	case mod&sdl.KMOD_ALT != 0:
		return sdl.KMOD_ALT
	}
	return 0
}
