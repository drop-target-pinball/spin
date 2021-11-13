package jdx

import (
	"context"
	"time"

	"github.com/drop-target-pinball/spin"
)

const (
	SniperCaught        = "SniperCaught"
	SniperCaughtDone    = "SniperCaughtDone"
	SniperModeIntroShow = "SniperDisplayModeShow"
	SniperFall          = "SniperFall"
	SniperFallDone      = "SniperFallDone"
	SniperHunt          = "SniperHunt"
	SniperMode          = "SniperMode"
)

func sniperModeIntroShow(ctx context.Context, e spin.Env) {
	r := e.Display("").Renderer()

	g := &spin.Graphics{
		Color: 0xffffffff,
		W:     r.Width(),
	}

	r.Clear()
	for i := 0; i < 8; i++ {
		g.Y = 2
		g.Font = PfArmaFive8
		r.Print(g, "SNIPER")
		g.Y = 12
		g.Font = PfRondaSevenBold8
		r.Print(g, "SHOOT")
		g.Y = 22
		r.Print(g, "SNIPER TOWER")

		if done := spin.Wait(ctx, 250*time.Millisecond); done {
			return
		}

		r.Clear()
		g.Y = 2
		g.Font = PfArmaFive8
		r.Print(g, "SNIPER")

		if done := spin.Wait(ctx, 100*time.Millisecond); done {
			return
		}
	}
}

func sniperHunt(ctx context.Context, e spin.Env) {
	defer func() {
		e.Do(spin.StopSpeech{})
	}()

	e.Do(spin.PlayMusic{ID: ModeTheme1, Vol: 100})
	e.Do(spin.PlaySpeech{ID: SniperIsShootingIntoCrowdFromJohnsonTower})
	e.Do(spin.PlayScript{ID: SniperModeIntroShow})

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
