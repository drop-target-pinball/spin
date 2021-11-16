package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
)

const (
	SniperCaughtDone = "SniperCaughtDone"
	SniperAdvance    = "SniperAdvance"
	SniperTimeout    = "SniperTimeout"
	SniperFallDone   = "SniperFallDone"
)

const (
	SniperScore = "SniperScore"
)

func sniperScoreFrame(e spin.Env, blinkOn bool) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.Font = PfArmaFive8
	r.Print(g, "SNIPER")
	g.Y = 12

	if blinkOn {
		g.Font = Bm8
		score := spin.Sprintf("%10d", e.Int(spin.Player, SniperScore))
		r.Print(g, score)
	}
}

func sniperScoreCountdownVideo(ctx context.Context, e spin.Env) {
	e.SetInt(spin.Player, SniperScore, 20_000_000)
	modeText := [3]string{"SNIPER", "SHOOT", "SNIPER TOWER"}
	if done := modeIntroVideo(ctx, e, modeText); done {
		return
	}

	sniperScoreFrame(e, true)
	if done := spin.Wait(ctx, 1000*time.Millisecond); done {
		return
	}

	expires := time.Now().Add(30 * time.Second)
	for time.Now().Before(expires) {
		e.AddInt(spin.Player, SniperScore, -78_330)
		sniperScoreFrame(e, true)
		if done := spin.Wait(ctx, 160*time.Millisecond); done {
			return
		}
	}

	e.SetInt(spin.Player, SniperScore, 5_000_000)
	sniperScoreFrame(e, true)
	if done := spin.Wait(ctx, 2000*time.Millisecond); done {
		return
	}
	e.Post(spin.Message{ID: SniperTimeout})
}

func sniperScoreCountdownAudio(ctx context.Context, e spin.Env) {
	e.Do(spin.PlayMusic{ID: ModeTheme1, Vol: 100})
	e.Do(spin.PlaySpeech{ID: SniperIsShootingIntoCrowdFromJohnsonTower})
	if done := spin.Wait(ctx, 4*time.Second); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: ShootSniperTower})
	if done := spin.Wait(ctx, 1*time.Second); done {
		return
	}

	for {
		e.Do(spin.PlaySound{ID: GunLoadSniper})
		if done := spin.Wait(ctx, 1500*time.Millisecond); done {
			return
		}
		e.Do(spin.PlaySound{ID: GunFire})
		if done := spin.Wait(ctx, 1500*time.Millisecond); done {
			return
		}
	}
}

func sniperScoreCountdown(ctx context.Context, e spin.Env) {
	e.Do(spin.PlayScript{ID: SniperScoreCountdownVideo})
	e.Do(spin.PlayScript{ID: SniperScoreCountdownAudio})
	defer e.Do(spin.StopScript{ID: SniperScoreCountdownVideo})
	defer e.Do(spin.StopScript{ID: SniperScoreCountdownAudio})

	evt, done := spin.WaitForEvents(ctx, e, []spin.Event{
		spin.Message{ID: SniperTimeout},
		spin.SwitchEvent{ID: jd.RightPopperSwitch},
	})
	if done || evt == (spin.Message{ID: SniperTimeout}) {
		return
	}
	e.Post(spin.Message{ID: SniperAdvance})
}

func sniperCaughtVideo(ctx context.Context, e spin.Env) {
	for i := 0; i < 6; i++ {
		sniperScoreFrame(e, true)
		if done := spin.Wait(ctx, 250*time.Millisecond); done {
			return
		}
		sniperScoreFrame(e, false)
		if done := spin.Wait(ctx, 100*time.Millisecond); done {
			return
		}
	}
	e.Post(spin.Message{ID: SniperAdvance})
}

func sniperCaughtAudio(ctx context.Context, e spin.Env) {
	e.Do(spin.VolumeMusic{Mul: 0.5})
	e.Do(spin.PlaySound{ID: Success})
	if done := spin.Wait(ctx, 1500*time.Millisecond); done {
		return
	}
	e.Do(spin.VolumeMusic{Mul: 2})
}

func sniperCaught(ctx context.Context, e spin.Env) {
	e.Do(spin.PlayScript{ID: SniperCaughtVideo})
	e.Do(spin.PlayScript{ID: SniperCaughtAudio})
	defer e.Do(spin.StopScript{ID: SniperCaughtVideo})
	defer e.Do(spin.StopScript{ID: SniperCaughtAudio})

	spin.WaitForEvent(ctx, e, spin.Message{ID: SniperAdvance})
}

func sniperFall(ctx context.Context, e spin.Env) {
	e.Do(spin.PlaySpeech{ID: ShootSniperTower})
	if done := spin.Wait(ctx, 1750*time.Millisecond); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: Ahhhhh})
	if done := spin.Wait(ctx, 3*time.Second); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: ItsALongWayDown})
	if done := spin.Wait(ctx, 2500*time.Millisecond); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: Ahhhhh})
	if done := spin.Wait(ctx, 3*time.Second); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: ICanSeeMyHouseFromHere})
	if done := spin.Wait(ctx, 2*time.Second); done {
		return
	}
	e.Do(spin.PlaySpeech{ID: Ahhhhh})
	if done := spin.Wait(ctx, 3*time.Second); done {
		return
	}
	e.Do(spin.StopSpeech{})
}

func sniperMode(ctx context.Context, e spin.Env) {
	// defer func() {
	// 	e.Do(spin.StopScript{ID: SniperCaught})
	// 	e.Do(spin.StopScript{ID: SniperFall})
	// }()

	// e.Do(spin.StopAudio{})
	// if done := sniperCountdown(ctx, e); done {
	// 	return
	// }

	// e.Do(spin.PlayScript{ID: SniperCaught})
	// if done, _ := spin.WaitForEvent(ctx, e, SniperCaughtDone); done {
	// 	return
	// }
	// e.Do(spin.PlayScript{ID: SniperFall})
	// if done, evt := spin.WaitForSwitchUntil(ctx, e, "popper", 16*time.Second); done || evt.ID == "" {
	// 	return
	// }
}
