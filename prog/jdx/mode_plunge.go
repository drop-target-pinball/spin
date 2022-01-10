package jdx

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func plungeModeScript(e *spin.ScriptEnv) {
	game := spin.GetGameVars(e)

	e.Do(spin.PlayScript{ID: builtin.ScriptScore})
	e.Do(spin.PlayMusic{ID: MusicPlungeLoop})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		s := spin.NewSequencer(e)

		s.Sleep(7_000)
		s.Do(spin.PlaySpeech{ID: SpeechUseFireButtonToLaunchBall})
		s.WaitFor(spin.SpeechFinishedEvent{})

		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		if done := e.Sleep(7_000); done {
			return
		}

		s := spin.NewSequencer(e)
		s.Do(spin.PlayScript{ID: ScriptUseFireButton})
		s.Sleep(13_000)
		s.Loop()
		s.Run()
	})

	e.NewCoroutine(func(e *spin.ScriptEnv) {
		spin.RenderFrameScript(e, spin.ScorePanel)
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
	e.Post(spin.ScriptFinishedEvent{ID: ScriptPlungeMode})
}

func useFireButtonPanel(e *spin.ScriptEnv, n int) {
	r, g := e.Display("").Renderer(spin.LayerPriority)

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

func useFireButtonScript(e *spin.ScriptEnv) {
	r, _ := e.Display("").Renderer(spin.LayerPriority)
	defer r.Clear()

	i := 0
	s := spin.NewSequencer(e)
	s.DoFunc(func() {
		useFireButtonPanel(e, i%4)
		i += 1
	})
	s.Sleep(100)
	s.LoopN(7 * 4)
	s.Run()
}
