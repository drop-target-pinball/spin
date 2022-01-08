package service

// type fontMode struct {
// 	offset   int32
// 	fonts    []string
// 	selected int
// }

// func fontPreviewFrame(e spin.Env, fm *fontMode) {
// 	r, g := e.Display("").Renderer("")

// 	r.Fill(spin.ColorBlack)
// 	g.Font = fm.fonts[fm.selected]
// 	g.X = fm.offset
// 	r.Print(g, "0123456,789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

// 	g.Font = builtin.Font04B_03_7px
// 	g.Y = 26
// 	g.X = 0
// 	r.Print(g, fm.fonts[fm.selected])
// }

// func fontPreviewVideoScript(e spin.Env, fm *fontMode) {
// 	for {
// 		fontPreviewFrame(e, fm)
// 		if done := e.Sleep(spin.FrameDuration); done {
// 			return
// 		}
// 	}
// }

// func fontPreviewScript(e spin.Env) {
// 	fm := fontMode{
// 		fonts:    make([]string, 0),
// 		selected: 0,
// 	}
// 	rsrc := spin.GetResourceVars(e)
// 	for name, _ := range rsrc.Fonts {
// 		fm.fonts = append(fm.fonts, name)
// 	}
// 	sort.Strings(fm.fonts)

// 	next := func() {
// 		fm.selected += 1
// 		if fm.selected >= len(fm.fonts) {
// 			fm.selected = 0
// 		}
// 	}

// 	prev := func() {
// 		fm.selected -= 1
// 		if fm.selected < 0 {
// 			fm.selected = len(fm.fonts) - 1
// 		}
// 	}

// 	ctx, cancel := e.Derive()
// 	e.NewCoroutine(ctx, func(e spin.Env) { fontPreviewVideoScript(e, &fm) })

// 	for {
// 		evt, done := e.WaitFor(
// 			spin.SwitchEvent{ID: e.Config.SwitchExitServiceButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchNextServiceButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchPreviousServiceButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
// 			spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
// 		)
// 		if done {
// 			cancel()
// 			return
// 		}
// 		switch evt {
// 		case spin.SwitchEvent{ID: e.Config.SwitchExitServiceButton}:
// 			cancel()
// 			return
// 		case spin.SwitchEvent{ID: e.Config.SwitchNextServiceButton}:
// 			next()
// 		case spin.SwitchEvent{ID: e.Config.SwitchPreviousServiceButton}:
// 			prev()
// 		case spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton}:
// 			fm.offset -= 1
// 		case spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton}:
// 			fm.offset += 1
// 		}
// 	}
// }
