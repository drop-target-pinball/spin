package terminal

const (
	// carriage return to go to the beginning of the line
	// then ansi escape sequence to clear the line
	AnsiClearLine    = "\r\033[2K"
	AnsiReset        = "\033[0m"
	AnsiLightBlue    = "\033[1;34m"
	AnsiLightGreen   = "\033[1;32m"
	AnsiPreviousLine = "\033[F"
	AnsiCyan         = "\033[0;36m"
)
