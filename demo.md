PlayScript ID=jdx.ScriptGame

SwitchEvent ID=jd.SwitchRightFireButton
SwitchEvent ID=jd.SwitchLeftSling

ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotRightRamp
ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotRightRamp

SwitchEvent ID=jd.SwitchRightFireButton

SwitchEvent ID=jd.SwitchRightPopper
SwitchEvent ID=jd.SwitchRightPopper
SwitchEvent ID=jd.SwitchRightPopper

ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotTopLeftRamp
SwitchEvent ID=jd.SwitchBankTargets

SetVar Vars=player.1 ID=Score Val=121881320
SwitchEvent ID=jd.SwitchTrough1

SetVar Vars=game ID=Ball Val=3
SwitchEvent ID=jd.SwitchRightFireButton
SwitchEvent ID=jd.SwitchTrough1

## Pursuit

PlayScript ID=jdx.ScriptPursuitMode
ShotEvent ID=jd.ShotRightRamp
ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotRightRamp
SetVar Vars=jdx.1 ID=Timer Val=3

## Sniper Tower

PlayScript ID=jdx.ScriptSniperMode
SwitchEvent ID=jd.SwitchRightPopper
SwitchEvent ID=jd.SwitchRightPopper

## Battle Tank

PlayScript ID=jdx.ScriptTankMode
ShotEvent ID=jd.ShotLeftRamp
ShotEvent ID=jd.ShotTopLeftRamp
SwitchEvent ID=jd.SwitchBankTargets

## Meltdown

PlayScript ID=jdx.ScriptMeltdownMode
SetVar Vars=jdx.1 ID=Timer Val=21
SetVar Vars=jdx.1 ID=Timer Val=11
SetVar Vars=jdx.1 ID=Timer Val=5
SwitchEvent ID=jd.SwitchCaptiveBall1
SwitchEvent ID=jd.SwitchCaptiveBall2
SwitchEvent ID=jd.SwitchCaptiveBall3

## Bad Impersonator

PlayScript ID=jdx.ScriptBadImpersonatorMode
SwitchEvent ID=jd.SwitchDropTargetE

## Safecracker

PlayScript ID=jdx.ScriptSafecrackerMode
SwitchEvent ID=jd.SwitchSubwayEnter1

## Manhunt

PlayScript ID=jdx.ScriptManhuntMode
ShotEvent ID=jd.ShotLeftRamp

# Stakeout

PlayScript ID=jdx.ScriptStakeoutMode
ShotEvent ID=jd.ShotRightRamp

# Blackout
PlayScript ID=jdx.ScriptBlackoutMode
ShotEvent ID=jd.ShotTopLeftRamp
BallDrainEvent BallsInPlay=1

## Switch Check

SwitchEvent ID=jd.SwitchRightFireButton
SwitchEvent ID=jd.SwitchBankTargets

SwitchEvent ID=jd.SwitchLeftSling
SwitchEvent ID=jd.SwitchRightSling
SwitchEvent ID=jd.SwitchLeftOutlane
SwitchEvent ID=jd.SwitchRightOutlane
