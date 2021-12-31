package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

const (
	ScriptAttractMode            = "jdx.ScriptAttractMode"
	ScriptBall                   = "jdx.ScriptBall"
	ScriptBasicMode              = "jdx.ScriptBasicMode"
	ScriptBlackoutMode           = "jdx.ScriptBlackoutMode"
	ScriptBonusMode              = "jdx.ScriptBonusMode"
	ScriptChain                  = "jdx.ScriptChain"
	ScriptDefaultLeftShooterLane = "jdx.ScriptDefaultLeftShooterLane"
	ScriptDefaultLeftPopper      = "jdx.ScriptDefaultLeftPopper"
	ScriptDefaultRightPopper     = "jdx.ScriptDefaultRightPopper"
	ScriptDebugExtraBall         = "jdx.ScriptDebugExtraBall"
	ScriptGame                   = "jdx.ScriptGame"
	ScriptMeltdownComplete       = "jdx.ScriptMeltdownComplete"
	ScriptMeltdownCountdown      = "jdx.ScriptMeltdownCountdown"
	ScriptMeltdownMode           = "jdx.ScriptMeltdownMode"
	ScriptMeltdownTimeout        = "jdx.ScriptMeltdownTimeout"
	ScriptMatchMode              = "jdx.ScriptMatchMode"
	ScriptOutlane                = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce         = "jdx.ScriptPlayerAnnounce"
	ScriptPlungeMode             = "jdx.ScriptPlungeMode"
	ScriptProgram                = "jdx.ScriptProgram"
	ScriptPursuitCaught          = "jdx.ScriptPursuitCaught"
	ScriptPursuitCountdown       = "jdx.ScriptPursuitCountdown"
	ScriptPursuitEscape          = "jdx.ScriptPursuitEscape"
	ScriptPursuitMode            = "jdx.ScriptPursuitMode"
	ScriptReturnLane             = "jdx.ScriptReturnLane"
	ScriptSling                  = "jdx.ScriptSling"
	ScriptSniperMode             = "jdx.ScriptSniperMode"
	ScriptSniperScoreCountdown   = "jdx.ScriptSniperScoreCountdown"
	ScriptSniperSplat            = "jdx.ScriptSniperSplat"
	ScriptSniperTakedown         = "jdx.ScriptSniperTakedown"
	ScriptSniperFallCountdown    = "jdx.ScriptSniperFallCountdown"
	ScriptTankCountdown          = "jdx.ScriptTankCountdown"
	ScriptTankDestroyed          = "jdx.ScriptTankDestroyed"
	ScriptTankMode               = "jdx.ScriptTankMode"
)

func defaultLeftShooterLaneScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftShooterLane}); done {
			return
		}
		e.Do(spin.PlayScript{ID: jd.ScriptRaiseDropTargets})
		if done := e.Sleep(1 * time.Second); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftShooterLane})
	}
}

func defaultLeftPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotLeftPopper}); done {
			return
		}
		for i := 0; i < 3; i++ {
			e.Do(spin.DriverPulse{ID: jd.FlasherSubwayExit})
			if done := e.Sleep(250 * time.Millisecond); done {
				return
			}
		}
		e.Do(spin.DriverPulse{ID: jd.CoilLeftPopper})
	}
}

func defaultRightPopperScript(e spin.Env) {
	for {
		if _, done := e.WaitFor(spin.ShotEvent{ID: jd.ShotRightPopper}); done {
			return
		}
		e.Do(spin.DriverPulse{ID: jd.CoilRightPopper})
	}
}

func modeIntroFrame(e spin.Env, blinkOn bool, text [3]string) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Y = 2
	g.Font = builtin.FontPfArmaFive8
	r.Print(g, text[0])
	if blinkOn {
		g.Y = 12
		g.Font = builtin.FontPfRondaSevenBold8
		r.Print(g, text[1])
		g.Y = 22
		r.Print(g, text[2])
	}
}

func modeIntroVideo(e spin.Env, text [3]string) bool {
	for i := 0; i < 8; i++ {
		modeIntroFrame(e, true, text)
		if done := e.Sleep(250 * time.Millisecond); done {
			return done
		}
		modeIntroFrame(e, false, text)
		if done := e.Sleep(100 * time.Millisecond); done {
			return done
		}
	}
	return false
}

func RegisterScripts(eng *spin.Engine) {
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractMode,
		Script: attractModeScript,
		Scope:  spin.ScopeProgram,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBall,
		Script: ballScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBasicMode,
		Script: basicModeScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBlackoutMode,
		Script: blackoutModeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBonusMode,
		Script: bonusModeScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptChain,
		Script: chainScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultLeftShooterLane,
		Script: defaultLeftShooterLaneScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultLeftPopper,
		Script: defaultLeftPopperScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDefaultRightPopper,
		Script: defaultRightPopperScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDebugExtraBall,
		Script: debugExtraBallScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGame,
		Script: gameScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownComplete,
		Script: meltdownCompleteScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownCountdown,
		Script: meltdownCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownMode,
		Script: meltdownModeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownTimeout,
		Script: meltdownTimeoutScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMatchMode,
		Script: matchModeScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptOutlane,
		Script: outlaneScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlayerAnnounce,
		Script: playerAnnounceScript,
		Scope:  spin.ScopeGame,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlungeMode,
		Script: plungeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptProgram,
		Script: programScript,
		Scope:  spin.ScopeProgram,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitCaught,
		Script: pursuitCaughtScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitCountdown,
		Script: pursuitCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitEscape,
		Script: pursuitEscapeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitMode,
		Script: pursuitModeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptReturnLane,
		Script: returnLaneScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSling,
		Script: slingScript,
		Scope:  spin.ScopeBall,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperTakedown,
		Script: sniperTakedownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperScoreCountdown,
		Script: sniperScoreCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperSplat,
		Script: sniperSplatScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperFallCountdown,
		Script: sniperFallCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode,
		Script: sniperModeScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankCountdown,
		Script: tankCountdownScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankDestroyed,
		Script: tankDestroyedScript,
		Scope:  spin.ScopeMode,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankMode,
		Script: tankModeScript,
		Scope:  spin.ScopeMode,
	})
}
