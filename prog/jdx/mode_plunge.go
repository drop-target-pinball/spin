package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

/*
SwitchEvent ID=jd.SwitchRightFireButton
*/

func plungeModeScript(e *spin.ScriptEnv) {
	game := spin.GetGameVars(e)
	vars := GetVars(e)

	vars.Mode = ModePlunge
	defer func() { vars.Mode = ModeNone }()

	e.Do(spin.PlayMusic{ID: MusicPlungeLoop})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)
		s.Sleep(7_000)
		s.DoScript(useFireButtonScript)
		s.Run()

		s = spin.NewSequencer(e)
		s.Sleep(13_000)
		s.DoScript(useFireButtonSilentScript)
		s.Loop()
		s.Run()
	})

	if game.Player == 1 && game.Ball == 1 && !game.IsExtraBall {
		e.Do(spin.PlaySpeech{ID: SpeechLawMasterComputerOnlineWelcomeAboard})
	}
	e.Do(spin.AddBall{})

	if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchRightFireButton}); done {
		return
	}

	e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Do(spin.PlaySound{ID: SoundMotorcycleStart})
}

func useFireButtonPanel(e *spin.ScriptEnv, r spin.Renderer, n int) {
	g := r.Graphics()

	// chevronsL := []string{
	// 	"   ",
	// 	"  <",
	// 	" <<",
	// 	"<<<",
	// }

	chevronsR := []string{
		"   ",
		">  ",
		">> ",
		">>>",
	}

	r.Fill(spin.ColorBlack)
	g.Font = spin.FontPfRondaSevenBold8
	g.Y = 7
	r.Print(g, "USE")
	g.Y = 18
	r.Print(g, "FIRE BUTTON")

	g.AnchorY = spin.AnchorMiddle
	g.Y = r.Height() / 2
	g.X = 4
	//r.Print(g, chevronsL[n])
	g.X = 110
	r.Print(g, chevronsR[n])
}

func useFireButtonAnimScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	i := 0
	s := spin.NewSequencer(e)
	s.DoFunc(func() {
		useFireButtonPanel(e, r, i%4)
		i += 1
	})
	s.Sleep(100)
	s.LoopN(7 * 4)
	s.Run()
}

func useFireButtonScript(e *spin.ScriptEnv) {
	e.Do(spin.PlaySpeech{ID: SpeechUseFireButtonToLaunchBall})
	useFireButtonAnimScript(e)
}

func useFireButtonSilentScript(e *spin.ScriptEnv) {
	useFireButtonAnimScript(e)
}
