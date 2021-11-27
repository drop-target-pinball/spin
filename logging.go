package spin

import "log"

type LoggingSystem struct {
}

func RegisterLoggingSystem(eng *Engine) {
	sys := &LoggingSystem{}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
}

func (s *LoggingSystem) HandleAction(act Action) {
	log.Println(String(act))
}

func (s *LoggingSystem) HandleEvent(evt Event) {
	log.Println(String(evt))
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
