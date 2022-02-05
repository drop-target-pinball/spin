package spin

import (
	"log"
	"math/rand"
	"os"
	"reflect"
	"runtime/pprof"
	"time"

	"github.com/drop-target-pinball/coroutine"
)

func init() {
	rand.Seed(time.Now().Unix())
}

var AssetDir = os.Getenv("SPIN_DIR")

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
	GetVars(string) (interface{}, bool)
	RegisterVars(string, interface{})
}

type Engine struct {
	Config         Config
	Options        Options
	Actions        map[string]Action
	Events         map[string]Event
	displays       map[string]Display
	queue          []interface{}
	actionHandlers []ActionHandler
	eventHandlers  []EventHandler
	servers        []Server
	vars           map[string]interface{}
	coroutines     *coroutine.Group
	watchdog       chan struct{}
	done           chan struct{}
}

func NewEngine(config Config, options Options) *Engine {
	eng := &Engine{
		Config:         config,
		Options:        options,
		Actions:        make(map[string]Action),
		Events:         make(map[string]Event),
		displays:       make(map[string]Display),
		queue:          make([]interface{}, 0),
		actionHandlers: make([]ActionHandler, 0),
		eventHandlers:  make([]EventHandler, 0),
		servers:        make([]Server, 0),
		vars:           make(map[string]interface{}),
		coroutines:     coroutine.NewGroup(),
		watchdog:       make(chan struct{}),
		done:           make(chan struct{}),
	}
	eng.RegisterVars("config", config)
	registerResourceSystem(eng)
	registerActions(eng)
	registerEvents(eng)
	registerGameSystem(eng)
	registerScriptSystem(eng)
	RegisterTrackerSystem(eng)
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

func (e *Engine) Do(act Action) {
	switch a := act.(type) {
	case RegisterDisplay:
		e.displays[a.ID] = a.Display
	}
	e.queue = append(e.queue, act)
}

func (e *Engine) Post(evt Event) {
	e.queue = append(e.queue, evt)
}

func (e *Engine) GetVars(name string) (interface{}, bool) {
	vars, ok := e.vars[name]
	return vars, ok
}

func (e *Engine) NewCoroutine(fn func(*ScriptEnv)) {
	e.coroutines.NewCoroutine(func(co *coroutine.C) {
		fn(newScriptEnv(e, co))
	})
}

func (e *Engine) Display(id string) Display {
	d, ok := e.displays[id]
	if !ok {
		log.Panicf("no such display: %v", id)
	}
	return d
}

func (e *Engine) Run() {
	ticker := time.NewTicker(16670 * time.Microsecond)
	watchdog := newWatchdog(1 * time.Second)

	defer func() {
		watchdog.Stop()
		e.coroutines.Stop()
	}()

	for {
		watchdog.Reset()

		select {
		case <-e.done:
			return
		case <-ticker.C:
		}

		for _, s := range e.servers {
			s.Service()
			e.coroutines.Tick()
		}

		for len(e.queue) > 0 {
			watchdog.Reset()
			var item interface{}
			item, e.queue = e.queue[0], e.queue[1:]
			switch i := item.(type) {
			case Action:
				for _, h := range e.actionHandlers {
					h.HandleAction(i)
				}
			case Event:
				for _, h := range e.eventHandlers {
					h.HandleEvent(i)
				}
				e.coroutines.Post(i)
			}
		}
	}
}

func (e *Engine) Stop() {
	e.done <- struct{}{}
}

// FIXME: maybe allow the coroutine watchdog to have a function handler?

type watchdog struct {
	reset chan struct{}
	done  chan struct{}
}

func newWatchdog(timeout time.Duration) *watchdog {
	w := &watchdog{
		reset: make(chan struct{}, 1),
		done:  make(chan struct{}, 1),
	}
	go func() {
		for {
			select {
			case <-w.reset:
			case <-w.done:
				return
			case <-time.After(timeout):
				println("watchdog timer expired")
				profile := pprof.Lookup("goroutine")
				profile.WriteTo(os.Stdout, 1)
				dumpLog()
				log.Panicf("watchdog panic")
			}
		}
	}()
	return w
}

func (w *watchdog) Reset() {
	w.reset <- struct{}{}
}

func (w *watchdog) Stop() {
	w.done <- struct{}{}
}
