package coroutine

import (
	"context"
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

var active []*coroutine

type coroutine struct {
	ctx      context.Context
	cancel   context.CancelFunc
	waitCond waitCond
	yield    chan waitCond
	resume   chan Selector
}

func New(parent context.Context, fn func(*Context)) {
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

	for i, entry := range active {
		if entry == nil {
			active[i] = cr
			return
		}
	}

	if active == nil {
		active = make([]*coroutine, 0)
	}
	active = append(active, cr)
}

func Service() {
	for i, entry := range active {
		if entry == nil || entry.waitCond.timer == nil {
			continue
		}
		select {
		case <-entry.waitCond.timer:
			entry.resume <- timeout{}
			entry.waitCond = <-entry.yield
		case <-entry.ctx.Done():
			active[i] = nil
		default:
		}
	}
}

func Post(s Selector) {
	for i, entry := range active {
		if entry == nil {
			continue
		}
		select {
		case <-entry.ctx.Done():
			active[i] = nil
			continue
		default:
			for _, wait := range entry.waitCond.selectors {
				if wait.Key() == s.Key() {
					entry.resume <- s
					entry.waitCond = <-entry.yield
					continue
				}
			}
		}
	}
}
