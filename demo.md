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



