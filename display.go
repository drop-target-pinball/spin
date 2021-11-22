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
	Clear()
	FillRect(*Graphics)
	Height() int32
	//Lock()
	Print(*Graphics, string, ...interface{})
	Println(*Graphics, string, ...interface{})
	//Unlock()
	Width() int32
}

type Display interface {
	Renderer() (Renderer, *Graphics)
	Width() int
	Height() int
}

type DisplayOptions struct {
	ID     string
	Width  int
	Height int
}
