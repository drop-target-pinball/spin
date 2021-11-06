package spin

import "log"

func Warn(format string, a ...interface{}) {
	log.Print("(!) ")
	log.Printf(format, a...)
}

func Log(format string, a ...interface{}) {
	log.Printf(format, a...)
}
