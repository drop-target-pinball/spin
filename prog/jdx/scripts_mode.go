package jdx

import "github.com/drop-target-pinball/spin"

// func defaultLeftShooterLaneScript(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftShooterLane}); done {
// 			return
// 		}
// 		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
// 		if done := e.Sleep(1 * time.Second); done {
// 			return
// 		}
// 		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
// 	}
// }

// func defaultLeftPopperScript(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftPopper}); done {
// 			return
// 		}
// 		for i := 0; i < 3; i++ {
// 			e.Do(spin.DriverPulse{ID: jd.FlasherSubwayExit})
// 			if done := e.Sleep(250 * time.Millisecond); done {
// 				return
// 			}
// 		}
// 		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
// 	}
// }

// func defaultRightPopperScript(e spin.Env) {
// 	for {
// 		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
// 			return
// 		}
// 		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
// 	}
// }

// func modeIntroVideo(e spin.Env, text [3]string) bool {
// 	for i := 0; i < 9; i++ {
// 		modeIntroPanel(e, true, text)
// 		if done := e.Sleep(250 * time.Millisecond); done {
// 			return done
// 		}
// 		modeIntroPanel(e, false, text)
// 		if done := e.Sleep(100 * time.Millisecond); done {
// 			return done
// 		}
// 	}
// 	return false
// }

func ModeIntroScript(e *spin.ScriptEnv, line1 string, line2 string, line3 string) bool {
	text := [3]string{line1, line2, line3}

	s := spin.NewSequencer(e)

	s.DoFunc(func() { ModeIntroPanel(e, true, text) })
	s.Sleep(250)
	s.DoFunc(func() { ModeIntroPanel(e, false, text) })
	s.Sleep(100)
	s.LoopN(9)

	return s.Run()
}

// func ModeAndBlinkingScoreSequence(e spin.Env, r spin.Renderer, title string, score int) bool {
// 	return spin.NewSequencer().
// 		Func(func() { modeAndBlinkingScorePanel(e, r, title, score, true) }).
// 		Sleep(250).
// 		Func(func() { modeAndBlinkingScorePanel(e, r, title, score, false) }).
// 		Sleep(100).
// 		Repeat(6).
// 		Run(e)
// }
