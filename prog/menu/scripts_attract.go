package menu

// const (
// 	MessageAttractAdvance = "menu.MessageAttractAdvance"
// 	MessageAttractDone    = "menu.MessageAttractDone"
// )

// var attractScripts = []spin.ScriptFn{
// 	gameOverScript,
// 	dropTargetPinballScript,
// 	superPinballSystemScript,
// 	freePlayScript,
// }

// func gameOverFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.AnchorY = spin.AnchorMiddle
// 	g.Y = r.Height() / 2
// 	g.Font = builtin.FontPfRondaSevenBold8
// 	r.Print(g, "GAME OVER")
// }

// func gameOverScript(e spin.Env) {
// 	gameOverFrame(e)
// 	if done := e.Sleep(4000 * time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func dropTargetPinballFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 7
// 	g.Font = builtin.FontPfArmaFive8
// 	r.Print(g, "DROP TARGET PINBALL")
// 	g.Y = 18
// 	g.Font = builtin.FontPfRondaSevenBold8
// 	r.Print(g, "PRESENTS")
// }

// func dropTargetPinballScript(e spin.Env) {
// 	dropTargetPinballFrame(e)
// 	if done := e.Sleep(4000 * time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func superPinballSystemFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 7
// 	g.Font = builtin.FontPfRondaSevenBold8
// 	r.Print(g, "SUPER")
// 	g.Y = 18
// 	r.Print(g, "PINBALL SYSTEM")
// }

// func superPinballSystemScript(e spin.Env) {
// 	superPinballSystemFrame(e)
// 	if done := e.Sleep(4000 * time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func freePlayFrame(e spin.Env, blinkOn bool) {
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Y = 7
// 	g.Font = builtin.FontPfRondaSevenBold8
// 	if blinkOn {
// 		r.Print(g, "PRESS START")
// 	}
// 	g.Y = 18
// 	r.Print(g, "FREE PLAY")
// }

// func freePlayScript(e spin.Env) {
// 	for i := 0; i < 5; i++ {
// 		freePlayFrame(e, true)
// 		if done := e.Sleep(200 * time.Millisecond); done {
// 			return
// 		}

// 		freePlayFrame(e, false)
// 		if done := e.Sleep(100 * time.Millisecond); done {
// 			return
// 		}
// 	}
// 	freePlayFrame(e, true)
// 	if done := e.Sleep(2500 * time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func attractModeScript(e spin.Env) {
// 	script := 0
// 	var ctx context.Context
// 	var cancel context.CancelFunc

// 	next := func() {
// 		script += 1
// 		if script >= len(attractScripts) {
// 			script = 0
// 		}
// 	}

// 	prev := func() {
// 		script -= 1
// 		if script < 0 {
// 			script = len(attractScripts) - 1
// 		}
// 	}

// 	e.Do(spin.DriverPWM{ID: e.Config.LampStartButton, On: 127, Off: 127})
// 	for {
// 		ctx, cancel = e.Derive()
// 		e.NewCoroutine(ctx, attractScripts[script])
// 		evt, done := e.WaitFor(
// 			spin.Message{ID: MessageAttractAdvance},
// 			spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchStartButton},
// 		)
// 		if done {
// 			cancel()
// 			return
// 		}
// 		switch evt {
// 		case spin.Message{ID: MessageAttractAdvance}:
// 			next()
// 		case spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton}:
// 			cancel()
// 			prev()
// 		case spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton}:
// 			cancel()
// 			next()
// 		case spin.SwitchEvent{ID: e.Config.SwitchStartButton}:
// 			cancel()
// 			e.Post(spin.Message{ID: MessageAttractDone})
// 			return
// 		}
// 	}
// }
