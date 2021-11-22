package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	MessageSniperAdvance = "jdx.MessageSniperAdvance"
	MessageSniperTimeout = "jdx.MessageSniperTimeout"
)

var sniperScore int

func sniperScoreFrame(e spin.Env, blinkOn bool) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.Font = FontPfArmaFive8
	r.Print(g, "SNIPER")
	g.Y = 12

	if blinkOn {
		g.Font = FontBm8
		score := spin.Sprintf("%10d", sniperScore)
		r.Print(g, score)
	}
}

func sniperScoreCountdownVideoScript(e spin.Env) {
	sniperScore = 20_000_000
	modeText := [3]string{"SNIPER", "SHOOT", "SNIPER TOWER"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	sniperScoreFrame(e, true)
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	expires := time.Now().Add(30 * time.Second)
	for time.Now().Before(expires) {
		sniperScore -= 78_330
		sniperScoreFrame(e, true)
		if done := e.Sleep(160 * time.Millisecond); done {
			return
		}
	}

	sniperScore = 5_000_000
	sniperScoreFrame(e, true)
	if done := e.Sleep(2000 * time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: MessageSniperTimeout})
}

func sniperScoreCountdownAudioScript(e spin.Env) {
	e.Do(spin.VolumeMusic{Mul: 0.5})
	e.Do(spin.PlaySpeech{ID: SpeechSniperIsShootingIntoCrowdFromJohnsonTower})
	if done := e.Sleep(3 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechSniperIsShootingIntoCrowdFromJohnsonTower})
		e.Do(spin.VolumeMusic{Mul: 2})
		return
	}

	e.Do(spin.VolumeMusic{Mul: 2})
	if done := e.Sleep(1 * time.Second); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechShootSniperTower})
	if done := e.Sleep(1 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechShootSniperTower})
		return
	}

	for {
		e.Do(spin.PlaySound{ID: SoundGunLoadSniper})
		if done := e.Sleep(1500 * time.Millisecond); done {
			return
		}
		e.Do(spin.PlaySound{ID: SoundGunFire})
		if done := e.Sleep(1500 * time.Millisecond); done {
			return
		}
	}
}

func sniperScoreCountdownScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptSniperScoreCountdownVideo})
	e.Do(spin.PlayScript{ID: ScriptSniperScoreCountdownAudio})

	evt, done := e.WaitFor(
		spin.Message{ID: MessageSniperTimeout},
		spin.SwitchEvent{ID: jd.SwitchRightPopper},
	)
	e.Do(spin.StopScript{ID: ScriptSniperScoreCountdownVideo})
	e.Do(spin.StopScript{ID: ScriptSniperScoreCountdownAudio})
	if done || evt == (spin.Message{ID: MessageSniperTimeout}) {
		return
	}
	e.Post(spin.Message{ID: MessageSniperAdvance})
}

func sniperTakedownVideoScript(e spin.Env) {
	for i := 0; i < 6; i++ {
		sniperScoreFrame(e, true)
		if done := e.Sleep(250 * time.Millisecond); done {
			return
		}
		sniperScoreFrame(e, false)
		if done := e.Sleep(100 * time.Millisecond); done {
			return
		}
	}
	e.Post(spin.Message{ID: MessageSniperAdvance})
}

func sniperTakedownAudioScript(e spin.Env) {
	e.Do(spin.VolumeMusic{Mul: 0.5})
	e.Do(spin.PlaySound{ID: SoundSuccess})
	if done := e.Sleep(1500 * time.Millisecond); done {
		return
	}
	e.Do(spin.VolumeMusic{Mul: 2})
}

func sniperTakedownScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptSniperTakedownVideo})
	e.Do(spin.PlayScript{ID: ScriptSniperTakedownAudio})
	e.WaitFor(spin.Message{ID: MessageSniperAdvance})

	e.Do(spin.StopScript{ID: ScriptSniperTakedownVideo})
	e.Do(spin.StopScript{ID: ScriptSniperTakedownAudio})
}

func sniperFallFrame(e spin.Env, seconds int) {
	r, g := e.Display("").Renderer()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.Font = FontPfArmaFive8
	r.Print(g, "SNIPER")
	g.Y = 12

	g.Font = FontBm8
	r.Print(g, "%v", seconds)
}

