package console

import (
	"fmt"

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
	fmt.Println(act)
}

func (s *LoggingSystem) HandleEvent(evt spin.Event) {
	fmt.Println(evt)
}
