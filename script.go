package spin

import (
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

type Env interface {
	Do(Action)
	Post(Event)
	// EventQueue() chan Event
	WaitForUntil(time.Duration, ...coroutine.Selector) coroutine.Selector
	WaitFor(time.Duration) bool
	WaitUntil(...coroutine.Selector) coroutine.Selector

	Display(string) Display
	// Int(string, string) int
	// SetInt(string, string, int)
	// AddInt(string, string, int)
	// String(string, string) string
	// SetString(string, string, string)
}

type Script func(Env)

// func Wait(ctx context.Context, d time.Duration) bool {
// 	select {
// 	case <-time.After(d):
// 		return false
// 	case <-ctx.Done():
// 		return true
// 	}
// }

// func WaitForEventsUntil(ctx context.Context, e Env, d time.Duration, watching []Event) (Event, bool) {
// 	timer := time.After(d)
// 	for {
// 		select {
// 		case event := <-e.EventQueue():
// 			for _, w := range watching {
// 				if event == w {
// 					return event, false
// 				}
// 			}
// 		case <-timer:
// 			return nil, false
// 		case <-ctx.Done():
// 			return nil, true
// 		}
// 	}
// }

// func WaitForEventUntil(ctx context.Context, e Env, d time.Duration, watching Event) (Event, bool) {
// 	return WaitForEventsUntil(ctx, e, d, []Event{watching})
// }

// func WaitForEvents(ctx context.Context, e Env, watching []Event) (Event, bool) {
// 	return WaitForEventsUntil(ctx, e, math.MaxInt64, watching)
// }

// func WaitForEvent(ctx context.Context, e Env, watching Event) (Event, bool) {
// 	return WaitForEventsUntil(ctx, e, math.MaxInt64, []Event{watching})
// }
