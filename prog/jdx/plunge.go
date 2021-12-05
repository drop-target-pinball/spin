package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func useFireButtonFrame(e spin.Env, n int) {
	r, g := e.Display("").Renderer()

	chevronsL := []string{
		"   ",
		"  <",
		" <<",
		"<<<",
	}

	chevronsR := []string{
		"   ",
		">  ",
		">> ",
		">>>",
	}

	r.Clear()
	g.Font = FontPfRondaSevenBold8
	g.W = r.Width()
	g.Y = 7
	r.Print(g, "USE")
	g.Y = 18
	r.Print(g, "FIRE BUTTON")

	g.W = 0
	g.H = r.Height()
	g.Y = 0
	g.X = 4
	r.Print(g, chevronsL[n])
	g.X = 110
	r.Print(g, chevronsR[n])
}

func useFireButtonVideo(e spin.Env) {
	game := spin.GameVars(e)
	game.HideScore = true
	defer func() { game.HideScore = false }()

	for i := 0; i < 7*4; i++ {
		useFireButtonFrame(e, i%4)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
}

func useFireButtonScript(e spin.Env) {
	ctx, cancel := e.Derive()
	defer cancel()

	if done := e.Sleep(7 * time.Second); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechUseFireButtonToLaunchBall})
	for {
		e.NewCoroutine(ctx, useFireButtonVideo)
		if done := e.Sleep(13 * time.Second); done {
			e.Do(spin.StopSpeech{ID: SpeechUseFireButtonToLaunchBall})
			return
		}
	}
}

func plungeScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicPlungeLoop})
	ctx, cancel := e.Derive()
	e.NewCoroutine(ctx, useFireButtonScript)

	e.Do(spin.PlaySpeech{ID: SpeechLawMasterComputerOnlineWelcomeAboard})
	e.Do(spin.DriverPulse{ID: jd.CoilTrough})
	_, done := e.WaitFor(
		spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
		spin.SwitchEvent{ID: jd.SwitchRightFireButton},
	)
	cancel()
	if done {
		return
	}

	e.Do(spin.DriverPulse{ID: jd.CoilRightShooterLane})
	e.Do(spin.PlayMusic{ID: MusicMain})
}
