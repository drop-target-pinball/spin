package coroutine

import (
	"context"
	"math"
	"time"
)

type Selector interface {
	Key() interface{}
}

type timeout struct{}

func (s timeout) Key() interface{} {
	return s
}

type cancel struct{}

func (s cancel) Key() interface{} {
	return s
}

var (
	Cancel  = cancel{}
	Timeout = timeout{}
)

type waitCond struct {
	timer     <-chan time.Time
	selectors []Selector
}

type Context struct {
	ctx    context.Context
	yield  chan waitCond
	resume chan Selector
}

func (c *Context) WaitForUntil(d time.Duration, s ...Selector) Selector {
	c.yield <- waitCond{timer: time.After(d), selectors: s}
	select {
	case s := <-c.resume:
		return s
	case <-c.ctx.Done():
		return Cancel
	}
}

func (c *Context) WaitFor(d time.Duration) bool {
	c.yield <- waitCond{timer: time.After(d)}
	select {
	case <-c.resume:
		return false
	case <-c.ctx.Done():
		return true
	}
}

func (c *Context) WaitUntil(s ...Selector) Selector {
	return c.WaitForUntil(math.MaxInt64, s...)
}

type Runner struct {
	active []*coroutine
}

type coroutine struct {
	ctx      context.Context
	cancel   context.CancelFunc
	waitCond waitCond
	yield    chan waitCond
	resume   chan Selector
}

func NewRunner() *Runner {
	return &Runner{}
}

func (e *Runner) Create(parent context.Context, fn func(*Context)) {
	cr := &coroutine{}

	cr.ctx, cr.cancel = context.WithCancel(parent)
	cr.yield = make(chan waitCond)
	cr.resume = make(chan Selector)

	context := &Context{
		ctx:    cr.ctx,
		yield:  cr.yield,
		resume: cr.resume,
	}
	go func() {
		fn(context)
		cr.yield <- waitCond{}
		cr.cancel()
	}()

	cr.waitCond = <-cr.yield

	for i, entry := range e.active {
		if entry == nil {
			e.active[i] = cr
			return
		}
	}
	e.active = append(e.active, cr)
}

func (e *Runner) Service() {
	for i, entry := range e.active {
		if entry == nil {
			continue
		}
		select {
		case <-entry.waitCond.timer:
			entry.resume <- Timeout
			entry.waitCond = <-entry.yield
		case <-entry.ctx.Done():
			e.active[i] = nil
		default:
		}
	}
}

func (e *Runner) Post(s Selector) {
	for i, entry := range e.active {
		if entry == nil {
			continue
		}
		select {
		case <-entry.ctx.Done():
			e.active[i] = nil
			continue
		default:
			for _, wait := range entry.waitCond.selectors {
				if wait == s {
					entry.resume <- s
					entry.waitCond = <-entry.yield
					continue
				}
			}
		}
	}
}
