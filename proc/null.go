package proc

import (
	"github.com/drop-target-pinball/spin"
)

type nullSystem struct {
	eng *spin.Engine
}

func RegisterNullSystem(eng *spin.Engine) {
	s := &nullSystem{eng: eng}
	eng.RegisterActionHandler(s)
}

func (s *nullSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.RegisterSwitch:
		s.registerSwitch(act)
	}
}

func (s *nullSystem) registerSwitch(act spin.RegisterSwitch) {
	rv := spin.GetResourceVars(s.eng)
	sw := spin.Switch{
		ID: act.ID,
		NC: act.NC,
	}
	rv.Switches[act.ID] = &sw
}
