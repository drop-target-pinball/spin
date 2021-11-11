package app

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/system"
)

type Options struct {
	WithLogging bool
	WithAudio   bool
}

func DefaultOptions() Options {
	return Options{
		WithLogging: true,
		WithAudio:   true,
	}
}

func NewEngine(opt Options) *spin.Engine {
	eng := spin.NewEngine()
	system.NewScriptRunner(eng)
	if opt.WithLogging {
		system.NewLoggingConsole(eng)
	}
	if opt.WithAudio {
		system.NewAudioSDL(eng)
	}
	return eng
}
