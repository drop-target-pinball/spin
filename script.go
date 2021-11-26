package spin

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin/coroutine"
)

type Env interface {
	Do(Action)
	Post(Event)
	Sleep(time.Duration) bool
	WaitFor(...coroutine.Selector) (coroutine.Selector, bool)
	WaitForUntil(time.Duration, ...coroutine.Selector) (coroutine.Selector, bool)
	Display(string) Display
	Derive() (context.Context, context.CancelFunc)
	NewCoroutine(ctx context.Context, scr Script)
}

type Script func(Env)
