package spin

import (
	"log"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

type RenderTargetSDL struct {
	Surface *sdl.Surface
	Mutex   sync.Mutex
}

func newRendererSDL(width int, height int) *RenderTargetSDL {
	surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(width), int32(height),
		32, sdl.PIXELFORMAT_RGB888)
	if err != nil {
		log.Fatalf("unable to create SDL surface: %v", err)
	}
	return &RenderTargetSDL{Surface: surf}
}

func (r *RenderTargetSDL) Graphics() *GraphicsSDL {
	return NewGraphicsSDL(r)
}

type GraphicsSDL struct {
	X     int
	Y     int
	Color uint32
	surf  *sdl.Surface
	mutex *sync.Mutex
}

func NewGraphicsSDL(r *RenderTargetSDL) *GraphicsSDL {
	return &GraphicsSDL{
		// Color: 0x0f,
		Color: 0xffffffff,
		surf:  r.Surface,
		mutex: &r.Mutex,
	}
}

func (g *GraphicsSDL) FillRect(w int, h int) {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	rect := sdl.Rect{X: int32(g.X), Y: int32(g.Y), W: int32(w), H: int32(h)}
	if err := g.surf.FillRect(&rect, g.Color); err != nil {
		log.Fatal(err)
	}
}
