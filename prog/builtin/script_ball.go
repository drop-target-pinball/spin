package builtin

// func ballClearJam(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.SwitchEvent{ID: e.Config.SwitchTroughJam}); done {
// 			return
// 		}
// 		evt, done := e.WaitForUntil(1000*time.Millisecond, spin.SwitchEvent{ID: e.Config.SwitchTroughJam, Released: true})
// 		if done {
// 			return
// 		}
// 		if evt != nil {
// 			continue
// 		}
// 		e.Do(spin.DriverPulse{ID: e.Config.CoilTrough})
// 	}
// }

// func ballDrainScript(e spin.Env) {
// 	game := spin.GetGameVars(e)
// 	for {
// 		if _, done := e.WaitFor(spin.SwitchEvent{ID: e.Config.SwitchDrain}); done {
// 			return
// 		}
// 		if game.BallsInPlay == 0 {
// 			continue
// 		}
// 		game.BallsInPlay -= 1
// 		e.Post(spin.BallDrainEvent{BallsInPlay: game.BallsInPlay})
// 	}
// }

// func ballLaunch(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.BallAddedEvent{}); done {
// 			return
// 		}
// 		e.Do(spin.DriverPulse{ID: e.Config.CoilTrough})
// 	}
// }

// func ballWillDrainScript(e spin.Env) {
// 	events := make([]coroutine.Selector, len(e.Config.SwitchWillDrain))
// 	for i, sw := range e.Config.SwitchWillDrain {
// 		events[i] = spin.SwitchEvent{ID: sw}
// 	}

// 	for {
// 		if _, done := e.WaitFor(events...); done {
// 			return
// 		}
// 		e.Post(spin.BallWillDrainEvent{})
// 	}
// }

// func ballTrackerScript(e spin.Env) {
// 	ctx, _ := e.Derive()
// 	e.NewCoroutine(ctx, ballClearJam)
// 	e.NewCoroutine(ctx, ballDrainScript)
// 	e.NewCoroutine(ctx, ballLaunch)
// 	e.NewCoroutine(ctx, ballWillDrainScript)

// 	for {
// 		if _, done := e.WaitFor(spin.Done{}); done {
// 			return
// 		}
// 	}
// }
