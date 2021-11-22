package menu

const (
	MessageGameUpdated = "menu.MessageGameUpdated"
	MessageSelectDone  = "menu.MessageSelectDone"
)

const (
	VariableGame = "menu.VariableGame"
)

// var games = []string{
// 	"DREDD REMIX",
// 	"MEGAMAN 3",
// 	"DR MARIO",
// 	"HIGH SPEED 3",
// 	"PINGOLF",
// }

// func gameSelectFrame(e spin.Env, blinkOn bool) {
// 	r, g := e.Display("").Renderer()
// 	defer r.Unlock()

// 	r.Clear()
// 	g.W = r.Width()
// 	g.Y = 7
// 	g.Font = FontPfTempestaFiveExtendedBold8
// 	r.Print(g, "GAME SELECT")
// 	g.Y = 18
// 	g.Font = FontPfTempestaFiveCompressedBold8
// 	game := games[e.Int(spin.System, VariableGame)]
// 	r.Print(g, game)

// 	if blinkOn {
// 		g.W = 0
// 		g.X = 20
// 		g.Y = 18
// 		r.Print(g, ">>")
// 		g.X = 96
// 		r.Print(g, "<<")
// 	}
// }

// func gameSelectScript(ctx context.Context, e spin.Env) {
// 	e.SetInt(spin.System, VariableGame, 0)
// 	for {
// 		gameSelectFrame(e, true)
// 		if _, done := spin.WaitForEventUntil(ctx, e, 256*time.Millisecond, spin.Message{ID: MessageGameUpdated}); done {
// 			return
// 		}

// 		gameSelectFrame(e, false)
// 		if _, done := spin.WaitForEventUntil(ctx, e, 100*time.Millisecond, spin.Message{ID: MessageGameUpdated}); done {
// 			return
// 		}
// 	}
// }

// func selectModeScript(ctx context.Context, e spin.Env) {
// 	next := func() {
// 		game := e.Int(spin.System, VariableGame)
// 		game += 1
// 		if game >= len(games) {
// 			game = 0
// 		}
// 		e.SetInt(spin.System, VariableGame, game)
// 	}

// 	prev := func() {
// 		game := e.Int(spin.System, VariableGame)
// 		game -= 1
// 		if game < 0 {
// 			game = len(games) - 1
// 		}
// 		e.SetInt(spin.System, VariableGame, game)
// 	}

// 	e.Do(spin.PlayScript{ID: ScriptGameSelect})
// 	e.Do(spin.PlayMusic{ID: MusicSelectMode})
// 	for {
// 		evt, done := spin.WaitForEvents(ctx, e, []spin.Event{
// 			spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
// 			spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
// 			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
// 			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
// 			spin.SwitchEvent{ID: jd.SwitchStartButton},
// 		})
// 		if done {
// 			return
// 		}
// 		if evt == (spin.SwitchEvent{ID: jd.SwitchStartButton}) {
// 			break
// 		}
// 		switch evt {
// 		case spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton}:
// 			prev()
// 			e.Do(spin.PlaySound{ID: SoundSelectScroll})
// 		case spin.SwitchEvent{ID: jd.SwitchRightFlipperButton}:
// 			next()
// 			e.Do(spin.PlaySound{ID: SoundSelectScroll})
// 		case spin.SwitchEvent{ID: jd.SwitchLeftFireButton}:
// 			prev()
// 			e.Do(spin.PlaySound{ID: SoundSelectScroll})
// 		case spin.SwitchEvent{ID: jd.SwitchRightFireButton}:
// 			next()
// 			e.Do(spin.PlaySound{ID: SoundSelectScroll})
// 		}
// 		e.Post(spin.Message{ID: MessageGameUpdated})
// 	}

// 	e.Post(spin.Message{ID: MessageSelectDone})
// }
