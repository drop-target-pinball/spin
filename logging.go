package spin

import "log"

type loggingSystem struct {
}

func RegisterLoggingSystem(eng *Engine) {
	sys := &loggingSystem{}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
}

func (s *loggingSystem) HandleAction(act Action) {
	log.Println(FormatAction(act))
}

func (s *loggingSystem) HandleEvent(evt Event) {
	log.Println(FormatEvent(evt))
}

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
