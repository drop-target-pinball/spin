package system

import (
	"image/color"
	"log"
	"sync"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type OptionsDotMatrixSDL struct {
	Name        string
	Width       int
	Height      int
	Scale       int
	Padding     int
	BackColor   color.RGBA
	BorderColor color.RGBA
	BorderSize  int
	Title       string
	Palette     []color.RGBA
}

func DefaultOptionsDotMatrixSDL() OptionsDotMatrixSDL {
	return OptionsDotMatrixSDL{
		Width:       128,
		Height:      32,
		Scale:       4,
		Padding:     1,
		BackColor:   color.RGBA{0x40, 0x40, 0x40, 0xff},
		BorderColor: color.RGBA{0x80, 0x80, 0x80, 0xff},
		BorderSize:  20,
		Title:       "Dot Matrix Display",
		Palette:     PaletteOrange,
	}
}

type DotMatrixSDL struct {
	width    int
	height   int
	renderer *sdl.Renderer
	surf     *sdl.Surface
	mutex    *sync.Mutex
	opts     OptionsDotMatrixSDL
	win      *sdl.Window
	borders  [4]sdl.Rect
}

func NewDotMatrixSDL(eng *spin.Engine, o OptionsDotMatrixSDL) *DotMatrixSDL {
	if err := sdl.Init(sdl.INIT_VIDEO); err != nil {
		log.Fatalf("unable to initialize SDL: %v", err)
	}

	rt, ok := eng.RenderTargetSDL[o.Name]
	if !ok {
		log.Fatalf("no such SDL renderer: %v", o.Name)
	}

	d := &DotMatrixSDL{surf: rt.Surface, mutex: &rt.Mutex, opts: o}
	surfW, surfH := int(rt.Surface.W), int(rt.Surface.H)
	d.width = ((surfW * o.Scale) + (o.Padding * surfW) + o.Padding +
		(o.BorderSize * 2))
	d.height = ((surfH * o.Scale) + (o.Padding * surfH) + o.Padding +
		(o.BorderSize * 2))

	win, err := sdl.CreateWindow(o.Title,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(d.width), int32(d.height),
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
	d.borders[0] = sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(d.width),
		H: int32(o.BorderSize),
	}
	// bottom
	d.borders[1] = sdl.Rect{
		X: 0,
		Y: int32(d.height - o.BorderSize),
		W: int32(d.width),
		H: int32(o.BorderSize),
	}
	// left
	d.borders[2] = sdl.Rect{
		X: 0,
		Y: 0,
		W: int32(o.BorderSize),
		H: int32(d.height),
	}
	// right
	d.borders[3] = sdl.Rect{
		X: int32(d.width - o.BorderSize),
		Y: 0,
		W: int32(o.BorderSize),
		H: int32(d.height),
	}

	d.win = win
	d.renderer = r

	eng.RegisterServer(d)
	sdl.PumpEvents()
	return d
}

func (s *DotMatrixSDL) Service() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	o := s.opts
	r := s.renderer

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
	for x := 0; x < int(s.surf.W); x++ {
		for y := 0; y < int(s.surf.H); y++ {
			dx := o.BorderSize + (x * o.Padding) + o.Padding + (x * o.Scale)
			dy := o.BorderSize + (y * o.Padding) + o.Padding + (y * o.Scale)
			y := rgbToGray(s.surf.At(x, y))
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

// https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
func rgbToGray(rgb color.Color) uint8 {
	r, g, b, _ := rgb.RGBA()
	lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return uint8(lum / 256)
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
