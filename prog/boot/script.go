package boot

import (
	"github.com/drop-target-pinball/spin"
)

// Scripts
const (
	ScriptSplashScreen = "boot.ScriptSplashScreen"
)

func splashScreenScript(e *spin.ScriptEnv) {
	r := e.Display("").Open()
	defer r.Close()

	e.Do(spin.StopAudio{})

	for _, gi := range e.Config.GI {
		e.Do(spin.DriverOn{ID: gi})
	}

	splashScreenPanel(e, r)
	e.Do(spin.PlayMusic{ID: MusicSplashScreen, Loops: 1})
	e.Do(spin.DriverPWM{ID: e.Config.LampStartButton, On: 127, Off: 127})
	evt, done := e.WaitForUntil(8000,
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
}

func splashScreenPanel(e *spin.ScriptEnv, r spin.Renderer) {
	g := r.Graphics()
	g.Font = spin.FontPfTempestaFiveCompressedBold8
	g.Y = 4
	r.Print(g, "SUPER PINBALL SYSTEM")
	g.Y = 14
	r.Print(g, spin.Version)
	g.Y = 24
	r.Print(g, spin.Date)
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptSplashScreen,
		Script: splashScreenScript,
	})
}
