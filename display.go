package spin

import (
	"image/color"
)

var (
	ColorBlack = color.RGBA{0x00, 0x00, 0x00, 0xff}
	ColorWhite = color.RGBA{0xff, 0xff, 0xff, 0xff}
)

const (
	LayerPriority = "LayerPriority"
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
	Color   uint32
	Font    string
	AnchorX AnchorX
	AnchorY AnchorY
}

type Renderer interface {
	Graphics() *Graphics
	Fill(color.RGBA)
	FillRect(*Graphics)
	Height() int32
	Print(*Graphics, string, ...interface{})
	Width() int32
	Clear()
}

type Display interface {
	Clear(string)
	Renderer(string) (Renderer, *Graphics)
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
