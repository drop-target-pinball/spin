package spin

func RenderFrameScript(e *ScriptEnv, fn func(*ScriptEnv)) bool {
	for {
		fn(e)
		if done := e.Sleep(16); done {
			return true
		}
	}
}

func CountdownScript(e *ScriptEnv, timer *int, tickMs int, end Event) bool {
	for *timer > 0 {
		if done := e.Sleep(tickMs); done {
			return true
		}
		*timer -= 1
	}
	e.Post(end)
	return false
}

func WatchTimerScript(e *ScriptEnv, timer *int, fn func(v int)) bool {
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

func ScoreHurryUpScript(e *ScriptEnv, score *int, tickMs int, decScore int, endScore int) bool {
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
