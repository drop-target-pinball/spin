package jdx

import (
	"github.com/drop-target-pinball/spin"
)

/*
SetVar Vars=jdx.1 ID=CrimeScenes Val=4
SetVar Vars=jdx.1 ID=SniperBonus Val=40000000
PlayScript ID=jdx.ScriptBonusMode
*/

func bonusModeScript(e *spin.ScriptEnv) {
	vars := GetVars(e)

	bonuses := []struct {
		name  string
		score int
	}{
		{"CRIME SCENES", vars.CrimeScenes * ScoreCrimeSceneBonus},
		{"PURSUIT", vars.PursuitBonus},
		{"SNIPER", vars.SniperBonus},
		{"BATTLE TANK", vars.TankBonus},
		{"BAD IMPERSONATOR", vars.BadImpersonatorBonus},
		{"MELTDOWN", vars.MeltdownBonus},
		{"SAFECRACKER", vars.SafecrackerBonus},
		{"MANHUNT", vars.ManhuntBonus},
		{"STAKEOUT", vars.StakeoutBonus},
	}

	totalBonus := 0
	for _, bonus := range bonuses {
		if bonus.score > 0 {
			bonusPanel(e, bonus.name, bonus.score)
			e.Do(spin.PlaySound{ID: SoundBonus})
			totalBonus += bonus.score
			if done := e.Sleep(1300); done {
				return
			}
		}
	}
	if totalBonus == 0 {
		totalBonus = ScoreMinimumBonus
	}
	bonusPanel(e, "TOTAL BONUS", totalBonus)
	e.Do(spin.PlaySound{ID: SoundBonus})
	e.Do(spin.AwardScore{Val: totalBonus})

	if done := e.Sleep(2000); done {
		return
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBonusMode})
}

func bonusPanel(e *spin.ScriptEnv, header string, score int) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Font = spin.FontPfArmaFive8
	g.Y = 4
	r.Print(g, header)
	g.Font = spin.Font14x10
	g.Y = 14
	r.Print(g, spin.FormatScore("%d", score))
}
