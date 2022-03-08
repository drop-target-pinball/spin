package sdl

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/proc"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type driverState struct {
	on               bool
	lastUpdate       time.Time
	procSchedule     uint32
	procCycleSeconds uint8
	procNow          bool
	pulse            bool
	pulseExpire      time.Time
	//onPWM            int
	//offPWM           int
}

type console struct {
	background *sdl.Texture
	r          *sdl.Renderer
	devices    map[string]spin.LayoutShape
	states     map[string]*driverState
}

type consoleSystem struct {
	consoles map[string]*console
}

func RegisterConsoleSystem(eng *spin.Engine) {
	sys := &consoleSystem{
		consoles: make(map[string]*console),
	}

	if err := img.Init(img.INIT_JPG | img.INIT_PNG); err != nil {
		log.Fatalf("unable to init image system: %v", err)
	}

	eng.RegisterActionHandler(sys)
	//eng.RegisterEventHandler(sys)
	eng.RegisterServer(sys)
}

func (s *consoleSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.AllLampsOff:
		s.allLampsOff()
	case spin.DriverOn:
		s.updateState(act.ID, driverState{on: true})
	case spin.DriverOff:
		s.updateState(act.ID, driverState{})
	case spin.DriverPulse:
		pulse := act.Time
		if pulse <= 0 {
			pulse = 25
		}
		expire := time.Now().Add(time.Duration(pulse) * time.Millisecond)
		s.updateState(act.ID, driverState{on: true, pulse: true, pulseExpire: expire})
	case proc.DriverSchedule:
		s.updateState(act.ID, driverState{on: true, procSchedule: act.Schedule, procNow: act.Now, procCycleSeconds: act.CycleSeconds})
	case spin.RegisterConsole:
		s.registerConsole(act)
	case spin.RegisterFlasher:
		s.registerConsoleDevice(act.ID, act.Layout)
	case spin.RegisterLamp:
		s.registerConsoleDevice(act.ID, act.Layout)
	}
}

func (s *consoleSystem) allLampsOff() {
	cons := s.consoles[""]
	for id := range cons.devices {
		cons.states[id] = &driverState{}
	}
}

func (s *consoleSystem) updateState(id string, state driverState) {
	cons := s.consoles[""]
	if _, ok := cons.states[id]; !ok {
		return
	}
	state.lastUpdate = time.Now()
	cons.states[id] = &state
}

func (s *consoleSystem) registerConsole(act spin.RegisterConsole) {
	file := path.Join(os.Getenv("SPIN_DIR"), act.Image)
	c := &console{
		devices: make(map[string]spin.LayoutShape),
		states:  make(map[string]*driverState),
	}
	image, err := img.Load(file)
	if err != nil {
		log.Panicf("unable to load image %v: %v", file, err)
	}

	mode, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		log.Panic(err)
	}

	win, err := sdl.CreateWindow("Console",
		mode.W-image.W, 0,
		int32(image.W), int32(image.H),
		sdl.WINDOW_HIDDEN)
	if err != nil {
		log.Panicf("unable to create window: %v", err)
	}

	r, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Panicf("unable to create renderer: %v", err)
	}

	texture, err := r.CreateTextureFromSurface(image)
	if err != nil {
		log.Panic(err)
	}
	c.background = texture

	win.Show()
	r.Copy(c.background, nil, nil)
	r.Present()

	c.r = r
	s.consoles[act.ID] = c
}

func (s *consoleSystem) registerConsoleDevice(id string, layout spin.LayoutShape) {
	c, ok := s.consoles[""]
	if !ok {
		spin.Warn("no such console")
	}
	c.devices[id] = layout
	c.states[id] = &driverState{}
}

func (s *consoleSystem) Service(t time.Time) {
	c, ok := s.consoles[""]
	if !ok {
		return
	}
	r := c.r

	r.Copy(c.background, nil, nil)
	r.SetDrawBlendMode(sdl.BLENDMODE_BLEND)

	mt := t.UnixMicro()
	for id, layout := range c.devices {
		state := c.states[id]
		if !state.on {
			continue
		}
		if state.procSchedule != 0 {
			micros := mt % 1e+6      // second remainder
			period := micros / 31250 // every 1/32 of a second
			bit := uint32(1 << period)

			if period == 0 && state.procCycleSeconds > 0 {
				state.procCycleSeconds--
				if state.procCycleSeconds == 0 {
					state.on = false
				}
			}

			if state.procSchedule&bit == 0 {
				continue
			}
		}
		if state.pulse && t.After(state.pulseExpire) {
			state.on = false
		}
		renderShape(r, layout)
	}
	r.Present()
}

func renderShape(r *sdl.Renderer, shape spin.LayoutShape) {
	switch s := shape.(type) {
	case spin.LayoutRect:
		rect := sdl.Rect{X: int32(s.X), Y: int32(s.Y), W: int32(s.W), H: int32(s.H)}
		r.SetDrawColor(s.Color.R, s.Color.G, s.Color.B, s.Color.A)
		r.FillRect(&rect)
	case spin.LayoutCircle:
		gfx.FilledCircleColor(r, int32(s.X), int32(s.Y), int32(s.R), sdl.Color(s.Color))
	case spin.LayoutMulti:
		for _, shape := range s.Shapes {
			renderShape(r, shape)
		}
	}

}
