package boot

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

// Scripts
const (
	ScriptSplashScreen = "boot.ScriptSplashScreen"
)

// Messages
const (
	MessageDone = "boot.MessageDone"
)

func splashScreenFrame(e spin.Env) {
	r, g := e.Display("").Renderer()
	g.Font = FontPfTempestaFiveCompressedBold8
	g.W = r.Width()
	g.PaddingV = 2
	g.Y = 4
	r.Println(g, "SUPER PINBALL SYSTEM")
	r.Println(g, spin.Version)
	r.Println(g, spin.Date)
}

func splashScreenScript(e spin.Env) {
	e.Do(spin.StopAudio{})
	splashScreenFrame(e)
	e.Do(spin.PlayMusic{ID: MusicSplashScreen})
	evt, done := e.WaitForUntil(8*time.Second,
		spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
		spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
		spin.SwitchEvent{ID: jd.SwitchLeftFireButton},
		spin.SwitchEvent{ID: jd.SwitchRightFireButton},
		spin.SwitchEvent{ID: jd.SwitchStartButton},
	)
	if done {
		return
	}
	if evt != nil {
		e.Do(spin.StopMusic{ID: MusicSplashScreen})
	}
	e.Post(spin.Message{ID: MessageDone})
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{ID: ScriptSplashScreen, Script: splashScreenScript})
}
