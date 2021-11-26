package menu

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
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
	"HIGH SPEED 3",
	"PINGOLF",
}

var game int
var selectBlinkOn bool

func gameSelectFrame(e spin.Env) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.W = r.Width()
	g.Y = 7
	g.Font = FontPfTempestaFiveExtendedBold8
	r.Print(g, "GAME SELECT")
	g.Y = 18
	g.Font = FontPfTempestaFiveCompressedBold8
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

func gameSelectGameScript(e spin.Env) {
	game = 0
	for {
		gameSelectFrame(e)
		_, done := e.WaitFor(spin.Message{ID: MessageGameUpdated})
		if done {
			return
		}
	}
}

func gameSelectBlinkScript(e spin.Env) {
	game = 0
	for {
		selectBlinkOn = true
		gameSelectFrame(e)
		if done := e.Sleep(256 * time.Millisecond); done {
			return
		}

		selectBlinkOn = false
		gameSelectFrame(e)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
}

func selectModeScript(e spin.Env) {
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
	e.NewCoroutine(ctx, gameSelectGameScript)
	e.NewCoroutine(ctx, gameSelectBlinkScript)
	e.Do(spin.PlayMusic{ID: MusicSelectMode})

	for {
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
			spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
			spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
			spin.SwitchEvent{ID: jd.SwitchRightFireButton},
			spin.SwitchEvent{ID: jd.SwitchStartButton},
		)
		if done {
			cancel()
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchStartButton}) {
			break
		}
		switch evt {
		case spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton}:
			prev()
			e.Do(spin.PlaySound{ID: SoundSelectScroll})
		case spin.SwitchEvent{ID: jd.SwitchRightFlipperButton}:
			next()
			e.Do(spin.PlaySound{ID: SoundSelectScroll})
		case spin.SwitchEvent{ID: jd.SwitchLeftFireButton}:
			prev()
			e.Do(spin.PlaySound{ID: SoundSelectScroll})
		case spin.SwitchEvent{ID: jd.SwitchRightFireButton}:
			next()
			e.Do(spin.PlaySound{ID: SoundSelectScroll})
		}
		e.Post(spin.Message{ID: MessageGameUpdated})
	}
	cancel()
	e.Post(spin.Message{ID: MessageSelectDone})
}
