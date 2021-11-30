package spin

import "image/color"

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
	X        int32
	Y        int32
	W        int32
	H        int32
	Color    uint32
	Font     string
	PaddingV int32
	AnchorX  AnchorX
	AnchorY  AnchorY
}

type Renderer interface {
	Clear()
	FillRect(*Graphics)
	Height() int32
	Print(*Graphics, string, ...interface{})
	Println(*Graphics, string, ...interface{})
	Width() int32
}

type Display interface {
	Renderer() (Renderer, *Graphics)
	Width() int
	Height() int
	At(int, int) color.Color
}

type DisplayOptions struct {
	ID     string
	Width  int
	Height int
}
