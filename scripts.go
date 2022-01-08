package spin

func RenderFrameScript(e *ScriptEnv, fn func(*ScriptEnv)) {
	for {
		fn(e)
		if done := e.Sleep(16); done {
			return
		}
	}
}

func CountdownScript(e *ScriptEnv, timer *int, tickMs int, end Event) {
	for *timer > 0 {
		if done := e.Sleep(tickMs); done {
			return
		}
		*timer -= 1
	}
	e.Post(end)
}

// func ScoreHurryUpScript(e Env, score *int, tickMs int, decScore int, endScore int, end Event) context.CancelFunc {
// 	ctx, cancel := e.Derive()
// 	e.NewCoroutine(ctx, func(e Env) {
// 		for *score > endScore {
// 			if done := e.Sleep(time.Duration(tickMs) * time.Millisecond); done {
// 				return
// 			}
// 			*score -= decScore
// 			if *score < endScore {
// 				*score = endScore
// 			}
// 		}
// 		e.Post(end)
// 	})
// 	return cancel
// }
