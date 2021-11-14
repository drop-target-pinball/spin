package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
)

const (
	SniperCaughtDone = "SniperCaughtDone"
	SniperFallDone   = "SniperFallDone"
)

func sniperScoreFrame(e spin.Env, score int) {
	r, g := e.Display("").Renderer()
	defer r.Unlock()

	r.Clear()
	g.Y = 2
	g.W = r.Width()
	g.Font = PfArmaFive8
	r.Print(g, "SNIPER")
	g.Y = 12
	g.Font = PfRondaSevenBold8
	r.Print(g, "%v", score)
}

func sniperScoreCountdown(ctx context.Context, e spin.Env) {
	sniperScore := 20_000_000

	modeText := [3]string{"SNIPER", "SHOOT", "SNIPER TOWER"}
	if done := modeIntroVideo(ctx, e, modeText); done {
		return
	}

	sniperScoreFrame(e, sniperScore)
	if done := spin.Wait(ctx, 1000*time.Millisecond); done {
		return
	}

	expires := time.Now().Add(30 * time.Second)
	for time.Now().Before(expires) {
		sniperScore -= 78_330
		sniperScoreFrame(e, sniperScore)
		if done := spin.Wait(ctx, 160*time.Millisecond); done {
			return
		}
	}
	sniperScore = 5_000_000
	sniperScoreFrame(e, sniperScore)
}

func sniperHunt(ctx context.Context, e spin.Env) {
	e.Do(spin.PlayMusic{ID: ModeTheme1, Vol: 100})
	e.Do(spin.PlaySpeech{ID: SniperIsShootingIntoCrowdFromJohnsonTower})
	e.Do(spin.PlayScript{ID: SniperScoreCountdown})

	if done := spin.Wait(ctx, 4*time.Second); done {
		return
	}

	e.Do(spin.PlaySpeech{ID: ShootSniperTower})

	if done := spin.Wait(ctx, 1*time.Second); done {
		return
	}

	for i := 0; i < 8; i++ {
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

func sniperCaught(ctx context.Context, e spin.Env) {
	e.Do(spin.VolumeMusic{Mul: 0.5})
	e.Do(spin.PlaySound{ID: Success})
	if done := spin.Wait(ctx, 1500*time.Millisecond); done {
		return
	}
	e.Do(spin.VolumeMusic{Mul: 2})
	e.Post(spin.Message{ID: SniperCaughtDone})
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
	defer func() {
		e.Do(spin.StopScript{ID: SniperCaught})
		e.Do(spin.StopScript{ID: SniperFall})
		e.Do(spin.StopScript{ID: SniperHunt})
	}()

	e.Do(spin.StopAudio{})
	e.Do(spin.PlayScript{ID: SniperHunt})

	if done, evt := spin.WaitForSwitchUntil(ctx, e, "popper", 30*time.Second); done || evt.ID == "" {
		return
	}
	e.Do(spin.StopScript{ID: SniperHunt})
	e.Do(spin.PlayScript{ID: SniperCaught})
	if done, _ := spin.WaitForMessage(ctx, e, SniperCaughtDone); done {
		return
	}
	e.Do(spin.PlayScript{ID: SniperFall})
	if done, evt := spin.WaitForSwitchUntil(ctx, e, "popper", 16*time.Second); done || evt.ID == "" {
		return
	}
}
