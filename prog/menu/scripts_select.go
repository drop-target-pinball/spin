package menu

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
	"github.com/drop-target-pinball/spin/prog/system"
)

const (
	MessageGameUpdated  = "menu.MessageGameUpdated"
	MessageGameSelected = "menu.MessageGameSelected"
	MessageExit         = "menu.Exit"
)

type selectVars struct {
	blinkOn bool
}

func selectGameMenuFrame(e spin.Env, local *selectVars) {
	r, g := e.Display("").Renderer("")
	vars := system.GetVars(e)

	r.Fill(spin.ColorBlack)
	g.W = r.Width()
	g.Y = 7
	g.Font = builtin.FontPfTempestaFiveExtendedBold8
	r.Print(g, "GAME SELECT")
	g.Y = 18
	g.Font = builtin.FontPfTempestaFiveCompressedBold8
	r.Print(g, vars.Games[vars.Game])

	if local.blinkOn {
		g.W = 0
		g.X = 20
		g.Y = 18
		r.Print(g, ">>")
		g.X = 96
		r.Print(g, "<<")
	}
}

func clearFrame(e spin.Env) {
	r, _ := e.Display("").Renderer("")
	r.Fill(spin.ColorBlack)
}

func selectGameMenuScript(e spin.Env, local *selectVars) {
	for {
		selectGameMenuFrame(e, local)
		_, done := e.WaitFor(spin.Message{ID: MessageGameUpdated})
		if done {
			return
		}
	}
}

func selectGameMenuBlinkScript(e spin.Env, local *selectVars) {
	for {
		local.blinkOn = true
		selectGameMenuFrame(e, local)
		if done := e.Sleep(256 * time.Millisecond); done {
			return
		}

		local.blinkOn = false
		selectGameMenuFrame(e, local)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
}

func selectGameScript(e spin.Env) {
	e.Do(spin.DriverPWM{ID: jd.LampSuperGameButton, On: 127, Off: 127})
	defer e.Do(spin.DriverOff{ID: jd.LampSuperGameButton})

	vars := system.GetVars(e)

	next := func() {
		vars.Game += 1
		if vars.Game >= len(vars.Games) {
			vars.Game = 0
		}
	}

	prev := func() {
		vars.Game -= 1
		if vars.Game < 0 {
			vars.Game = len(vars.Games) - 1
		}
	}

	local := &selectVars{}
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, func(e spin.Env) { selectGameMenuScript(e, local) })
	e.NewCoroutine(ctx, func(e spin.Env) { selectGameMenuBlinkScript(e, local) })
	e.Do(spin.PlayMusic{ID: MusicSelectMode})

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchStartButton},
			spin.SwitchEvent{ID: jd.SwitchSuperGameButton},
		)
		if done {
			e.Do(spin.StopMusic{ID: MusicSelectMode})
			cancel()
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchSuperGameButton}) {
			e.Do(spin.StopMusic{ID: MusicSelectMode})
			cancel()
			e.Post(spin.Message{ID: MessageExit})
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchStartButton}) {
			break
		}
		switch evt {
		case spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton}:
			prev()
			e.Do(spin.PlaySound{ID: SoundScroll})
		case spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton}:
			next()
			e.Do(spin.PlaySound{ID: SoundScroll})
		}
		e.Post(spin.Message{ID: MessageGameUpdated})
	}
	cancel()

	local.blinkOn = false
	selectGameMenuFrame(e, local)
	e.Do(spin.DriverOn{ID: e.Config.LampStartButton})
	e.Do(spin.PlaySound{ID: SoundSelect})
	e.Do(spin.FadeOutMusic{Time: 1500})
	if done := e.Sleep(1500 * time.Millisecond); done {
		e.Do(spin.StopMusic{ID: MusicSelectMode})
		return
	}

	clearFrame(e)
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	e.Post(spin.Message{ID: MessageGameSelected})
}
