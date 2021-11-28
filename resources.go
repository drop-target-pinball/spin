package spin

import "log"

type Resources struct {
	Fonts map[string]struct{}
}

type resourceSystem struct {
	resources Resources
}

func registerResourceSystem(eng *Engine) {
	s := &resourceSystem{
		resources: Resources{
			Fonts: make(map[string]struct{}),
		},
	}
	eng.RegisterVars("resources", &s.resources)
	eng.RegisterActionHandler(s)
}

func (s *resourceSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case RegisterFont:
		s.resources.Fonts[act.ID] = struct{}{}
	}
}

func ResourceVars(store Store) *Resources {
	v, ok := store.Vars("resources")
	if !ok {
		log.Panicf("resources vars not defined")
	}
	return v.(*Resources)
}
