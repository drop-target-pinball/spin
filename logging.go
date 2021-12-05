package spin

import (
	"fmt"
	"log"
	"time"
)

type loggingSystem struct {
	startTime time.Time
}

func RegisterLoggingSystem(eng *Engine) {
	sys := &loggingSystem{
		startTime: time.Now(),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
}

func (s *loggingSystem) HandleAction(act Action) {
	log.Printf("[%v] %v", s.elapsedTime(), FormatAction(act))
}

func (s *loggingSystem) HandleEvent(evt Event) {
	log.Printf("[%v] %v", s.elapsedTime(), FormatEvent(evt))
}

func (s *loggingSystem) elapsedTime() string {
	now := time.Now()
	diff := float64(now.Sub(s.startTime).Milliseconds()) / 1000
	return fmt.Sprintf("%10.3f", diff)
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
