package app

import (
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
	"github.com/drop-target-pinball/spin/prog/builtin"
	"github.com/drop-target-pinball/spin/sdl"
)

type Options struct {
	WithPROC       bool
	WithLogging    bool
	WithAudio      bool
	WithVirtualDMD bool
	WithEOS        bool
}

func DefaultOptions() Options {
	return Options{
		WithLogging:    true,
		WithAudio:      true,
		WithVirtualDMD: true,
	}
}

func NewEngine(appOptions Options, spinOptions spin.Options) *spin.Engine {
	eng := spin.NewEngine(jd.Config, spinOptions)
	if appOptions.WithLogging {
		spin.RegisterLoggingSystem(eng)
	}
	if appOptions.WithAudio {
		sdl.RegisterAudioSystem(eng)
	}
	if appOptions.WithVirtualDMD {
		opts := sdl.DefaultOptionsDotMatrix()
		sdl.RegisterDotMatrixSystem(eng, opts)
	}
	if appOptions.WithPROC {
		opts := proc.Options{
			MachType:                wpc.MachType,
			DMDConfig:               wpc.DMDConfigDefault,
			SwitchConfig:            wpc.SwitchConfigDefault,
			DefaultCoilPulseTime:    25, // milliseconds
			DefaultFlasherPulseTime: 20, // milliseconds
		}
		proc.RegisterSystem(eng, opts)
	} else {
		proc.RegisterNullSystem(eng)
	}
	sdl.RegisterDisplaySystem(eng, spin.DisplayOptions{Width: 128, Height: 32})
	sdl.RegisterInputSystem(eng)
	builtin.Load(eng)

	return eng
}
