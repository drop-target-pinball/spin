package spin

import (
	"os"
	"reflect"
	"time"
)

var AssetDir = os.Getenv("SPIN_ASSET_DIR")

type ActionHandler interface {
	HandleAction(Action)
}

type EventHandler interface {
	HandleEvent(Event)
}

type Server interface {
	Service()
}

type Store interface {
	RegisterVars(string, interface{})
	Vars(string) (interface{}, bool)
}

type Engine struct {
	Config         Config
	Actions        map[string]Action
	Events         map[string]Event
	actionQueue    []Action
	eventQueue     []Event
	actionHandlers []ActionHandler
	eventHandlers  []EventHandler
	servers        []Server
	running        bool
	vars           map[string]interface{}
}

func NewEngine(config Config) *Engine {
	eng := &Engine{
		Config:         config,
		Actions:        make(map[string]Action),
		Events:         make(map[string]Event),
		actionQueue:    make([]Action, 0),
		eventQueue:     make([]Event, 0),
		actionHandlers: make([]ActionHandler, 0),
		eventHandlers:  make([]EventHandler, 0),
		servers:        make([]Server, 0),
		vars:           make(map[string]interface{}),
	}
	registerResourceSystem(eng)
	registerActions(eng)
	registerEvents(eng)
	registerGameSystem(eng)
	registerScriptSystem(eng)
	return eng
}

func (e *Engine) RegisterActionHandler(h ActionHandler) {
	e.actionHandlers = append(e.actionHandlers, h)
}

func (e *Engine) RegisterEventHandler(h EventHandler) {
	e.eventHandlers = append(e.eventHandlers, h)
}

func (e *Engine) RegisterServer(s Server) {
	e.servers = append(e.servers, s)
}

func (e *Engine) RegisterAction(act Action) {
	t := reflect.TypeOf(act)
	name := t.Name()
	if _, exists := e.Actions[name]; exists {
		Warn("duplicate action: %v", name)
		return
	}
	e.Actions[name] = act
}

func (e *Engine) RegisterEvent(evt Event) {
	t := reflect.TypeOf(evt)
	name := t.Name()
	if _, exists := e.Actions[name]; exists {
		Warn("duplicate event: %v", name)
		return
	}
	e.Events[name] = evt
}

func (e *Engine) RegisterVars(name string, vars interface{}) {
	e.vars[name] = vars
}

func (e *Engine) Run() {
	e.running = true
	e.loop()
}

func (e *Engine) Do(act Action) {
	e.actionQueue = append(e.actionQueue, act)
}

func (e *Engine) Post(evt Event) {
	e.eventQueue = append(e.eventQueue, evt)
}

func (e *Engine) Vars(name string) (interface{}, bool) {
	vars, ok := e.vars[name]
	return vars, ok
}

func (e *Engine) loop() {
	ticker := time.NewTicker(16670 * time.Microsecond)
	for {
		<-ticker.C
		for len(e.actionQueue) > 0 {
			var act Action
			act, e.actionQueue = e.actionQueue[0], e.actionQueue[1:]
			for _, h := range e.actionHandlers {
				h.HandleAction(act)
			}
		}
		for len(e.eventQueue) > 0 {
			var evt Event
			evt, e.eventQueue = e.eventQueue[0], e.eventQueue[1:]
			for _, h := range e.eventHandlers {
				h.HandleEvent(evt)
			}
		}
		for _, s := range e.servers {
			s.Service()
		}
	}
}
