package jdx

import "github.com/drop-target-pinball/spin"

const (
	FontBm8                           = "FontBm8"
	FontBmsf                          = "FontBmsf"
	FontPfArmaFive8                   = "FontPfArmaFive8"
	FontPfRondaSevenBold8             = "FontPfRondaSevenBold8"
	FontPfTempestaFiveCompressedBold8 = "FontPfTempestaFiveCompressedBold8"
)

func RegisterFonts(eng *spin.Engine) {
	eng.Do(spin.RegisterFont{
		ID:   FontPfArmaFive8,
		Size: 8,
		Path: "pf-fonts/pf_arma_five.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontPfRondaSevenBold8,
		Size: 8,
		Path: "pf-fonts/pf_ronda_seven_bold.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontPfTempestaFiveCompressedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed_bold.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBm8,
		Path: "proc-shared/bm8.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBmsf,
		Path: "proc-shared/bmsf.dmd",
	})
}
