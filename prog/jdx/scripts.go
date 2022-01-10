package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptAttractMode             = "jdx.ScriptAttractMode"
	ScriptBadImpersonatorComplete = "jdx.ScriptBadImpersonatorComplete"
	ScriptBadImpersonatorCrowd    = "jdx.ScriptBadImpersonatorCrowd"
	ScriptBadImpersonatorHit      = "jdx.ScriptBadImpersonatorHit"
	ScriptBadImpersonatorMode     = "jdx.ScriptBadImpersonatorMode"
	ScriptBall                    = "jdx.ScriptBall"
	ScriptBasicMode               = "jdx.ScriptBasicMode"
	ScriptBlackoutJackpot         = "jdx.ScriptBlackoutJackpot"
	ScriptBlackoutMode            = "jdx.ScriptBlackoutMode"
	ScriptBonusMode               = "jdx.ScriptBonusMode"
	ScriptChain                   = "jdx.ScriptChain"
	ScriptDefaultLeftShooterLane  = "jdx.ScriptDefaultLeftShooterLane"
	ScriptDefaultLeftPopper       = "jdx.ScriptDefaultLeftPopper"
	ScriptDefaultRightPopper      = "jdx.ScriptDefaultRightPopper"
	ScriptDebugExtraBall          = "jdx.ScriptDebugExtraBall"
	ScriptDemo                    = "jdx.ScriptDemo"
	ScriptGame                    = "jdx.ScriptGame"
	ScriptManhuntComplete         = "jdx.ScriptManhuntComplete"
	ScriptManhuntMode             = "jdx.ScriptManhuntMode"
	ScriptMeltdownComplete        = "jdx.ScriptMeltdownComplete"
	ScriptMeltdownIncomplete      = "jdx.ScriptMeltdownIncomplete"
	ScriptMeltdownMode            = "jdx.ScriptMeltdownMode"
	ScriptMatchMode               = "jdx.ScriptMatchMode"
	ScriptOutlane                 = "jdx.ScriptOutlane"
	ScriptPlayerAnnounce          = "jdx.ScriptPlayerAnnounce"
	ScriptPlungeMode              = "jdx.ScriptPlungeMode"
	ScriptProgram                 = "jdx.ScriptProgram"
	ScriptPursuitComplete         = "jdx.ScriptPursuitComplete"
	ScriptPursuitIncomplete       = "jdx.ScriptPursuitIncomplete"
	ScriptPursuitMode             = "jdx.ScriptPursuitMode"
	ScriptReturnLane              = "jdx.ScriptReturnLane"
	ScriptSafecrackerComplete     = "jdx.ScriptSafecrackerComplete"
	ScriptSafecrackerIncomplete   = "jdx.ScriptSafecrackerIncomplete"
	ScriptSafecrackerMode         = "jdx.ScriptSafecrackerMode"
	ScriptSafecrackerMode2        = "jdx.ScriptSafecrackerMode2"
	ScriptSafecrackerOpenThatSafe = "jdx.ScriptSafecrackerOpenThatSafe"
	ScriptSling                   = "jdx.ScriptSling"
	ScriptSniperComplete          = "jdx.ScriptSniperComplete"
	ScriptSniperIncomplete        = "jdx.ScriptSniperIncomplete"
	ScriptSniperMode              = "jdx.ScriptSniperMode"
	ScriptSniperMode2             = "jdx.ScriptSniperMode2"
	ScriptStakeoutComplete        = "jdx.ScriptStakeoutComplete"
	ScriptStakeoutMode            = "jdx.ScriptStakeoutMode"
	ScriptStakeoutInteresting     = "jdx.ScriptStakeoutInteresting"
	ScriptTankHit                 = "jdx.ScriptTankHit"
	ScriptTankComplete            = "jdx.ScriptTakComplete"
	ScriptTankIncomplete          = "jdx.ScriptTankIncomplete"
	ScriptTankMode                = "jdx.ScriptTankMode"
	ScriptUseFireButton           = "jdx.ScriptUseFireButton"
)

func RegisterScripts(eng *spin.Engine) {
	// eng.Do(spin.RegisterScript{
	// 	ID:     ScriptAttractMode,
	// 	Script: attractModeScript,
	// 	Scope:  spin.ScopeProgram,
	// })
	eng.Do(spin.RegisterScript{
		ID:     ScriptBadImpersonatorComplete,
		Script: impersonatorCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBadImpersonatorCrowd,
		Script: impersonatorCrowdScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBadImpersonatorHit,
		Script: impersonatorHitScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBadImpersonatorMode,
		Script: impersonatorModeScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptManhuntComplete,
		Script: manhuntCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptManhuntMode,
		Script: manhuntModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownComplete,
		Script: meltdownCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownComplete,
		Script: meltdownCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownIncomplete,
		Script: meltdownIncompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptMeltdownMode,
		Script: meltdownModeScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlungeMode,
		Script: plungeModeScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerComplete,
		Script: safecrackerCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerIncomplete,
		Script: safecrackerIncompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerMode,
		Script: safecrackerModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerMode2,
		Script: safecrackerMode2Script,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerOpenThatSafe,
		Script: safecrackerOpenThatSafeScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptStakeoutComplete,
		Script: stakeoutCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptStakeoutMode,
		Script: stakeoutModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptStakeoutInteresting,
		Script: stakeoutInterestingScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankComplete,
		Script: tankCompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankHit,
		Script: tankHitScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankIncomplete,
		Script: tankIncompleteScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptTankMode,
		Script: tankModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptUseFireButton,
		Script: useFireButtonScript,
	})
}
