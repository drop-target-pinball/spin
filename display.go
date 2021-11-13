package spin

type Graphics interface {
	Move(x int, y int)
	FillRect(w int, h int)
}

type Display interface {
	Graphics() Graphics
	Width() int
	Height() int
}

type DisplayOptions struct {
	ID     string
	Width  int
	Height int
}
