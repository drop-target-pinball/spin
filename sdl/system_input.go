package sdl

import (
	"log"
	"os"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type key struct {
	key sdl.Keycode
	mod uint16
}

type handler struct {
	eventDown  spin.Event
	actionDown spin.Action
}

type inputSystem struct {
	eng     *spin.Engine
	keys    map[key]handler
	buttons map[sdl.GameControllerButton]handler
}

func RegisterInputSystem(eng *spin.Engine) {
	if err := sdl.InitSubSystem(sdl.INIT_JOYSTICK); err != nil {
		log.Fatalf("unable to initialize joystick subsystem: %v", err)
	}
	if err := sdl.InitSubSystem(sdl.INIT_GAMECONTROLLER); err != nil {
		log.Fatalf("unable to initialize game controller subsystem: %v", err)
	}

	s := &inputSystem{
		eng:     eng,
		keys:    make(map[key]handler),
		buttons: make(map[sdl.GameControllerButton]handler),
	}
	eng.RegisterServer(s)
	eng.RegisterActionHandler(s)
}

func (s *inputSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case RegisterButton:
		s.registerButton(act)
	case RegisterKey:
		s.registerKey(act)
	}
}

func (s *inputSystem) registerKey(act RegisterKey) {
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
	v := handler{eventDown: act.EventDown}
	s.keys[k] = v
}

func (s *inputSystem) registerButton(act RegisterButton) {
	button := sdl.GameControllerGetButtonFromString(act.Button)
	if button == sdl.CONTROLLER_BUTTON_INVALID {
		spin.Warn("unrecognized button: %v", act.Button)
		return
	}
	v := handler{actionDown: act.ActionDown}
	s.buttons[button] = v
}

func (s *inputSystem) Service(_ time.Time) {
	for e := sdl.PollEvent(); e != nil; e = sdl.PollEvent() {
		switch event := e.(type) {
		case *sdl.KeyboardEvent:
			s.handleKey(event)
		case *sdl.ControllerButtonEvent:
			s.handleControllerButton(event)
		case *sdl.ControllerDeviceEvent:
			s.handleControllerDevice(event)
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

func (s *inputSystem) handleControllerButton(evt *sdl.ControllerButtonEvent) {
	handler, ok := s.buttons[sdl.GameControllerButton(evt.Button)]
	if !ok {
		return
	}
	if evt.Type == sdl.CONTROLLERBUTTONDOWN {
		if handler.actionDown != nil {
			s.eng.Do(handler.actionDown)
		}
	}
}

func (s *inputSystem) handleControllerDevice(evt *sdl.ControllerDeviceEvent) {
	if evt.Type == sdl.CONTROLLERDEVICEADDED {
		if evt.Which == 0 {
			gc := sdl.GameControllerOpen(0)
			if gc == nil {
				spin.Error("unable to open game controller")
			}
			spin.Log("added controller: %v", gc.Name())
		}
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
