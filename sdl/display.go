package sdl

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/drop-target-pinball/spin"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type displaySystem struct {
	id    string
	surf  *sdl.Surface
	fonts map[string]font
}

func (d *displaySystem) Width() int {
	return int(d.surf.W)
}

func (d *displaySystem) Height() int {
	return int(d.surf.H)
}

func (d *displaySystem) Renderer() (spin.Renderer, *spin.Graphics) {
	return &rendererSDL{
		surf:  d.surf,
		fonts: d.fonts,
	}, &spin.Graphics{}
}

func RegisterDisplaySystem(eng *spin.Engine, opts spin.DisplayOptions) {
	if err := ttf.Init(); err != nil {
		log.Fatalf("unable to initialize ttf: %v", err)
	}
	surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(opts.Width), int32(opts.Height),
		32, sdl.PIXELFORMAT_RGB888)
	if err != nil {
		log.Fatalf("unable to create SDL surface: %v", err)
	}

	s := &displaySystem{
		id:    opts.ID,
		surf:  surf,
		fonts: make(map[string]font),
	}
	eng.Do(spin.RegisterDisplaySDL{
		ID:      s.id,
		Display: s,
		Surface: s.surf,
		//Mutex:   &s.mutex,
	})
	eng.RegisterActionHandler(s)
}

func (s *displaySystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterFont:
		s.registerFont(act)
	}
}

// ----------------------------------------------------------------------------

type rendererSDL struct {
	surf  *sdl.Surface
	fonts map[string]font
}

func (r *rendererSDL) Clear() {
	rect := sdl.Rect{X: 0, Y: 0, W: r.surf.W, H: r.surf.H}
	if err := r.surf.FillRect(&rect, 0); err != nil {
		log.Panic(err)
	}
}

func (r *rendererSDL) Width() int32 {
	return r.surf.W
}

func (r *rendererSDL) Height() int32 {
	return r.surf.H
}

func (r *rendererSDL) FillRect(g *spin.Graphics) {
	rect := sdl.Rect{X: g.X, Y: g.Y, W: g.W, H: g.H}
	if err := r.surf.FillRect(&rect, g.Color); err != nil {
		log.Panic(err)
	}
}

func (r *rendererSDL) Print(g *spin.Graphics, format string, a ...interface{}) {
	font := r.getFont(g)
	if font == nil {
		return
	}

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

func (s *displaySystem) registerFont(act spin.RegisterFont) {
	if _, exists := s.fonts[act.ID]; exists {
		return
	}
	switch {
	case strings.HasSuffix(act.Path, ".ttf"):
		s.registerFontTTF(act)
	case strings.HasSuffix(act.Path, ".dmd"):
		s.registerFontBitmap(act)
	default:
		spin.Warn("unknown font format: %v", act.Path)
	}
}

func (s *displaySystem) registerFontTTF(act spin.RegisterFont) {
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

func (s *displaySystem) registerFontBitmap(act spin.RegisterFont) {
	fontPath := path.Join(spin.AssetDir, act.Path)
	dots, err := ioutil.ReadFile(fontPath)
	if err != nil {
		spin.Warn("unable to read file: %v", err)
		return
	}
	surf, err := DecodeDMD1(dots)
	if err != nil {
		spin.Warn("unable to decode bitmap file: %v", err)
		return
	}
	tmFile := regexpExt.ReplaceAllString(fontPath, ".json")
	tileMap, err := loadTileMap(tmFile)
	if err != nil {
		spin.Warn("unable to read tile map: %v", err)
		return
	}
	s.fonts[act.ID] = &fontBitmap{surf: surf, tileMap: tileMap, tracking: 1}
}

type tile struct {
	X int32
	Y int32
	W int32
	H int32
}

type tileMap map[string]tile

func loadTileMap(name string) (tileMap, error) {
	tmText, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("unable to load tile map: %v", err)
	}
	tilemap := make(tileMap)
	if err := json.Unmarshal(tmText, &tilemap); err != nil {
		return nil, fmt.Errorf("unable to parse tile map: %v", err)
	}
	return tilemap, nil
}

type fontBitmap struct {
	surf     *sdl.Surface
	tileMap  tileMap
	tracking int32
}

func (f *fontBitmap) render(target *sdl.Surface, x int32, y int32, text string) {
	for _, ch := range text {
		t, ok := f.tileMap[string(ch)]
		if !ok {
			continue
		}
		srcRect := sdl.Rect{X: t.X, Y: t.Y, W: t.W, H: t.H}
		tgtRect := sdl.Rect{X: x, Y: y, W: t.W, H: t.H}
		if err := f.surf.Blit(&srcRect, target, &tgtRect); err != nil {
			log.Panic(err)
		}
		x += t.W + f.tracking
	}
}

func (f *fontBitmap) size(text string) (int32, int32) {
	var w, h int32
	for _, ch := range text {
		tile, ok := f.tileMap[string(ch)]
		if !ok {
			continue
		}
		w += tile.W + f.tracking
		if tile.H > h {
			h = tile.H
		}
	}
	w -= f.tracking
	return w, h
}

// ----------------------------------------------------------------------------

func DecodeDMD(data []byte) ([]*sdl.Surface, error) {
	in := bytes.NewReader(data)
	var header, nFrames, width, height uint32
	if err := binary.Read(in, binary.LittleEndian, &header); err != nil {
		return nil, err
	}
	if err := binary.Read(in, binary.LittleEndian, &nFrames); err != nil {
		return nil, err
	}
	if err := binary.Read(in, binary.LittleEndian, &width); err != nil {
		return nil, err
	}
	if err := binary.Read(in, binary.LittleEndian, &height); err != nil {
		return nil, err
	}

	frames := make([]*sdl.Surface, nFrames)
	dots := make([]byte, width*height)
	for i := uint32(0); i < nFrames; i++ {
		surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(width), int32(height),
			32, sdl.PIXELFORMAT_RGB888)
		if err != nil {
			return nil, fmt.Errorf("unable to create RGB surface: %v", err)
		}

		if err := binary.Read(in, binary.LittleEndian, &dots); err != nil {
			return nil, err
		}
		x, y := 0, 0
		for _, dot := range dots {
			// Values in file are going to be between 0x0 and 0xf. Copy the
			// lower nibble to the higher nibble.
			dot = dot<<4 + dot
			surf.Set(x, y, color.RGBA{R: dot, G: dot, B: dot, A: 0xff})
			x++
			if x >= int(width) {
				x = 0
				y++
			}
		}
		frames[i] = surf
	}
	return frames, nil
}

func DecodeDMD1(data []byte) (*sdl.Surface, error) {
	frames, err := DecodeDMD(data)
	if err != nil {
		return nil, err
	}
	return frames[0], nil
}
