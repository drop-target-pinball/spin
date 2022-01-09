package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptAttractMode                   = "jdx.ScriptAttractMode"
	ScriptBadImpersonatorComplete       = "jdx.ScriptBadImpersonatorComplete"
	ScriptBadImpersonatorCountdown      = "jdx.ScriptBadImpersonatorCountdown"
	ScriptBadImpersonatorCountdownAudio = "jdx.ScriptBadImpersonatorCountdownAudio"
	ScriptBadImpersonatorHit            = "jdx.ScriptBadImpersonatorHit"
	ScriptBadImpersonatorMode           = "jdx.ScriptBadImpersonatorMode"
	ScriptBall                          = "jdx.ScriptBall"
	ScriptBasicMode                     = "jdx.ScriptBasicMode"
	ScriptBlackoutJackpot               = "jdx.ScriptBlackoutJackpot"
	ScriptBlackoutMode                  = "jdx.ScriptBlackoutMode"
	ScriptBonusMode                     = "jdx.ScriptBonusMode"
	ScriptChain                         = "jdx.ScriptChain"
	ScriptDefaultLeftShooterLane        = "jdx.ScriptDefaultLeftShooterLane"
	ScriptDefaultLeftPopper             = "jdx.ScriptDefaultLeftPopper"
	ScriptDefaultRightPopper            = "jdx.ScriptDefaultRightPopper"
	ScriptDebugExtraBall                = "jdx.ScriptDebugExtraBall"
	ScriptDemo                          = "jdx.ScriptDemo"
	ScriptGame                          = "jdx.ScriptGame"
	ScriptManhuntComplete               = "jdx.ScriptManhuntComplete"
	ScriptManhuntMode                   = "jdx.ScriptManhuntMode"
	ScriptMeltdownComplete              = "jdx.ScriptMeltdownComplete"
	ScriptMeltdownCountdown             = "jdx.ScriptMeltdownCountdown"
	ScriptMeltdownMode                  = "jdx.ScriptMeltdownMode"
	ScriptMeltdownTimeout               = "jdx.ScriptMeltdownTimeout"
	ScriptMatchMode                     = "jdx.ScriptMatchMode"
	ScriptOutlane                       = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce                = "jdx.ScriptPlayerAnnounce"
	ScriptPlungeMode                    = "jdx.ScriptPlungeMode"
	ScriptProgram                       = "jdx.ScriptProgram"
	ScriptPursuitComplete               = "jdx.ScriptPursuitComplete"
	ScriptPursuitIncomplete             = "jdx.ScriptPursuitIncomplete"
	ScriptPursuitMode                   = "jdx.ScriptPursuitMode"
	ScriptReturnLane                    = "jdx.ScriptReturnLane"
	ScriptSafecrackerComplete           = "jdx.ScriptSafecrackerComplete"
	ScriptSafecrackerCountdown1         = "jdx.ScriptSafecrackerCountdown1"
	ScriptSafecrackerIncomplete         = "jdx.ScriptSafecrackerIncomplete"
	ScriptSafecrackerMode               = "jdx.ScriptSafecrackerMode"
	ScriptSafecrackerOpenThatSafe       = "jdx.ScriptSafecrackerOpenThatSafe"
	ScriptSling                         = "jdx.ScriptSling"
	ScriptSniperComplete                = "jdx.ScriptSniperComplete"
	ScriptSniperIncomplete              = "jdx.ScriptSniperIncomplete"
	ScriptSniperMode                    = "jdx.ScriptSniperMode"
	ScriptSniperMode2                   = "jdx.ScriptSniperMode2"
	ScriptStakeoutComplete              = "jdx.ScriptStakeoutComplete"
	ScriptStakeoutMode                  = "jdx.ScriptStakeoutMode"
	ScriptStakeoutInteresting           = "jdx.ScriptStakeoutInteresting"
	ScriptTankHit                       = "jdx.ScriptTankHit"
	ScriptTankComplete                  = "jdx.ScriptTakComplete"
	ScriptTankIncomplete                = "jdx.ScriptTankIncomplete"
	ScriptTankMode                      = "jdx.ScriptTankMode"
)

