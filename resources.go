package spin

import (
	"reflect"
	"strconv"
)

type Resources struct {
	Fonts    map[string]struct{}
	Switches map[string]*Switch
}

type resourceSystem struct {
	eng       *Engine
	resources *Resources
}

func registerResourceSystem(eng *Engine) {
	s := &resourceSystem{
		eng:       eng,
		resources: GetResourceVars(eng),
	}
	eng.RegisterActionHandler(s)
}

func (s *resourceSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case RegisterFont:
		s.registerFont(act)
	case SetVar:
		s.setVar(act)
	}
}

func (s *resourceSystem) registerFont(act RegisterFont) {
	s.resources.Fonts[act.ID] = struct{}{}
}

func (s *resourceSystem) setVar(act SetVar) {
	vars, ok := s.eng.Vars(act.Vars)
	if !ok {
		Warn("no such variables: %v", act.Vars)
		return
	}

	t := reflect.TypeOf(vars).Elem()
	v := reflect.ValueOf(vars).Elem()
	f, ok := t.FieldByName(act.ID)
	if !ok {
		Warn("no such %v variable: %v", act.Vars, act.ID)
		return
	}
	switch f.Type.Kind() {
	case reflect.String:
		v.FieldByName(act.ID).SetString(act.Val)
	case reflect.Int:
		i, err := strconv.Atoi(act.Val)
		if err != nil {
			Warn("not an integer: %v", act.Val)
			return
		}
		v.FieldByName(act.ID).SetInt(int64(i))
	default:
		Warn("cannot handle type %v: %v", f.Type.Kind(), act.Val)
	}
}

func GetResourceVars(store Store) *Resources {
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
