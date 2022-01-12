package spin

import (
	"time"
)

func RenderFrameLoop(e *ScriptEnv, fn func(*ScriptEnv)) bool {
	for {
		fn(e)
		if done := e.Sleep(16); done {
			fn(e)
			return true
		}
	}
}

func CountdownLoop(e *ScriptEnv, timer *int, tickMs int, end Event) bool {
	for *timer > 0 {
		if done := e.Sleep(tickMs); done {
			return true
		}
		*timer -= 1
	}
	e.Post(end)
	return false
}

func WatcherTimerLoop(e *ScriptEnv, timer *int, fn func(v int)) bool {
	seen := *timer
	for {
		if done := e.Sleep(16); done {
			return true
		}
		if *timer != seen {
			seen = *timer
			fn(*timer)
		}
	}
}

func ScoreHurryUpLoop(e *ScriptEnv, score *int, tickMs int, decScore int, endScore int) bool {
	for *score > endScore {
		if done := e.Sleep(tickMs); done {
			return true
		}
		*score -= decScore
		if *score < endScore {
			*score = endScore
		}
	}
	return false
}

func WaitForBallArrival(e *ScriptEnv, sw string, timeMs int) bool {
	for {
		if _, done := e.WaitFor(SwitchEvent{ID: sw}); done {
			return true
		}
		evt, done := e.WaitForUntil(500*time.Millisecond, SwitchEvent{ID: sw, Released: true})
		if done {
			return true
		}
		if evt == nil {
			return false
		}
	}
}

func WaitForBallDeparture(e *ScriptEnv, sw string, timeMs int) bool {
	for {
		if _, done := e.WaitFor(SwitchEvent{ID: sw, Released: true}); done {
			return true
		}
		evt, done := e.WaitForUntil(500*time.Millisecond, SwitchEvent{ID: sw})
		if done {
			return true
		}
		if evt == nil {
			return false
		}
	}
}
