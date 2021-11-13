package system

import (
	"log"

	"github.com/drop-target-pinball/spin"
)

type LoggingConsole struct {
}

func RegisterLoggingConsole(eng *spin.Engine) {
	sys := &LoggingConsole{}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
}

func (s *LoggingConsole) HandleAction(act spin.Action) {
	log.Println(spin.String(act))
}

func (s *LoggingConsole) HandleEvent(evt spin.Event) {
	log.Println(spin.String(evt))
}
