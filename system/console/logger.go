package console

import (
	"log"

	"github.com/drop-target-pinball/spin"
)

type LoggingSystem struct {
}

func NewLoggingSystem(eng *spin.Engine) *LoggingSystem {
	sys := &LoggingSystem{}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
	return sys
}

func (s *LoggingSystem) HandleAction(act spin.Action) {
	log.Println(spin.String(act))
}

func (s *LoggingSystem) HandleEvent(evt spin.Event) {
	log.Println(spin.String(evt))
}
