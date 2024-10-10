package spin

import (
	"errors"
	"fmt"
	"slices"
	"time"

	"github.com/drop-target-pinball/spin/v2/msg"
)

type Provider interface {
	Provides() []string
	Handle(Message)
	Service(time.Time)
}

type Settings struct {
	Addr         string
	TickInterval time.Duration
}

func DefaultSettings() Settings {
	return Settings{
		Addr:         DefaultRedisAddr,
		TickInterval: 16670 * time.Microsecond,
	}
}

type Engine struct {
	providers []Provider
	settings  Settings
	ticker    *time.Ticker
	sendQueue []Packet
}

func NewEngine(s Settings) *Engine {
	e := &Engine{settings: s}
	return e
}

func (e *Engine) Send(m Message) {
	header := Header{
		Chan: m.Chan(),
		Name: m.Name(),
	}
	e.sendQueue = append(e.sendQueue, Packet{header, m})
}

func (e *Engine) Alert(format string, args ...any) {
	e.Send(msg.Alert{
		Desc: fmt.Sprintf(format, args...),
	})
}

func (e *Engine) Run() error {
	c := NewClient(e.settings.Addr)
	e.ticker = time.NewTicker(e.settings.TickInterval)

	var subs []string
	for _, p := range e.providers {
		subs = append(subs, p.Provides()...)
	}
	slices.Sort(subs)
	subs = slices.Compact(subs)
	c.Subscribe(subs...)

	for {
		e.sendQueue = nil
		for {
			m, err := c.ReceiveWithTimeout(e.ticker.C)
			if errors.Is(err, ErrTimeout) {
				break
			}
			if err != nil {
				e.Alert("invalid message: %v", err)
				continue
			}
			for _, p := range e.providers {
				p.Handle(m.Body)
			}
		}

		t := time.Now()
		for _, p := range e.providers {
			p.Service(t)
		}

		for _, p := range e.sendQueue {
			if err := c.Send(p); err != nil {
				return fmt.Errorf("queue unavailable: %v", err)
			}
		}
	}
}
