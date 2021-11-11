package system

import (
	"log"

	"github.com/drop-target-pinball/spin"
)

type LoggingConsole struct {
}

func NewLoggingConsole(eng *spin.Engine) *LoggingConsole {
	sys := &LoggingConsole{}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	return sys
}

func (s *LoggingConsole) HandleAction(act spin.Action) {
	log.Println(spin.String(act))
}

func (s *LoggingConsole) HandleEvent(evt spin.Event) {
	log.Println(spin.String(evt))
}
