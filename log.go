package spin

import "log"

func Error(format string, a ...interface{}) {
	format = "(*) " + format
	log.Printf(format, a...)
}

func Warn(format string, a ...interface{}) {
	format = "(!) " + format
	log.Printf(format, a...)
}

func Log(format string, a ...interface{}) {
	log.Printf(format, a...)
}
