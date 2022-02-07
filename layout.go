package spin

import "image/color"

var (
	ColorLampBlue   = color.RGBA{R: 0x00, G: 0x40, B: 0xff, A: 0xc0}
	ColorLampGreen  = color.RGBA{R: 0x00, G: 0xa0, B: 0x00, A: 0xc0}
	ColorLampOrange = color.RGBA{R: 0xff, G: 0x80, B: 0x00, A: 0xc0}
	ColorLampRed    = color.RGBA{R: 0xff, G: 0x00, B: 0x00, A: 0xc0}
	ColorLampWhite  = color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xc0}
	ColorLampYellow = color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0xc0}
)

type LayoutShape interface {
	shape()
}

type LayoutRect struct {
	LayoutShape
	X     int
	Y     int
	W     int
	H     int
	Color color.RGBA
}

func NewLayoutRect(x int, y int, w int, h int, color color.RGBA) LayoutRect {
	return LayoutRect{X: x, Y: y, W: w, H: h, Color: color}
}

type LayoutCircle struct {
	LayoutShape
	X     int
	Y     int
	R     int
	Color color.RGBA
}

func NewLayoutCircle(x int, y int, r int, color color.RGBA) LayoutCircle {
	return LayoutCircle{X: x, Y: y, R: r, Color: color}
}

func NewLayoutCircleFromRect(x int, y int, w int, h int, color color.RGBA) LayoutCircle {
	cx := x + (w / 2)
	cy := y + (h / 2)
	r := w
	if h < w {
		r = h
	}
	return LayoutCircle{X: cx, Y: cy, R: r, Color: color}
}

type LayoutMulti struct {
	LayoutShape
	Shapes []LayoutShape
}

func NewLayoutMulti(shapes ...LayoutShape) LayoutMulti {
	return LayoutMulti{Shapes: shapes}
}
