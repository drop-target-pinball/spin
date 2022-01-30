package gamepad

import (
	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/jdx"
	"github.com/drop-target-pinball/spin/sdl"
)

func RegisterGamePad(eng *spin.Engine) {
	eng.Do(sdl.RegisterButton{
		Button:     "a",
		ActionDown: spin.AddBall{},
	})
	eng.Do(sdl.RegisterButton{
		Button:     "x",
		ActionDown: spin.PlayScript{ID: jd.ScriptBallCollect},
	})
	eng.Do(sdl.RegisterButton{
		Button:     "y",
		ActionDown: spin.PlayScript{ID: jd.ScriptBallSearch},
	})
	eng.Do(sdl.RegisterButton{
		Button:     "start",
		ActionDown: spin.PlayScript{ID: jdx.ScriptBase},
	})
}
