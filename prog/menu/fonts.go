package menu

import "github.com/drop-target-pinball/spin"

const (
	PfArmaFive8       = "PfArmaFive8"
	PfRondaSevenBold8 = "PfRondaSevenBold8"
)

func RegisterFonts(eng *spin.Engine) {
	eng.Do(spin.RegisterFont{
		ID:   PfArmaFive8,
		Size: 8,
		Path: "pf-fonts/pf_arma_five.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   PfRondaSevenBold8,
		Size: 8,
		Path: "pf-fonts/pf_ronda_seven_bold.ttf",
	})
}
