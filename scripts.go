package spin

import (
	"context"
	"time"
)

func RenderFrameScript(e Env, fn func(Env)) context.CancelFunc {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e Env) {
		for {
			fn(e)
			if done := e.Sleep(FrameDuration); done {
				return
			}
		}
	})
	return cancel
}

func CountdownScript(e Env, timer *int, tickMs int, end Event) context.CancelFunc {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e Env) {
		for *timer > 0 {
			if done := e.Sleep(time.Duration(tickMs) * time.Millisecond); done {
				return
			}
			*timer -= 1
		}
		e.Post(end)
	})
	return cancel
}

func ScoreHurryUpScript(e Env, score *int, tickMs int, decScore int, endScore int, end Event) context.CancelFunc {
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e Env) {
		for *score > endScore {
			if done := e.Sleep(time.Duration(tickMs) * time.Millisecond); done {
				return
			}
			*score -= decScore
			if *score < endScore {
				*score = endScore
			}
		}
		e.Post(end)
	})
	return cancel
}
