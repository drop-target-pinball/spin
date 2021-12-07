package boot

import (
	"time"

	"github.com/drop-target-pinball/spin"
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
	g.Y = 4
	r.Print(g, "SUPER PINBALL SYSTEM")
	g.Y = 14
	r.Print(g, spin.Version)
	g.Y = 24
	r.Print(g, spin.Date)
}

func splashScreenScript(e spin.Env) {
	e.Do(spin.StopAudio{})

	for _, gi := range e.Config.GI {
		e.Do(spin.DriverOn{ID: gi})
	}

	splashScreenFrame(e)
	e.Do(spin.PlayMusic{ID: MusicSplashScreen})
	e.Do(spin.DriverPWM{ID: e.Config.LampStartButton, TimeOn: 127, TimeOff: 127})
	evt, done := e.WaitForUntil(8*time.Second,
		spin.SwitchEvent{ID: e.Config.SwitchLeftFlipperButton},
		spin.SwitchEvent{ID: e.Config.SwitchRightFlipperButton},
		spin.SwitchEvent{ID: e.Config.SwitchStartButton},
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
