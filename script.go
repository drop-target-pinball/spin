package spin

import (
	"context"
	"time"
)

type Env interface {
	Do(Action)
	Post(Event)
	EventQueue() chan Event
	Display(string) Display
}

type Script func(context.Context, Env)

func Wait(ctx context.Context, d time.Duration) bool {
	select {
	case <-time.After(d):
		return false
	case <-ctx.Done():
		return true
	}
}

func WaitForSwitch(ctx context.Context, e Env, id string) (bool, SwitchEvent) {
	for {
		select {
		case event := <-e.EventQueue():
			sw, ok := event.(SwitchEvent)
			if ok && sw.ID == id {
				return false, sw
			}
		case <-ctx.Done():
			return true, SwitchEvent{}
		}
	}
}

func WaitForMessage(ctx context.Context, e Env, id string) (bool, Message) {
	for {
		select {
		case event := <-e.EventQueue():
			msg, ok := event.(Message)
			if ok && msg.ID == id {
				return false, msg
			}
		case <-ctx.Done():
			return true, Message{}
		}
	}
}

func WaitForSwitchUntil(ctx context.Context, e Env, id string, d time.Duration) (bool, SwitchEvent) {
	timer := time.After(d)
	for {
		select {
		case event := <-e.EventQueue():
			sw, ok := event.(SwitchEvent)
			if ok && sw.ID == id {
				return false, sw
			}
		case <-timer:
			return false, SwitchEvent{}
		case <-ctx.Done():
			return true, SwitchEvent{}
		}
	}
}

func WaitForEventsUntil(ctx context.Context, e Env, d time.Duration, watching []Event) (bool, Event) {
	timer := time.After(d)
	for {
		select {
		case event := <-e.EventQueue():
			for _, w := range watching {
				if event == w {
					return false, event
				}
			}
		case <-timer:
			return false, nil
		case <-ctx.Done():
			return true, nil
		}
	}
}
