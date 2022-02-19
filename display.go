package spin

import (
	"image/color"
)

var (
	ColorOff = color.RGBA{0x00, 0x00, 0x00, 0xff}
	ColorOn8 = color.RGBA{0x88, 0x88, 0x88, 0xff}
	ColorOn  = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

var (
	DotMatrixColors = []color.RGBA{
		{0x00, 0x00, 0x00, 0xff},
		{0x11, 0x11, 0x11, 0xff},
		{0x22, 0x22, 0x22, 0xff},
		{0x33, 0x33, 0x33, 0xff},
		{0x44, 0x44, 0x44, 0xff},
		{0x55, 0x55, 0x55, 0xff},
		{0x66, 0x66, 0x66, 0xff},
		{0x77, 0x77, 0x77, 0xff},
		{0x88, 0x88, 0x88, 0xff},
		{0x99, 0x99, 0x99, 0xff},
		{0xaa, 0xaa, 0xaa, 0xff},
		{0xbb, 0xbb, 0xbb, 0xff},
		{0xcc, 0xcc, 0xcc, 0xff},
		{0xdd, 0xdd, 0xdd, 0xff},
		{0xee, 0xee, 0xee, 0xff},
		{0xff, 0xff, 0xff, 0xff},
	}
)

const (
	LayerPriority = "LayerPriority"
)

const (
	PriorityAnnounce = 1
)

type AnchorX int

const (
	AnchorLeft AnchorX = iota
	AnchorCenter
	AnchorRight
)

type AnchorY int

const (
	AnchorTop AnchorY = iota
	AnchorMiddle
	AnchorBottom
)

type Graphics struct {
	X       int32
	Y       int32
	W       int32
	H       int32
	Color   color.RGBA
	Font    string
	AnchorX AnchorX
	AnchorY AnchorY
}

type Renderer interface {
	Graphics() *Graphics // TODO: Make this not a pointer
	Fill(color.RGBA)
	FillRect(*Graphics)
	Height() int32
	Print(*Graphics, string, ...interface{})
	Width() int32
	Close()
}

type Display interface {
	Open(int) Renderer
	Width() int
	Height() int
	At(int, int) color.Color
}

type DisplayOptions struct {
	ID     string
	Width  int
	Height int
	Layers []string
}

// https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
func RGBToGray(rgb color.Color) uint8 {
	r, g, b, _ := rgb.RGBA()
	lum := 0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b)
	return uint8(lum / 256)
}

func RGBAToUint32(c color.RGBA) uint32 {
	return uint32(c.R)<<24 | uint32(c.G)<<16 | uint32(c.B)<<8 | uint32(c.A)
}
