package jdx

import (
	"context"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

func blackoutScoreFrame(e spin.Env) {
	player := spin.GetPlayerVars(e)
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, "BLACKOUT")
	g.Y = 12

	g.Font = builtin.Font14x10
	score := spin.FormatScore("%10d", player.Score)
	r.Print(g, score)
}

func blackoutVideoScript(e spin.Env) {
	vars := GetVars(e)
	vars.SniperScore = 20_000_000
	modeText := [3]string{"BLACKOUT", "EVERYTHING", "2X"}
	if done := modeIntroVideo(e, modeText); done {
		return
	}

	for {
		blackoutScoreFrame(e)
		if done := e.Sleep(spin.FrameDuration); done {
			return
		}
	}
}

func blackoutAudioScript(e spin.Env) {
	e.Do(spin.PlayMusic{ID: MusicMode1})
	e.Do(spin.PlaySpeech{ID: SpeechMegaCityOneIsBlackedOutBeOnTheAlertForLooters, Notify: true})
	if _, done := e.WaitFor(spin.SpeechFinishedEvent{}); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: SpeechSendBackupUnits})
}

func blackoutAwardJackpot(e spin.Env) {
	e.Do(spin.AwardScore{Val: ScoreBlackoutJackpot})
	e.Do(spin.MusicVolume{Mul: 0.25})
	e.Do(spin.PlaySound{ID: SoundBlackoutJackpot, Notify: true})
	if _, done := e.WaitFor(spin.SoundFinishedEvent{ID: SoundBlackoutJackpot}); done {
		e.Do(spin.MusicVolume{Mul: 4})
		return
	}
	e.Do(spin.MusicVolume{Mul: 4})
}

func blackoutWatchJackpot(e spin.Env) {
	var ctx context.Context
	var cancel context.CancelFunc

	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotTopLeftRamp}); done {
			return
		}
		if cancel != nil {
			cancel()
		}
		ctx, cancel = e.Derive()
		e.NewCoroutine(ctx, blackoutAwardJackpot)
	}
}

func blackoutModeScript(e spin.Env) {
	vars := GetVars(e)
	vars.Multiplier = 2

	ctx, cancel := e.Derive()

	defer func() {
		cancel()
		vars.Multiplier = 1
	}()

	e.NewCoroutine(ctx, blackoutVideoScript)
	e.NewCoroutine(ctx, blackoutAudioScript)
	e.NewCoroutine(ctx, blackoutWatchJackpot)
	e.Do(spin.AddBall{})
	if _, done := e.WaitFor(spin.BallDrainEvent{}); done {
		return
	}
	cancel()
	e.Do(spin.PlayMusic{ID: MusicMain})
}
