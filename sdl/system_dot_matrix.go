package sdl

import (
	"image/color"
	"log"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type OptionsDotMatrix struct {
	ID          string
	Scale       int
	Padding     int
	BackColor   color.RGBA
	BorderColor color.RGBA
	BorderSize  int
	Title       string
	Palette     []color.RGBA
}

func DefaultOptionsDotMatrix() OptionsDotMatrix {
	return OptionsDotMatrix{
		Scale:       4,
		Padding:     1,
		BackColor:   color.RGBA{0x40, 0x40, 0x40, 0xff},
		BorderColor: color.RGBA{0x80, 0x80, 0x80, 0xff},
		BorderSize:  20,
		Title:       "Dot Matrix Display",
		Palette:     PaletteOrange,
	}
}

type dotMatrixSystem struct {
	eng     *spin.Engine
	opts    OptionsDotMatrix
	winW    int
	winH    int
	source  spin.Display
	target  *sdl.Renderer
	win     *sdl.Window
	borders [4]sdl.Rect
}

func RegisterDotMatrixSystem(eng *spin.Engine, opts OptionsDotMatrix) {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("unable to initialize SDL: %v", err)
	}
	s := &dotMatrixSystem{eng: eng, opts: opts}
	eng.RegisterActionHandler(s)
}

func (s *dotMatrixSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterDisplay:
		s.registerDisplay(act)
	}
}

func (s *dotMatrixSystem) registerDisplay(act spin.RegisterDisplay) {
	if act.ID != s.opts.ID {
		return
	}
	s.source = act.Display

	sourceW, sourceH := s.source.Width(), s.source.Height()
	s.winW = ((sourceW * s.opts.Scale) + (s.opts.Padding * sourceW) +
		s.opts.Padding + (s.opts.BorderSize * 2))
	s.winH = ((sourceH * s.opts.Scale) + (s.opts.Padding * sourceH) +
		s.opts.Padding + (s.opts.BorderSize * 2))

	win, err := sdl.CreateWindow(s.opts.Title,
		//sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		0, 0,
		int32(s.winW), int32(s.winH),
		sdl.WINDOW_HIDDEN)
	if err != nil {
		log.Fatalf("unable to create window: %v", err)
	}

	r, err := sdl.CreateRenderer(win, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)
	if err != nil {
		log.Fatalf("unable to create renderer: %v", err)
	}

	// Delay the showing of the window until after the renderer is created. If
	// not, the window will be created twice with the first window opening
	// and closing really quickly.
	win.Show()

	info, err := r.GetInfo()
	if err != nil {
		log.Fatalf("unable to get renderer info: %v", err)
	}
	// FIXME: Change this to check for OpenGL
	// Or not? On macOS, this reports as "Metal". Check Linux.
	if info.Name != "direct3d" {
		if _, err := win.GLCreateContext(); err != nil {
			log.Printf("unable to create GL context: %v", err)
		}
		if err = sdl.GLSetSwapInterval(1); err != nil {
			log.Printf("unable to set swap interval: %v", err)
		}
	}

	// top
	s.borders[0] = sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(s.winW),
		H: int32(s.opts.BorderSize),
	}
	// bottom
	s.borders[1] = sdl.Rect{
		X: 0,
		Y: int32(s.winH - s.opts.BorderSize),
		W: int32(s.winW),
		H: int32(s.opts.BorderSize),
	}
	// left
	s.borders[2] = sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(s.opts.BorderSize),
		H: int32(s.winH),
	}
	// right
	s.borders[3] = sdl.Rect{
		X: int32(s.winW - s.opts.BorderSize),
		Y: 0,
		W: int32(s.opts.BorderSize),
		H: int32(s.winH),
	}

	s.win = win
	s.target = r

	s.eng.RegisterServer(s)
	sdl.PumpEvents()
}

func (s *dotMatrixSystem) Service(_ time.Time) {
	if s.target == nil {
		return
	}

	o := s.opts
	r := s.target

	// Background
	bg := o.BackColor
	r.SetDrawColor(bg.R, bg.G, bg.B, 0xff)
	r.Clear()

	// Border
	bdr := o.BorderColor
	r.SetDrawColor(bdr.R, bdr.G, bdr.B, 0xff)
	for _, b := range s.borders {
		r.FillRect(&b)
	}

	// Dots
	for x := 0; x < s.source.Width(); x++ {
		for y := 0; y < int(s.source.Height()); y++ {
			dx := o.BorderSize + (x * o.Padding) + o.Padding + (x * o.Scale)
			dy := o.BorderSize + (y * o.Padding) + o.Padding + (y * o.Scale)
			y := spin.RGBToGray(s.source.At(x, y))
			y = y >> 4
			c := o.Palette[y]
			r.SetDrawColor(c.R, c.G, c.B, 0xff)
			dest := sdl.Rect{
				X: int32(dx),
				Y: int32(dy),
				W: int32(o.Scale),
				H: int32(o.Scale),
			}
			r.FillRect(&dest)
		}
	}

	r.Present()
}

var PaletteOrange = []color.RGBA{
	{0, 0, 0, 0xff},
	{15, 8, 0, 0xff},
	{33, 17, 0, 0xff},
	{51, 25, 0, 0xff},
	{66, 33, 0, 0xff},
	{84, 42, 0, 0xff},
	{102, 51, 0, 0xff},
	{117, 58, 0, 0xff},
	{135, 67, 0, 0xff},
	{153, 76, 0, 0xff},
	{168, 84, 0, 0xff},
	{186, 93, 0, 0xff},
	{204, 102, 0, 0xff},
	{219, 109, 0, 0xff},
	{237, 118, 0, 0xff},
	{255, 127, 0, 0xff},
}
