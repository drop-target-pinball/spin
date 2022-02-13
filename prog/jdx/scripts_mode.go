package jdx

import "github.com/drop-target-pinball/spin"

func ModeIntroScript(e *spin.ScriptEnv, r spin.Renderer, line1 string, line2 string, line3 string) bool {
	text := [3]string{line1, line2, line3}

	s := spin.NewSequencer(e)

	s.DoFunc(func() { ModeIntroPanel(e, r, true, text) })
	s.Sleep(250)
	s.DoFunc(func() { ModeIntroPanel(e, r, false, text) })
	s.Sleep(100)
	s.LoopN(9)

	return s.Run()
}

func ModeAndBlinkingScoreScript(e *spin.ScriptEnv, r spin.Renderer, title string, score int) bool {
	s := spin.NewSequencer(e)

	s.DoFunc(func() { ModeAndBlinkingScorePanel(e, r, title, score, true) })
	s.Sleep(250)
	s.DoFunc(func() { ModeAndBlinkingScorePanel(e, r, title, score, false) })
	s.Sleep(100)
	s.LoopN(6)

	return s.Run()
}
