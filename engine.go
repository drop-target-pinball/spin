package spin

import (
	"time"
)

type ActionHandler interface {
	HandleAction(Action)
}

type EventHandler interface {
	HandleEvent(Event)
}

type Server interface {
	Service()
}

type Engine struct {
	actionQueue    chan Action
	eventQueue     chan Event
	actionHandlers []ActionHandler
	eventHandlers  []EventHandler
	servers        []Server
}

func NewEngine() *Engine {
	return &Engine{
		actionQueue:    make(chan Action, 1),
		eventQueue:     make(chan Event, 1),
		actionHandlers: make([]ActionHandler, 0),
		eventHandlers:  make([]EventHandler, 0),
		servers:        make([]Server, 0),
	}
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

func (e *Engine) Do(act Action) {
	e.actionQueue <- act
}

func (e *Engine) Post(evt Event) {
	e.eventQueue <- evt
}

func (e *Engine) loop() {
	ticker := time.NewTicker(100 * time.Millisecond)
	for {
		select {
		case act := <-e.actionQueue:
			for _, h := range e.actionHandlers {
				h.HandleAction(act)
			}
		case evt := <-e.eventQueue:
			for _, h := range e.eventHandlers {
				h.HandleEvent(evt)
			}
		case <-ticker.C:
			for _, s := range e.servers {
				s.Service()
			}
		}
	}
}

func (e *Engine) Start() {
	go e.loop()
}
