package system

import (
	"log"
	"sync"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
)

type displaySDL struct {
	id    string
	surf  *sdl.Surface
	mutex sync.Mutex
}

func (d *displaySDL) Width() int {
	return int(d.surf.W)
}

func (d *displaySDL) Height() int {
	return int(d.surf.H)
}

func (d *displaySDL) Graphics() spin.Graphics {
	return &graphicsSDL{
		surf:  d.surf,
		mutex: &d.mutex,
		color: 0xffffffff,
	}
}

func NewDisplaySDL(eng *spin.Engine, opts spin.DisplayOptions) {
	surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(opts.Width), int32(opts.Height),
		32, sdl.PIXELFORMAT_RGB888)
	if err != nil {
		log.Fatalf("unable to create SDL surface: %v", err)
	}

	s := &displaySDL{
		id:   opts.ID,
		surf: surf,
	}
	eng.Do(spin.RegisterDisplaySDL{
		ID:      s.id,
		Display: s,
		Surface: s.surf,
		Mutex:   &s.mutex,
	})
}

type graphicsSDL struct {
	surf  *sdl.Surface
	mutex *sync.Mutex
	x     int
	y     int
	color uint32
}

func (g *graphicsSDL) Move(x int, y int) {
	g.x = x
	g.y = y
}

func (g *graphicsSDL) Color(c uint32) {
	g.color = c
}

func (g *graphicsSDL) FillRect(w int, h int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	rect := sdl.Rect{X: int32(g.x), Y: int32(g.y), W: int32(w), H: int32(h)}
	if err := g.surf.FillRect(&rect, g.color); err != nil {
		log.Fatal(err)
	}
}
