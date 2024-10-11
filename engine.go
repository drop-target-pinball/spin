package spin

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"time"

	"github.com/drop-target-pinball/spin/v2/msg"
)

type Provider interface {
	Provides() []string
	Handle(Message)
	Service(time.Time)
}

type Service interface {
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
	Dir       string
	client    *Client
	providers []Provider
	settings  Settings
	ticker    *time.Ticker
	sendQueue []Packet
	running   bool
}

func NewEngine(s Settings) *Engine {
	e := &Engine{settings: s}
	e.Dir = os.Getenv("SPIN_DIR")
	if e.Dir == "" {
		e.Dir = "."
	}
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

func (e *Engine) Abort(format string, args ...any) {
	e.Alert(format, args...)
	if err := e.flush(); err != nil {
		log.Print(err)
	}
	log.Fatalf(format, args...)
}

func (e *Engine) PathTo(tail string) string {
	return path.Join(e.Dir, tail)
}

func (e *Engine) Run() error {
	if e.running {
		return fmt.Errorf("already running")
	}
	e.running = true

	c := NewClient(e.settings.Addr)
	e.client = c
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

		if err := e.flush(); err != nil {
			return err
		}
	}
}

func (e *Engine) flush() error {
	if e.client == nil {
		return nil
	}
	for _, p := range e.sendQueue {
		if err := e.client.Send(p); err != nil {
			return fmt.Errorf("queue unavailable: %v", err)
		}
	}
	return nil
}
