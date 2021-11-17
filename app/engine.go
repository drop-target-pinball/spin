package app

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/system"
)

type Options struct {
	WithLogging    bool
	WithAudio      bool
	WithVirtualDMD bool
}

func DefaultOptions() Options {
	return Options{
		WithLogging:    true,
		WithAudio:      true,
		WithVirtualDMD: true,
	}
}

func NewEngine(opt Options) *spin.Engine {
	eng := spin.NewEngine()
	system.RegisterScriptRunner(eng)
	if opt.WithLogging {
		system.RegisterLoggingConsole(eng)
	}
	if opt.WithAudio {
		system.RegisterAudioSDL(eng)
	}
	if opt.WithVirtualDMD {
		sdlOpts := system.DefaultOptionsDotMatrixSDL()
		system.RegisterDotMatrixSDL(eng, sdlOpts)
	}
	system.RegisterDisplaySDL(eng, spin.DisplayOptions{Width: 128, Height: 32})
	system.RegisterInputSDL(eng)
	return eng
}
