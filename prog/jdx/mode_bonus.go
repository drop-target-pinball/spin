package jdx

import (
	"github.com/drop-target-pinball/spin"
)

/*
SetVar Vars=jdx.1 ID=CrimeScenesCollected Val=4
SetVar Vars=jdx.1 ID=SniperBonus Val=40000000
SetVar Vars=jdx.1 ID=TankBonus Val=33000000
PlayScript ID=jdx.ScriptBonusMode
*/

func bonusModeScript(e *spin.ScriptEnv) {
	r := e.Display("").Open(0)
	defer r.Close()

	vars := GetVars(e)

	bonuses := []struct {
		name  string
		score int
	}{
		{"CRIME SCENES", vars.CrimeScenesCollected * ScoreCrimeSceneBonus},
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
			bonusPanel(e, r, bonus.name, bonus.score)
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
	bonusPanel(e, r, "TOTAL BONUS", totalBonus)
	e.Do(spin.PlaySound{ID: SoundBonus})
	e.Do(spin.AwardScore{Val: totalBonus})

	e.Sleep(2000)
}

func bonusPanel(e *spin.ScriptEnv, r spin.Renderer, header string, score int) {
	g := r.Graphics()

	r.Fill(spin.ColorOff)
	g.Font = spin.FontPfArmaFive8
	g.Y = 4
	r.Print(g, header)
	g.Font = spin.Font14x10
	g.Y = 14
	r.Print(g, spin.FormatScore("%d", score))
}
