package cli

const (
	// carriage return to go to the beginning of the line
	// then ansi escape sequence to clear the line
	AnsiClearScreen  = "\033[2J"
	AnsiBrightBlack  = "\033[0;90m"
	AnsiBrightRed    = "\033[0;91m"
	AnsiClearLine    = "\r\033[2K"
	AnsiCyan         = "\033[0;36m"
	AnsiLightBlue    = "\033[1;34m"
	AnsiLightGreen   = "\033[1;32m"
	AnsiMoveToBottom = "\033[200;0H" // go to line 200, column 0
	AnsiPreviousLine = "\033[F"
	AnsiReset        = "\033[0m"
)
