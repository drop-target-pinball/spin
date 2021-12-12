package spin

type Resources struct {
	Fonts    map[string]struct{}
	Switches map[string]*Switch
}

type resourceSystem struct {
	resources *Resources
}

func registerResourceSystem(eng *Engine) {
	s := &resourceSystem{
		resources: ResourceVars(eng),
	}
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
	var vars *Resources
	if ok {
		vars = v.(*Resources)
	} else {
		vars = &Resources{
			Fonts:    make(map[string]struct{}),
			Switches: make(map[string]*Switch),
		}
		store.RegisterVars("resources", vars)
	}
	return vars
}
