package menu

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

const (
	MessageGameSelected = "menu.MessageGameSelected"
	MessageExit         = "menu.Exit"
)

func selectGameScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	e.Do(proc.DriverSchedule{ID: jd.LampSuperGameButton, Schedule: proc.BlinkSchedule})
	defer e.Do(spin.DriverOff{ID: jd.LampSuperGameButton})

	vars := GetVars(e)

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

	blink := true

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.RenderFrameLoop(e, func(e *spin.ScriptEnv) {
			selectGameMenuPanel(e, r, blink)
		})
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.DoFunc(func() { blink = true })
		s.Sleep(256)
		s.DoFunc(func() { blink = false })
		s.Sleep(100)
		s.Loop()

		s.Run()
	})

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
			return
		}
		if evt == (spin.SwitchEvent{ID: jd.SwitchSuperGameButton}) {
			e.Do(spin.StopMusic{ID: MusicSelectMode})
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
	}
	e.Do(spin.PlayScript{ID: ScriptGameSelected})
}

func gameSelectedScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	selectGameMenuPanel(e, r, false)
	e.Do(spin.DriverOn{ID: e.Config.LampStartButton})
	e.Do(spin.PlaySound{ID: SoundSelect})
	e.Do(spin.FadeOutMusic{Time: 1500})
	if done := e.Sleep(1500); done {
		e.Do(spin.StopMusic{ID: MusicSelectMode})
		return
	}

	r.Fill(spin.ColorOff)
	if done := e.Sleep(1000); done {
		return
	}

	e.Post(spin.Message{ID: MessageGameSelected})
}

func selectGameMenuPanel(e *spin.ScriptEnv, r spin.Renderer, blinkOn bool) {
	g := r.Graphics()
	vars := GetVars(e)

	r.Fill(spin.ColorOff)
	g.Y = 7
	g.Font = spin.FontPfTempestaFiveExtendedBold8
	r.Print(g, "GAME SELECT")
	g.Y = 18
	g.Font = spin.FontPfTempestaFiveCompressedBold8
	r.Print(g, vars.Games[vars.Game])

	g.AnchorX = spin.AnchorLeft
	if blinkOn {
		g.X = 20
		g.Y = 18
		r.Print(g, ">>")
		g.X = 96
		r.Print(g, "<<")
	}
}
