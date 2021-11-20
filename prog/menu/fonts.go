package menu

import "github.com/drop-target-pinball/spin"

const (
	FontPfArmaFive8                   = "FontPfArmaFive8"
	FontPfRondaSevenBold8             = "FontPfRondaSevenBold8"
	FontPfTempestaFiveExtendedBold8   = "FontPfTempestaFiveExtendedBold8"
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
		ID:   FontPfTempestaFiveExtendedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_extended_bold.ttf",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontPfTempestaFiveCompressedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed_bold.ttf",
	})
}
