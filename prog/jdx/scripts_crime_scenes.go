package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/mach/jd"
	"github.com/drop-target-pinball/spin/proc"
)

var crimeCallouts = map[int]map[int]string{
	jd.CrimeSceneInnerLoop: {
		jd.CrimeLevelMisdemeanor:  SpeechFifteenYearsForExhibitionism,
		jd.CrimeLevelFelony:       SpeechTwentyFiveYearsForAutoTheft,
		jd.CrimeLevelClassXFelony: SpeechLifeForTerrorism,
	},
	jd.CrimeSceneLeftLoop: {
		jd.CrimeLevelWarning:      SpeechWarningForJaywalking,
		jd.CrimeLevelMisdemeanor:  SpeechOneYearForFlatulism,
		jd.CrimeLevelFelony:       SpeechFifteenYearsForArson,
		jd.CrimeLevelClassXFelony: SpeechThirtyFiveYearsForArmsDealing,
	},
	jd.CrimeSceneRightLoop: {
		jd.CrimeLevelWarning:      SpeechWarningForLittering,
		jd.CrimeLevelMisdemeanor:  SpeechTwoYearsForSpeeding,
		jd.CrimeLevelFelony:       SpeechEighteenYearsForAssault,
		jd.CrimeLevelClassXFelony: SpeechEightyFiveYearsForTreason,
	},
	jd.CrimeSceneRightPopper: {
		jd.CrimeLevelWarning:      SpeechWarningForSmoking,
		jd.CrimeLevelMisdemeanor:  SpeechThreeYearsForVandalism,
		jd.CrimeLevelFelony:       SpeechTwentyYearsForStalking,
		jd.CrimeLevelClassXFelony: SpeechNinetyYearsForKidnapping,
	},
	jd.CrimeSceneRightRamp: {
		jd.CrimeLevelWarning:      SpeechWarningForSpitting,
		jd.CrimeLevelMisdemeanor:  SpeechTenYearsForSolicitation,
		jd.CrimeLevelFelony:       SpeechTwentyThreeYearsForExtortion,
		jd.CrimeLevelClassXFelony: SpeechLifeForCannibalism,
	},
}

var crimeText = map[int]map[int][]string{
	jd.CrimeSceneInnerLoop: {
		jd.CrimeLevelMisdemeanor:  []string{"15 YEARS FOR", "EXHIBITIONISM"},
		jd.CrimeLevelFelony:       []string{"25 YEARS FOR", "AUTO THEFT"},
		jd.CrimeLevelClassXFelony: []string{"LIFE FOR", "TERRORISM"},
	},
	jd.CrimeSceneLeftLoop: {
		jd.CrimeLevelWarning:      []string{"WARNING FOR", "JAYWALKING"},
		jd.CrimeLevelMisdemeanor:  []string{"1 YEAR FOR", "FLATULISM"},
		jd.CrimeLevelFelony:       []string{"15 YEARS FOR", "ARSON"},
		jd.CrimeLevelClassXFelony: []string{"35 YEARS FOR", "ARMS DEALING"},
	},
	jd.CrimeSceneRightLoop: {
		jd.CrimeLevelWarning:      []string{"WARNING FOR", "LITTERING"},
		jd.CrimeLevelMisdemeanor:  []string{"2 YEARS FOR", "SPEEDING"},
		jd.CrimeLevelFelony:       []string{"18 YEARS FOR", "ASSAULT"},
		jd.CrimeLevelClassXFelony: []string{"80 YEARS FOR", "TREASON"},
	},
	jd.CrimeSceneRightPopper: {
		jd.CrimeLevelWarning:      []string{"WARNING FOR", "SMOKING"},
		jd.CrimeLevelMisdemeanor:  []string{"3 YEARS FOR", "VANDALISM"},
		jd.CrimeLevelFelony:       []string{"20 YEARS FOR", "STALKING"},
		jd.CrimeLevelClassXFelony: []string{"90 YEARS FOR", "KIDNAPPING"},
	},
	jd.CrimeSceneRightRamp: {
		jd.CrimeLevelWarning:      []string{"WARNING FOR", "SPITTING"},
		jd.CrimeLevelMisdemeanor:  []string{"10 YEARS FOR", "SOLICITATION"},
		jd.CrimeLevelFelony:       []string{"23 YEARS FOR", "EXTORTION"},
		jd.CrimeLevelClassXFelony: []string{"LIFE FOR", "CANNIBALISM"},
	},
}

/*
SwitchEvent ID=jd.SwitchRightFireButton
SwitchEvent ID=jd.SwitchOuterLoopLeft
SwitchEvent ID=jd.SwitchOuterLoopRight
SwitchEvent ID=jd.SwitchInnerLoop
SwitchEvent ID=jd.SwitchRightPopper
SwitchEvent ID=jd.SwitchRightRampExit
SwitchEvent ID=jd.SwitchBankTargets
*/

func crimeScenesScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	if vars.CrimeLevel == jd.CrimeLevelNone {
		vars.CrimeLevel = jd.CrimeLevelWarning
		vars.CrimeScenesLit = jd.CrimeSceneLeftLoop | jd.CrimeSceneRightLoop | jd.CrimeSceneRightPopper | jd.CrimeSceneRightRamp
	}

	if vars.AdvanceCrimeSceneLit {
		e.Do(proc.DriverSchedule{ID: jd.LampAdvanceCrimeLevel})
		e.Do(proc.DriverSchedule{ID: jd.CrimeLevelLamps[vars.CrimeLevel]})
		return
	} else {
		e.Do(spin.DriverOn{ID: jd.CrimeLevelLamps[vars.CrimeLevel]})
		for _, cs := range jd.CrimeScenes {
			if vars.CrimeScenesLit&cs != 0 {
				e.Do(proc.DriverSchedule{ID: jd.CrimeSceneLamps[cs][vars.CrimeLevel], Schedule: proc.BlinkSchedule})
			}
		}
	}

	defer func() {
		if vars.AdvanceCrimeSceneLit {
			e.Do(spin.DriverOff{ID: jd.LampAdvanceCrimeLevel})
		}
		if vars.CrimeLevel != jd.CrimeLevelNone {
			e.Do(spin.DriverOff{ID: jd.CrimeLevelLamps[vars.CrimeLevel]})
		}
		for _, cs := range jd.CrimeScenes {
			if vars.CrimeScenesLit&cs != 0 {
				e.Do(spin.DriverOff{ID: jd.CrimeSceneLamps[cs][vars.CrimeLevel]})
			}
		}
	}()

	var lastCollect time.Time

	for {
		if vars.AdvanceCrimeSceneLit {
			if _, done := e.WaitFor(spin.SwitchEvent{ID: jd.SwitchBankTargets}); done {
				return
			}
			vars.AdvanceCrimeSceneLit = false
			vars.CrimeLevelLast = vars.CrimeLevel
			e.Do(spin.PlayScript{ID: ScriptCrimeLevelAdvance})
			e.Do(spin.AwardScore{Val: CrimeLevelScores[vars.CrimeLevel] * vars.Multiplier})
			e.Do(spin.DriverOff{ID: jd.CrimeLevelLamps[vars.CrimeLevel]})
			e.Do(spin.DriverOff{ID: jd.LampAdvanceCrimeLevel})

			vars.CrimeLevel += 1
			if vars.CrimeLevel > jd.CrimeLevelClassXFelony {
				vars.CrimeLevel = jd.CrimeLevelNone
				return
			}
			vars.CrimeScenesLit = jd.CrimeSceneLeftLoop | jd.CrimeSceneRightLoop | jd.CrimeSceneRightPopper | jd.CrimeSceneRightRamp | jd.CrimeSceneInnerLoop
			for _, cs := range jd.CrimeScenes {
				e.Do(proc.DriverSchedule{ID: jd.CrimeSceneLamps[cs][vars.CrimeLevel], Schedule: proc.BlinkSchedule})
			}
		} else {
			for {
				evt, done := e.WaitFor(jd.CrimeSceneSwitchEvents...)
				if done {
					return
				}
				if time.Since(lastCollect) < 1*time.Second {
					continue
				}
				sw := evt.(spin.SwitchEvent).ID
				scene := jd.CrimeSceneSwitches[sw]
				if vars.CrimeScenesLit&scene != 0 {
					vars.CrimeSceneLastCollected = scene
					vars.CrimeScenesLit &^= scene
					e.Do(spin.DriverOff{ID: jd.CrimeSceneLamps[scene][vars.CrimeLevel]})
					e.Do(spin.AwardScore{Val: CrimeLevelScores[vars.CrimeLevel] * vars.Multiplier})
					e.Do(spin.PlayScript{ID: ScriptCrimeSceneCollect})

					if vars.CrimeScenesLit == 0 {
						vars.AdvanceCrimeSceneLit = true
						e.Do(proc.DriverSchedule{ID: jd.LampAdvanceCrimeLevel, Schedule: proc.BlinkSchedule})
						e.Do(proc.DriverSchedule{ID: jd.CrimeLevelLamps[vars.CrimeLevel], Schedule: proc.BlinkSchedule})
						break
					}
				}
			}
		}
	}
}

func collectCrimeScenePanel(e *spin.ScriptEnv, r spin.Renderer, text []string, score int) {
	g := r.Graphics()

	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, text[0])

	g.Y = 12
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, text[1])

	g.Y = 24
	g.Font = spin.FontPfArmaFive8
	r.Print(g, spin.FormatScore("%v", score))
}

/*
SetVar Vars=jdx.1 ID=CrimeSceneLastCollected Val=1
SetVar Vars=jdx.1 ID=CrimeLevel Val=1
PlayScript ID=jdx.ScriptCrimeSceneCollect
*/
func crimeSceneCollectScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	callout := crimeCallouts[vars.CrimeSceneLastCollected][vars.CrimeLevel]
	text := crimeText[vars.CrimeSceneLastCollected][vars.CrimeLevel]
	score := CrimeLevelScores[vars.CrimeLevel] * vars.Multiplier

	collectCrimeScenePanel(e, r, text, score)

	s := spin.NewSequencer(e)
	s.Do(spin.PlaySpeech{ID: callout})
	s.Sleep(2_000)
	s.Run()
}

func crimeLevelAdvancePanel(e *spin.ScriptEnv, r spin.Renderer, score int) {
	g := r.Graphics()

	g.Y = 2
	g.Font = spin.FontPfArmaFive8
	r.Print(g, "ADVANCE")

	g.Y = 12
	g.Font = spin.FontPfRondaSevenBold8
	r.Print(g, "CRIME LEVEL")

	g.Y = 24
	g.Font = spin.FontPfArmaFive8
	r.Print(g, spin.FormatScore("%v", score))
}

func crimeLevelAdvanceScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(spin.PriorityAnnounce)
	defer r.Close()

	vars := GetVars(e)
	score := CrimeLevelScores[vars.CrimeLevelLast] * vars.Multiplier

	crimeLevelAdvancePanel(e, r, score)

	s := spin.NewSequencer(e)
	s.Do(spin.PlaySpeech{ID: SpeechIfYouDoTheCrimeYouDoTheTime})
	s.Sleep(2_000)
	s.Run()
}
