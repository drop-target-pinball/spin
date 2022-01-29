package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

func attractModeScript(e *spin.ScriptEnv) {

	e.Do(proc.DriverSchedule{ID: jd.LampStartButton, Schedule: proc.Blink})
	for {
		e.Do(spin.PlayScript{ID: ScriptAttractModeSlide})
		evt, done := e.WaitFor(
			spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
			spin.SwitchEvent{ID: e.Config.SwitchStartButton},
		)
		if done {
			e.Do(spin.StopScript{ID: ScriptAttractModeSlide})
			return
		}
		switch evt {
		case spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton}:
			attractPreviousSlide(e)
		case spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton}:
			attractNextSlide(e)
		case spin.SwitchEvent{ID: e.Config.SwitchStartButton}:
			e.Do(spin.StopScript{ID: ScriptAttractModeSlide})
			return
		}
	}
}

var attractSlides = []func(*spin.ScriptEnv) bool{
	attractGameOver,
	attractDropTargetPinball,
	attractJudgeDreddRemix,
	attractFreePlay,
}

func attractGameOver(e *spin.ScriptEnv) bool {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.AnchorY = spin.AnchorMiddle
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "GAME OVER")

	return e.Sleep(4000)
}

func attractDropTargetPinball(e *spin.ScriptEnv) bool {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 7
	g.Font = spin.FontPfArmaFive8
	r.Print(g, "DROP TARGET PINBALL")
	g.Y = 18
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "PRESENTS")

	return e.Sleep(4000)
}

func attractJudgeDreddRemix(e *spin.ScriptEnv) bool {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 7
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "JUDGE DREDD")
	g.Y = 18
	r.Print(g, "REMIX")

	return e.Sleep(4000)
}

func freePlayPanel(e *spin.ScriptEnv, blinkOn bool) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 7
	g.Font = spin.FontPfRondaSevenBold8
	if blinkOn {
		r.Print(g, "PRESS START")
	}
	g.Y = 18
	r.Print(g, "FREE PLAY")
}

func attractFreePlay(e *spin.ScriptEnv) bool {
	s := spin.NewSequencer(e)

	s.DoFunc(func() { freePlayPanel(e, true) })
	s.Sleep(200)
	s.DoFunc(func() { freePlayPanel(e, false) })
	s.Sleep(100)
	s.LoopN(5)
	s.Run()

	freePlayPanel(e, true)
	return e.Sleep(2500)
}

func attractModeSlideScript(e *spin.ScriptEnv) {
	vars := GetVars(e)
	for {
		fn := attractSlides[vars.AttractModeSlide]
		if done := fn(e); done {
			return
		}
		attractNextSlide(e)
	}
}

func attractNextSlide(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.AttractModeSlide += 1
	if vars.AttractModeSlide >= len(attractSlides) {
		vars.AttractModeSlide = 0
	}
}

func attractPreviousSlide(e *spin.ScriptEnv) {
	vars := GetVars(e)
	vars.AttractModeSlide -= 1
	if vars.AttractModeSlide < 0 {
		vars.AttractModeSlide = len(attractSlides) - 1
	}
}