func RegisterScripts(eng *spin.Engine) {
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptAttractMode,
	// 	Script: attractModeScript,
	// 	Scope:  spin.ScopeProgram,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBadImpersonatorComplete,
	// 	Script: impersonatorCompleteScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBadImpersonatorCountdown,
	// 	Script: impersonatorCountdownScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBadImpersonatorCountdownAudio,
	// 	Script: impersonatorCountdownAudioScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBadImpersonatorHit,
	// 	Script: impersonatorHitScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBadImpersonatorMode,
	// 	Script: impersonatorModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBall,
	// 	Script: ballScript,
	// 	Scope:  spin.ScopeBall,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBasicMode,
	// 	Script: basicModeScript,
	// 	Scope:  spin.ScopeBall,
	// })
	eng.Do(spin.RegisterScript{
		ID:     ScriptBlackoutJackpot,
		Script: blackoutJackpotScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBlackoutMode,
		Script: blackoutModeScript,
	})
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptBonusMode,
	// 	Script: bonusModeScript,
	// 	Scope:  spin.ScopeGame,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptChain,
	// 	Script: chainScript,
	// 	Scope:  spin.ScopeBall,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDefaultLeftShooterLane,
	// 	Script: defaultLeftShooterLaneScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDefaultLeftPopper,
	// 	Script: defaultLeftPopperScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDefaultRightPopper,
	// 	Script: defaultRightPopperScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDebugExtraBall,
	// 	Script: debugExtraBallScript,
	// 	Scope:  spin.ScopeBall,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptDemo,
	// 	Script: demoScript,
	// 	Scope:  spin.ScopeInit,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptGame,
	// 	Script: gameScript,
	// 	Scope:  spin.ScopeGame,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptManhuntComplete,
	// 	Script: manhuntCompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptManhuntMode,
	// 	Script: manhuntModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptMeltdownComplete,
	// 	Script: meltdownCompleteScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptMeltdownCountdown,
	// 	Script: meltdownCountdownScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptMeltdownMode,
	// 	Script: meltdownModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptMeltdownTimeout,
	// 	Script: meltdownTimeoutScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptMatchMode,
	// 	Script: matchModeScript,
	// 	Scope:  spin.ScopeGame,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptOutlane,
	// 	Script: outlaneScript,
	// 	Scope:  spin.ScopeBall,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptPlayerAnnounce,
	// 	Script: playerAnnounceScript,
	// 	Scope:  spin.ScopeGame,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptPlungeMode,
	// 	Script: plungeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptProgram,
	// 	Script: programScript,
	// 	Scope:  spin.ScopeProgram,
	// })
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitComplete,
		Script: pursuitCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitIncomplete,
		Script: pursuitIncompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPursuitMode,
		Script: pursuitModeScript,
	})
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptReturnLane,
	// 	Script: returnLaneScript,
	// 	Scope:  spin.ScopeBall,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSafecrackerComplete,
	// 	Script: safecrackerCompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSafecrackerCountdown1,
	// 	Script: safecrackerCountdown1Script,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSafecrackerIncomplete,
	// 	Script: safecrackerIncompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSafecrackerMode,
	// 	Script: safecrackerModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSafecrackerOpenThatSafe,
	// 	Script: safecrackerOpenThatSafeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptSling,
	// 	Script: slingScript,
	// 	Scope:  spin.ScopeBall,
	// })
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperComplete,
		Script: sniperCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperIncomplete,
		Script: sniperIncompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode,
		Script: sniperModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSniperMode2,
		Script: sniperMode2Script,
	})
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptStakeoutComplete,
	// 	Script: stakeoutCompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptStakeoutMode,
	// 	Script: stakeoutModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptStakeoutInteresting,
	// 	Script: stakeoutInterestingScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptTankComplete,
	// 	Script: tankCompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptTankHit,
	// 	Script: tankHitScript,
	// 	Scope:  spin.ScopeMode,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptTankIncomplete,
	// 	Script: tankIncompleteScript,
	// 	Scope:  spin.ScopePriority,
	// })
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptTankMode,
	// 	Script: tankModeScript,
	// 	Scope:  spin.ScopeMode,
	// })
}
