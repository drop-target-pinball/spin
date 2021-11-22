package boot

import (
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

// func splashScreenFrame(e spin.Env) {
// 	r, g := e.Display("").Renderer()
// 	defer r.Unlock()

// 	g.Font = FontPfTempestaFiveCompressedBold8
// 	g.W = r.Width()
// 	g.PaddingV = 2
// 	g.Y = 4
// 	r.Println(g, "SUPER PINBALL SYSTEM")
// 	r.Println(g, spin.Version)
// 	r.Println(g, spin.Date)
// }

// func splashScreenScript(ctx context.Context, e spin.Env) {

// 	e.Do(spin.StopAudio{})
// 	splashScreenFrame(e)
// 	e.Do(spin.PlayMusic{ID: MusicSplashScreen})
// 	if _, done := spin.WaitForEventsUntil(ctx, e, 8*time.Second, []spin.Event{
// 		spin.SwitchEvent{ID: jd.SwitchLeftFlipperButton},
// 		spin.SwitchEvent{ID: jd.SwitchRightFlipperButton},
// 	}); done {
// 		return
// 	}
// 	e.Post(spin.Message{ID: MessageDone})
// }

func RegisterScripts(eng *spin.Engine) {
	//eng.Do(spin.RegisterScript{ID: ScriptSplashScreen, Script: splashScreenScript})
}
