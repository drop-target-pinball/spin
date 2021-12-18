package coroutine

import (
	"context"
	"time"
)

type ID int

var (
	nextID = ID(1)
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

type waitCond struct {
	timer     <-chan time.Time
	selectors []Selector
}

type Context struct {
	ctx    context.Context
	yield  chan waitCond
	resume chan Selector
}

func (c *Context) WaitFor(s ...Selector) (Selector, bool) {
	c.yield <- waitCond{selectors: s}
	select {
	case s := <-c.resume:
		return s, false
	case <-c.ctx.Done():
		return cancel{}, true
	}
}

func (c *Context) WaitForUntil(d time.Duration, s ...Selector) (Selector, bool) {
	c.yield <- waitCond{timer: time.After(d), selectors: s}
	select {
	case s := <-c.resume:
		if s == (timeout{}) {
			return nil, false
		}
		return s, false
	case <-c.ctx.Done():
		return cancel{}, true
	}
}

func (c *Context) Sleep(d time.Duration) bool {
	c.yield <- waitCond{timer: time.After(d)}
	select {
	case <-c.resume:
		return false
	case <-c.ctx.Done():
		return true
	}
}

func (c *Context) Derive() (context.Context, context.CancelFunc) {
	return context.WithCancel(c.ctx)
}

var active map[ID]*coroutine

type coroutine struct {
	id       ID
	ctx      context.Context
	cancel   context.CancelFunc
	waitCond waitCond
	yield    chan waitCond
	resume   chan Selector
}

func New(parent context.Context, fn func(*Context)) ID {
	if active == nil {
		active = make(map[ID]*coroutine)
	}

	cr := &coroutine{id: nextID}
	nextID++

	cr.ctx, cr.cancel = context.WithCancel(parent)
	cr.yield = make(chan waitCond)
	cr.resume = make(chan Selector)

	active[cr.id] = cr

	context := &Context{
		ctx:    cr.ctx,
		yield:  cr.yield,
		resume: cr.resume,
	}
	go func() {
		fn(context)
		select {
		case cr.yield <- waitCond{}:
		default:
		}
		cr.cancel()
	}()

	cr.waitCond = <-cr.yield
	return cr.id
}

func Service() {
	for id, entry := range active {
		select {
		case <-entry.ctx.Done():
			delete(active, id)
			entry.waitCond.timer = nil
			continue
		default:
		}

		if entry.waitCond.timer == nil {
			continue
		}

		select {
		case <-entry.waitCond.timer:
			entry.resume <- timeout{}
			entry.waitCond = <-entry.yield
		default:
		}
	}
}

func Post(s Selector) {
	for id, entry := range active {
		select {
		case <-entry.ctx.Done():
			delete(active, id)
			entry.waitCond.selectors = []Selector{}
		default:
		}

		for _, wait := range entry.waitCond.selectors {
			if wait.Key() == s.Key() {
				entry.resume <- s
				entry.waitCond = <-entry.yield
				continue
			}
		}
	}
}

func Cancel(id ID) {
	cr, ok := active[id]
	if ok {
		cr.cancel()
	}
}

func IsActive(id ID) bool {
	_, ok := active[id]
	return ok
}
