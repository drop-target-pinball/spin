package spin

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

func WaitForBallArrivalLoop(e *ScriptEnv, sw string, timeMs int) bool {
	for {
		if _, done := e.WaitFor(SwitchEvent{ID: sw}); done {
			return true
		}
		evt, done := e.WaitForUntil(timeMs, SwitchEvent{ID: sw, Released: true})
		if done {
			return true
		}
		if evt == nil {
			return false
		}
	}
}

func WaitForBallArrivalFunc(e *ScriptEnv, sw string, timeMs int) func() bool {
	return func() bool {
		return WaitForBallArrivalLoop(e, sw, timeMs)
	}
}

func WaitForBallDepartureLoop(e *ScriptEnv, sw string, timeMs int) bool {
	for {
		if _, done := e.WaitFor(SwitchEvent{ID: sw, Released: true}); done {
			return true
		}
		evt, done := e.WaitForUntil(500, SwitchEvent{ID: sw})
		if done {
			return true
		}
		if evt == nil {
			return false
		}
	}
}

func ShotSequenceLoop(e *ScriptEnv, shot string, timeMs int, switches ...string) {
	for {
		if _, done := e.WaitFor(SwitchEvent{ID: switches[0]}); done {
			return
		}
		shotMade := true
		for _, sw := range switches[1:] {
			evt, done := e.WaitForUntil(timeMs, SwitchEvent{ID: sw})
			if done {
				return
			}
			if evt == nil {
				shotMade = false
				break
			}
		}
		if shotMade {
			e.Post(ShotEvent{ID: shot})
		}
	}
}
