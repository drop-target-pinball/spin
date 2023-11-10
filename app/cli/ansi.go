package cli

const (
	// carriage return to go to the beginning of the line
	// then ansi escape sequence to clear the line
	AnsiBrightBlack  = "\033[0;90m"
	AnsiClearLine    = "\r\033[2K"
	AnsiCyan         = "\033[0;36m"
	AnsiLightBlue    = "\033[1;34m"
	AnsiLightGreen   = "\033[1;32m"
	AnsiPreviousLine = "\033[F"
	AnsiReset        = "\033[0m"
)
