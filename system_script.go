package spin

import (
	"sort"

	"github.com/drop-target-pinball/coroutine"
)

const (
	ScriptGroupMode = "GroupMode"
	ScriptGroupBall = "GroupBall"
)

type ScriptFn func(*ScriptEnv)

type script struct {
	id     string
	fn     ScriptFn
	groups []string
	cancel coroutine.CancelFunc
}

type scriptSystem struct {
	eng     *Engine
	scripts map[string]*script
	//displays map[string]Display
}

func registerScriptSystem(eng *Engine) {
	sys := &scriptSystem{
		eng:     eng,
		scripts: make(map[string]*script),
		//displays: make(map[string]Display),
	}
	eng.RegisterActionHandler(sys)
}

func (s *scriptSystem) HandleAction(action Action) {
	switch act := action.(type) {
	case Debug:
		s.debug(act)
	// case RegisterDisplay:
	// 	s.registerDisplay(act)
	case RegisterScript:
		s.registerScript(act)
	case PlayScript:
		s.playScript(act)
	case StopScript:
		s.stopScript(act)
	case StopScriptGroup:
		s.stopScriptGroup(act)
	}
}

// func (s *scriptSystem) registerDisplay(act RegisterDisplay) {
// 	s.displays[act.ID] = act.Display
// }

func (s *scriptSystem) registerScript(a RegisterScript) {
	var groups []string
	if a.Groups != nil {
		groups = a.Groups
	}
	if a.Group != "" {
		groups = []string{a.Group}
	}
	s.scripts[a.ID] = &script{
		id:     a.ID,
		groups: groups,
		fn:     a.Script,
	}
}

func (s *scriptSystem) playScript(a PlayScript) {
	script, ok := s.scripts[a.ID]
	if !ok {
		Warn("no such script: %v", a.ID)
		return
	}
	if script.cancel != nil {
		script.cancel()
	}
	script.cancel = s.eng.coroutines.NewCoroutine(func(co *coroutine.C) {
		s.eng.Post(ScriptStartedEvent(a))
		script.fn(newScriptEnv(s.eng, co))
		s.eng.Post(ScriptFinishedEvent(a))
		script.cancel = nil
	})
}

func (s *scriptSystem) stopScript(act StopScript) {
	script, ok := s.scripts[act.ID]
	if !ok {
		Warn("no such script: %v", act.ID)
		return
	}
	if script.cancel != nil {
		script.cancel()
	}
}

func (s *scriptSystem) stopScriptGroup(act StopScriptGroup) {
	for _, script := range s.scripts {
		for _, group := range script.groups {
			if group == act.ID && script.cancel != nil {
				script.cancel()
				break
			}
		}
	}
}

func (s *scriptSystem) debug(evt Debug) {
	switch evt.ID {
	case "Scripts":
		s.debugScripts()
	}
}

func (s *scriptSystem) debugScripts() {
	running := make([]string, 0)
	for _, script := range s.scripts {
		if script.cancel != nil {
			running = append(running, script.id)
		}
	}
	sort.Strings(running)
	for _, name := range running {
		Log(name)
	}
}
