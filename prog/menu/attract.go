package menu

const (
	MessageAttractAdvance = "menu.MessageAttractAdvance"
	MessageAttractDone    = "menu.MessageAttractDone"
)

// var attractScripts = []string{
// 	ScriptGameOver,
// 	ScriptDropTargetPinball,
// 	ScriptFreePlay,
// }

// func gameOverFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer()
// 	defer r.Unlock()

// 	r.Clear()
// 	g.Y = 2
// 	g.W = r.Width()
// 	g.H = r.Height()
// 	g.Font = FontPfRondaSevenBold8
// 	r.Print(g, "GAME OVER")
// }

// func gameOverScript(ctx context.Context, e spin.Env) {
// 	gameOverFrame(e)
// 	if done := spin.Wait(ctx, 4000*time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func dropTargetPinballFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer()
// 	defer r.Unlock()

// 	r.Clear()
// 	g.W = r.Width()
// 	g.Y = 7
// 	g.Font = FontPfArmaFive8
// 	r.Println(g, "DROP TARGET PINBALL")
// 	g.Y = 18
// 	g.Font = FontPfRondaSevenBold8
// 	r.Println(g, "PRESENTS")
// }

// func dropTargetPinballScript(ctx context.Context, e spin.Env) {
// 	dropTargetPinballFrame(e)
// 	if done := spin.Wait(ctx, 4000*time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func freePlayFrame(e spin.Env, blinkOn bool) {
// 	r, g := e.Display("").Renderer()
// 	defer r.Unlock()

// 	r.Clear()
// 	g.W = r.Width()
// 	g.Y = 7
// 	g.Font = FontPfRondaSevenBold8
// 	if blinkOn {
// 		r.Print(g, "PRESS START")
// 	}
// 	g.Y = 18
// 	r.Print(g, "FREE PLAY")
// }

// func freePlayScript(ctx context.Context, e spin.Env) {
// 	for i := 0; i < 5; i++ {
// 		freePlayFrame(e, true)
// 		if done := spin.Wait(ctx, 200*time.Millisecond); done {
// 			return
// 		}

// 		freePlayFrame(e, false)
// 		if done := spin.Wait(ctx, 100*time.Millisecond); done {
// 			return
// 		}
// 	}
// 	freePlayFrame(e, true)
// 	if done := spin.Wait(ctx, 2500*time.Millisecond); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageAttractAdvance})
// }

// func attractModeScript(ctx context.Context, e spin.Env) {
// 	script := 0

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

// 	for {
// 		e.Do(spin.PlayScript{ID: attractScripts[script]})
// 		evt, done := spin.WaitForEvents(ctx, e, []spin.Event{
// 			spin.Message{ID: MessageAttractAdvance},
// 			spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
// 			spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
// 			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
// 			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
// 			spin.SwitchEvent{ID: jd.SwitchStartButton},
// 		})
// 		if done {
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			return
// 		}
// 		switch evt {
// 		case spin.Message{ID: MessageAttractAdvance}:
// 			next()
// 		case spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton}:
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			prev()
// 		case spin.SwitchEvent{ID: jd.SwitchRightFlipperButton}:
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			next()
// 		case spin.SwitchEvent{ID: jd.SwitchLeftFireButton}:
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			prev()
// 		case spin.SwitchEvent{ID: jd.SwitchRightFireButton}:
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			next()
// 		case spin.SwitchEvent{ID: jd.SwitchStartButton}:
// 			e.Do(spin.StopScript{ID: attractScripts[script]})
// 			e.Post(spin.Message{ID: MessageAttractDone})
// 			return
// 		}
// 	}
// }
