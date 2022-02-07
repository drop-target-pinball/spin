package sdl

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type console struct {
	background *sdl.Texture
	r          *sdl.Renderer
	devices    map[string]spin.LayoutShape
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
	eng.RegisterEventHandler(sys)
	eng.RegisterServer(sys)
}

func (s *consoleSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterConsole:
		s.registerConsole(act)
	case spin.RegisterFlasher:
		s.registerConsoleDevice(act.ID, act.Layout)
	case spin.RegisterLamp:
		s.registerConsoleDevice(act.ID, act.Layout)
	}
}

func (s *consoleSystem) HandleEvent(event spin.Event) {
}

func (s *consoleSystem) registerConsole(act spin.RegisterConsole) {
	file := path.Join(os.Getenv("SPIN_DIR"), act.Image)
	c := &console{
		devices: make(map[string]spin.LayoutShape),
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
}

func (s *consoleSystem) Service(t time.Time) {
	c, ok := s.consoles[""]
	if !ok {
		return
	}
	r := c.r

	t0 := t.UnixMilli()

	r.Copy(c.background, nil, nil)

	r.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if t0%1000 < 500 {
		for _, shape := range c.devices {
			renderShape(r, shape)
		}
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
