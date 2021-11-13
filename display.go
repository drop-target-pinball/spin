package spin

type Graphics struct {
	X        int32
	Y        int32
	W        int32
	H        int32
	Color    uint32
	Font     string
	PaddingV int32
}

type Renderer interface {
	FillRect(*Graphics)
	Height() int32
	Print(*Graphics, string, ...interface{})
	Println(*Graphics, string, ...interface{})
	Width() int32
}

type Display interface {
	Renderer() Renderer
	Width() int
	Height() int
}

type DisplayOptions struct {
	ID     string
	Width  int
	Height int
}
