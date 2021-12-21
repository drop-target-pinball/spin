package jdx

import (
	"time"

	"github.com/drop-target-pinball/spin"
	"github.com/drop-target-pinball/spin/prog/builtin"
)

/*
SetVar Vars=jdx.0 ID=CrimeScenes Val=4
SetVar Vars=jdx.0 ID=SniperBonus Val=40000000
PlayScript ID=jdx.ScriptBonusMode
*/

func bonusFrame(e spin.Env, header string, score int) {
	r, g := e.Display("").Renderer("")

	r.Fill(spin.ColorBlack)
	g.Font = builtin.FontPfArmaFive8
	g.Y = 4
	r.Print(g, header)
	g.Font = builtin.Font14x10
	g.Y = 14
	r.Print(g, spin.FormatScore("%d", score))
}

func bonusModeScript(e spin.Env) {
	vars := GetVars(e)

	bonuses := []struct {
		name  string
		score int
	}{
		{"CRIME SCENES", vars.CrimeScenes * ScoreCrimeSceneBonus},
		{"SNIPER", vars.SniperBonus},
	}

	totalBonus := 0
	for _, bonus := range bonuses {
		if bonus.score > 0 {
			bonusFrame(e, bonus.name, bonus.score)
			e.Do(spin.PlaySound{ID: SoundBonus})
			totalBonus += bonus.score
			if done := e.Sleep(1300 * time.Millisecond); done {
				return
			}
		}
	}
	if totalBonus == 0 {
		totalBonus = ScoreMinimumBonus
	}
	bonusFrame(e, "TOTAL BONUS", totalBonus)
	e.Do(spin.PlaySound{ID: SoundBonus})
	e.Do(spin.AwardScore{Val: totalBonus})

	if done := e.Sleep(2 * time.Second); done {
		return
	}
	e.Post(spin.ScriptFinishedEvent{ID: ScriptBonusMode})
}
