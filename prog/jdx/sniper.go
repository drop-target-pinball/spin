package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
)

const (
	SniperCaught     = "SniperCaught"
	SniperCaughtDone = "SniperCaughtDone"
	SniperFall       = "SniperFall"
	SniperFallDone   = "SniperFallDone"
	SniperHunt       = "SniperHunt"
	SniperMode       = "SniperMode"
)

func sniperHunt(ctx context.Context, e spin.Env) {
	defer func() {
		e.Do(spin.StopSpeech{})
	}()

	e.Do(spin.PlayMusic{ID: ModeTheme1, Vol: 100})
	e.Do(spin.PlaySpeech{ID: SniperIsShootingIntoCrowdFromJohnsonTower})

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
		e.Do(spin.StopAudio{})
		e.Do(spin.StopScript{ID: SniperCaught})
		e.Do(spin.StopScript{ID: SniperFall})
		e.Do(spin.StopScript{ID: SniperHunt})
	}()

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

func sniperDMD(ctx context.Context, e spin.Env) {
	g := e.Display("").Graphics()
	g.FillRect(10, 10)
	if done := spin.Wait(ctx, 1*time.Second); done {
		return
	}
	g.Move(10, 10)
	g.FillRect(10, 10)
}
