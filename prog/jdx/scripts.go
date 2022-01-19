package jdx

import (
	"github.com/drop-target-pinball/spin"
)

const (
	ScriptAttractMode             = "jdx.ScriptAttractMode"
	ScriptAttractModeSlide        = "jdx.ScriptAttractModeSlide"
	ScriptBadImpersonatorComplete = "jdx.ScriptBadImpersonatorComplete"
	ScriptBadImpersonatorCrowd    = "jdx.ScriptBadImpersonatorCrowd"
	ScriptBadImpersonatorHit      = "jdx.ScriptBadImpersonatorHit"
	ScriptBadImpersonatorMode     = "jdx.ScriptBadImpersonatorMode"
	ScriptBall                    = "jdx.ScriptBall"
	ScriptBlackoutJackpot         = "jdx.ScriptBlackoutJackpot"
	ScriptBlackoutMode            = "jdx.ScriptBlackoutMode"
	ScriptBonusMode               = "jdx.ScriptBonusMode"
	ScriptChain                   = "jdx.ScriptChain"
	ScriptDemo                    = "jdx.ScriptDemo"
	ScriptGame                    = "jdx.ScriptGame"
	ScriptManhuntComplete         = "jdx.ScriptManhuntComplete"
	ScriptManhuntMode             = "jdx.ScriptManhuntMode"
	ScriptMeltdownComplete        = "jdx.ScriptMeltdownComplete"
	ScriptMeltdownIncomplete      = "jdx.ScriptMeltdownIncomplete"
	ScriptMeltdownMode            = "jdx.ScriptMeltdownMode"
	ScriptMatchMode               = "jdx.ScriptMatchMode"
	ScriptPlungeMode              = "jdx.ScriptPlungeMode"
	ScriptProgram                 = "jdx.ScriptProgram"
	ScriptPursuitComplete         = "jdx.ScriptPursuitComplete"
	ScriptPursuitIncomplete       = "jdx.ScriptPursuitIncomplete"
	ScriptPursuitMode             = "jdx.ScriptPursuitMode"
	ScriptSafecrackerComplete     = "jdx.ScriptSafecrackerComplete"
	ScriptSafecrackerIncomplete   = "jdx.ScriptSafecrackerIncomplete"
	ScriptSafecrackerMode         = "jdx.ScriptSafecrackerMode"
	ScriptSafecrackerMode1        = "jdx.ScriptSafecrackerMode1"
	ScriptSafecrackerMode2        = "jdx.ScriptSafecrackerMode2"
	ScriptSafecrackerOpenThatSafe = "jdx.ScriptSafecrackerOpenThatSafe"
	ScriptSniperComplete          = "jdx.ScriptSniperComplete"
	ScriptSniperIncomplete        = "jdx.ScriptSniperIncomplete"
	ScriptSniperMode              = "jdx.ScriptSniperMode"
	ScriptSniperMode1             = "jdx.ScriptSniperMode1"
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractMode,
		Script: attractModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptAttractModeSlide,
		Script: attractModeSlideScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptBall,
		Script: ballScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBlackoutJackpot,
		Script: blackoutJackpotScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBlackoutMode,
		Script: blackoutModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptBonusMode,
		Script: bonusModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptChain,
		Script: chainScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptDemo,
		Script: demoScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptGame,
		Script: gameScript,
	})
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
	eng.Do(spin.RegisterScript{
		ID:     ScriptMatchMode,
		Script: matchModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptPlungeMode,
		Script: plungeModeScript,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptProgram,
		Script: programScript,
	})
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
		ID:     ScriptSafecrackerMode1,
		Script: safecrackerMode1Script,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerMode2,
		Script: safecrackerMode2Script,
	})
	eng.Do(spin.RegisterScript{
		ID:     ScriptSafecrackerOpenThatSafe,
		Script: safecrackerOpenThatSafeScript,
	})
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
		ID:     ScriptSniperMode1,
		Script: sniperMode1Script,
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
