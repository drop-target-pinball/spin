package ansi

const (
	// carriage return to go to the beginning of the line
	// then ansi escape sequence to clear the line
	BrightBlack  = "\033[0;90m"
	ClearLine    = "\r\033[2K"
	Cyan         = "\033[0;36m"
	LightBlue    = "\033[1;34m"
	LightGreen   = "\033[1;32m"
	PreviousLine = "\033[F"
	Reset        = "\033[0m"
)