func sniperFallCountdownVideoScript(e spin.Env) {
	seconds := 10

	sniperFallFrame(e, seconds)
	if done := e.Sleep(200 * time.Millisecond); done {
		return
	}

	for seconds > 0 {
		if done := e.Sleep(1500 * time.Millisecond); done {
			return
		}
		seconds -= 1
		sniperFallFrame(e, seconds)
	}
	e.Post(spin.Message{ID: MessageSniperTimeout})
}

func sniperFallCountdownAudioScript(e spin.Env) {
	e.Do(spin.PlaySpeech{ID: SpeechShootSniperTower})
	if done := e.Sleep(1750 * time.Millisecond); done {
		e.Do(spin.StopSpeech{ID: SpeechShootSniperTower})
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechAaaaah})
	if done := e.Sleep(3 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechAaaaah})
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechItsALongWayDown})
	if done := e.Sleep(2500 * time.Millisecond); done {
		e.Do(spin.StopSpeech{ID: SpeechItsALongWayDown})
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechAaaaah})
	if done := e.Sleep(3 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechAaaaah})
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechICanSeeMyHouseFromHere})
	if done := e.Sleep(2 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechICanSeeMyHouseFromHere})
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechAaaaah})
	if done := e.Sleep(3 * time.Second); done {
		e.Do(spin.StopSpeech{ID: SpeechAaaaah})
		return
	}
	e.Do(spin.StopSpeech{ID: SpeechAaaaah})
}

func sniperFallCountdownScript(e spin.Env) {
	e.Do(spin.PlayScript{ID: ScriptSniperFallCountdownVideo})
	e.Do(spin.PlayScript{ID: ScriptSniperFallCountdownAudio})
	evt, done := e.WaitFor(
		spin.Message{ID: MessageSniperTimeout},
		spin.SwitchEvent{ID: jd.SwitchRightPopper},
	)
	e.Do(spin.StopScript{ID: ScriptSniperFallCountdownVideo})
	e.Do(spin.StopScript{ID: ScriptSniperFallCountdownAudio})
	if done || evt == (spin.Message{ID: MessageSniperTimeout}) {
		return
	}

	e.Post(spin.Message{ID: MessageSniperAdvance})
}

func sniperSplatTimeoutScript(e spin.Env) {
	e.Do(spin.StopMusic{ID: MusicMode1})
	e.Do(spin.PlayMusic{ID: MusicMain})
	e.Do(spin.PlaySound{ID: SoundSniperSplat})
	if done := e.Sleep(1000 * time.Millisecond); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: SpeechSniperEliminated})
	if done := e.Sleep(2000 * time.Millisecond); done {
		e.Do(spin.StopSpeech{ID: SpeechSniperEliminated})
		return
	}
	e.Post(spin.Message{ID: MessageSniperAdvance})
}

func sniperModeScript(e spin.Env) {
	e.Do(spin.StopAudio{})
	e.Do(spin.PlayMusic{ID: MusicMode1})
	e.Do(spin.PlayScript{ID: ScriptSniperScoreCountdown})
	evt, done := e.WaitFor(
		spin.Message{ID: MessageSniperTimeout},
		spin.Message{ID: MessageSniperAdvance},
	)
	if done || evt == (spin.Message{ID: MessageSniperTimeout}) {
		e.Do(spin.StopMusic{ID: MusicMode1})
		return
	}

	e.Do(spin.PlayScript{ID: ScriptSniperTakedown})
	if _, done := e.WaitFor(spin.Message{ID: MessageSniperAdvance}); done {
		e.Do(spin.StopMusic{ID: MusicMode1})
		return
	}

	e.Do(spin.PlayScript{ID: ScriptSniperFallCountdown})
	evt, done = e.WaitFor(
		spin.Message{ID: MessageSniperTimeout},
		spin.Message{ID: MessageSniperAdvance},
	)
	if done || evt == (spin.Message{ID: MessageSniperTimeout}) {
		e.Do(spin.StopMusic{ID: MusicMode1})
		return
	}

	success := evt == spin.Message{ID: MessageSniperAdvance}
	e.Do(spin.PlayScript{ID: ScriptSniperSplatTimeout})
	if _, done := e.WaitFor(spin.Message{ID: MessageSniperAdvance}); done || !success {
		return
	}

	e.Do(spin.PlayScript{ID: ScriptSniperTakedown})
	e.WaitFor(spin.Message{ID: MessageSniperAdvance})
}
