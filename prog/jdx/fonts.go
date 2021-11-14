package jdx

import "github.com/drop-target-pinball/spin"

const (
	Bm8                           = "Bm8"
	Bmsf                          = "Bmsf"
	PfArmaFive8                   = "PfArmaFive8"
	PfRondaSevenBold8             = "PfRondaSevenBold8"
	PfTempestaFiveCompressedBold8 = "PfTempestaFiveCompressedBold8"
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
	eng.Do(spin.RegisterFont{
		ID:   PfTempestaFiveCompressedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed_bold.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   Bm8,
		Path: "proc-shared/bm8.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Bmsf,
		Path: "proc-shared/bmsf.dmd",
	})
}
