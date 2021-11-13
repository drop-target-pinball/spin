package system

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"sync"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type displaySDL struct {
	id    string
	surf  *sdl.Surface
	mutex sync.Mutex
	fonts map[string]font
}

func (d *displaySDL) Width() int {
	return int(d.surf.W)
}

func (d *displaySDL) Height() int {
	return int(d.surf.H)
}

func (d *displaySDL) Renderer() spin.Renderer {
	return &rendererSDL{
		surf:  d.surf,
		mutex: &d.mutex,
		fonts: d.fonts,
	}
}

func NewDisplaySDL(eng *spin.Engine, opts spin.DisplayOptions) {
	if err := ttf.Init(); err != nil {
		log.Fatalf("unable to initialize ttf: %v", err)
	}
	surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(opts.Width), int32(opts.Height),
		32, sdl.PIXELFORMAT_RGB888)
	if err != nil {
		log.Fatalf("unable to create SDL surface: %v", err)
	}

	s := &displaySDL{
		id:    opts.ID,
		surf:  surf,
		fonts: make(map[string]font),
	}
	eng.Do(spin.RegisterDisplaySDL{
		ID:      s.id,
		Display: s,
		Surface: s.surf,
		Mutex:   &s.mutex,
	})
	eng.RegisterActionHandler(s)
}

func (s *displaySDL) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterFont:
		s.registerFont(act)
	}
}

// ----------------------------------------------------------------------------

type rendererSDL struct {
	surf  *sdl.Surface
	mutex *sync.Mutex
	fonts map[string]font
}

func (r *rendererSDL) Width() int32 {
	return r.surf.W
}

func (r *rendererSDL) Height() int32 {
	return r.surf.H
}

func (r *rendererSDL) FillRect(g *spin.Graphics) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	rect := sdl.Rect{X: g.X, Y: g.Y, W: g.W, H: g.H}
	if err := r.surf.FillRect(&rect, g.Color); err != nil {
		log.Fatal(err)
	}
}

func (r *rendererSDL) Print(g *spin.Graphics, format string, a ...interface{}) {
	font := r.getFont(g)
	if font == nil {
		return
	}
	r.mutex.Lock()
	defer r.mutex.Unlock()

	text := fmt.Sprintf(format, a...)
	if g.W == 0 {
		font.render(r.surf, g.X, g.Y, text)
		return
	}
	w, h := font.size(text)
	x, y := g.X, g.Y
	if g.W > 0 {
		x += (r.surf.W - w) / 2
	}
	if g.H > 0 {
		y += (r.surf.H - h) / 2
	}
	font.render(r.surf, x, y, text)
}

func (r *rendererSDL) Println(g *spin.Graphics, text string, a ...interface{}) {
	font := r.getFont(g)
	if font == nil {
		return
	}
	r.Print(g, text, a...)
	_, h := font.size(text)
	g.X = 0
	g.Y = g.Y + h + g.PaddingV
}

func (r *rendererSDL) getFont(g *spin.Graphics) font {
	if g.Font == "" {
		spin.Warn("no font selected")
		return nil
	}
	font, ok := r.fonts[g.Font]
	if !ok {
		spin.Warn("no such font: %v", g.Font)
		return nil
	}
	return font
}

// ----------------------------------------------------------------------------

type font interface {
	render(s *sdl.Surface, x int32, y int32, text string)
	size(string) (int32, int32)
}

type infoTTF struct {
	OffsetY int32
}

type fontTTF struct {
	font *ttf.Font
	info infoTTF
}

func (f *fontTTF) render(target *sdl.Surface, x int32, y int32, text string) {
	surf, err := f.font.RenderUTF8Solid(text,
		sdl.Color{R: 0xff, G: 0xff, B: 0xff, A: 0xff})
	if err != nil {
		panic(err)
	}
	srcRect := sdl.Rect{X: 0, Y: 0, W: surf.W, H: surf.H}
	tgtRect := sdl.Rect{X: x, Y: y + f.info.OffsetY, W: surf.W, H: surf.H}
	surf.Blit(&srcRect, target, &tgtRect)
}

func (f *fontTTF) size(text string) (int32, int32) {
	w, h, err := f.font.SizeUTF8(text)
	if err != nil {
		panic(err)
	}
	return int32(w), int32(h) + f.info.OffsetY
}

var regexpExt = regexp.MustCompile(`\.[^\.]+$`)

func (s *displaySDL) registerFont(act spin.RegisterFont) {
	s.registerFontTTF(act)
}

func (s *displaySDL) registerFontTTF(act spin.RegisterFont) {
	if _, exists := s.fonts[act.ID]; exists {
		return
	}
	fontPath := path.Join(spin.AssetDir, act.Path)
	font, err := ttf.OpenFont(fontPath, act.Size)
	if err != nil {
		spin.Warn("unable to load font: %v", err)
		return
	}

	var info infoTTF
	infoFile := regexpExt.ReplaceAllString(fontPath, ".json")
	_, err = os.Stat(infoFile)
	if err != nil && !os.IsNotExist(err) {
		spin.Warn("unable to load font descriptor: %v", err)
		return
	}
	if !os.IsNotExist(err) {
		infoText, err := ioutil.ReadFile(infoFile)
		if err != nil {
			spin.Warn("unable to load font descriptor: %v", err)
			return
		}
		if err := json.Unmarshal(infoText, &info); err != nil {
			spin.Warn("unable to parse font descriptor: %v", err)
			return
		}
	}
	s.fonts[act.ID] = &fontTTF{font, info}
}
