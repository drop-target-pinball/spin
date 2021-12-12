package menu

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	MessageGameUpdated = "menu.MessageGameUpdated"
	MessageSelectDone  = "menu.MessageSelectDone"
)

const (
	VariableGame = "menu.VariableGame"
)

var games = []string{
	"DREDD REMIX",
	"MEGAMAN 3",
	"DR MARIO",
	"PINGOLF",
	"PRACTICE",
}

var game int
var selectBlinkOn bool

func selectGameMenuFrame(e spin.Env) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = builtin.FontPfTempestaFiveExtendedBold8
	r.Print(g, "GAME SELECT")
	g.Y = 18
	g.Font = builtin.FontPfTempestaFiveCompressedBold8
	r.Print(g, games[game])

	if selectBlinkOn {
		g.W = 0
		g.X = 20
		g.Y = 18
		r.Print(g, ">>")
		g.X = 96
		r.Print(g, "<<")
	}
}

func clearFrame(e spin.Env) {
	r, _ := e.Display("").Renderer()
	r.Clear()
}

func selectGameMenuScript(e spin.Env) {
	game = 0
	for {
		selectGameMenuFrame(e)
		_, done := e.WaitFor(spin.Message{ID: MessageGameUpdated})
		if done {
			return
		}
	}
}

func selectGameMenuBlinkScript(e spin.Env) {
	game = 0
	for {
		selectBlinkOn = true
		selectGameMenuFrame(e)
		if done := e.Sleep(256 * time.Millisecond); done {
			return
		}

		selectBlinkOn = false
		selectGameMenuFrame(e)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
}

func selectGameScript(e spin.Env) {
	next := func() {
		game += 1
		if game >= len(games) {
			game = 0
		}
	}

	prev := func() {
		game -= 1
		if game < 0 {
			game = len(games) - 1
		}
	}

	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, selectGameMenuScript)
	e.NewCoroutine(ctx, selectGameMenuBlinkScript)
	e.Do(spin.PlayMusic{ID: MusicSelectMode})

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchStartButton},
		)
		if done {
			e.Do(spin.StopMusic{ID: MusicSelectMode})
			cancel()
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

	selectBlinkOn = false
	selectGameMenuFrame(e)
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

	e.Post(spin.Message{ID: MessageSelectDone})
}
