package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

func defaultLeftShooterLaneScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftShooterLane}); done {
			return
		}
		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
		if done := e.Sleep(1 * time.Second); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
	}
}

func defaultLeftPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftPopper}); done {
			return
		}
		for i := 0; i < 3; i++ {
			e.Do(spin.DriverPulse{ID: jd.FlasherSubwayExit})
			if done := e.Sleep(250 * time.Millisecond); done {
				return
			}
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
	}
}

func defaultRightPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

func modeIntroVideo(e spin.Env, text [3]string) bool {
	for i := 0; i < 9; i++ {
		modeIntroPanel(e, true, text)
		if done := e.Sleep(250 * time.Millisecond); done {
			return done
		}
		modeIntroPanel(e, false, text)
		if done := e.Sleep(100 * time.Millisecond); done {
			return done
		}
	}
	return false
}

func ModeIntroSequence(e spin.Env, line1 string, line2 string, line3 string) *spin.Sequencer {
	text := [3]string{line1, line2, line3}

	return spin.NewSequencer().
		Func(func() { modeIntroPanel(e, true, text) }).
		Sleep(250).
		Func(func() { modeIntroPanel(e, false, text) }).
		Sleep(100).
		Repeat(9)
}

func ModeAndBlinkingScoreSequence(e spin.Env, r spin.Renderer, title string, score int) bool {
	return spin.NewSequencer().
		Func(func() { modeAndBlinkingScorePanel(e, r, title, score, true) }).
		Sleep(250).
		Func(func() { modeAndBlinkingScorePanel(e, r, title, score, false) }).
		Sleep(100).
		Repeat(6).
		Run(e)
}
