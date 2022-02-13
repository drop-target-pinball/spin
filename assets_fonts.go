package spin

const (
	Font04B_03_7px = "04B-03-7px"
	Font09x5       = "Font09x5"
	Font09x6       = "Font09x6"
	Font09x7       = "Font09x7"
	Font14x9       = "Font14x9"
	Font14x8       = "Font14x8"
	Font14x10      = "Font14x10"
	Font18x10      = "Font18x10"
	Font18x11      = "Font18x11"
	Font18x12      = "Font18x12"

	FontPfArmaFive8                   = "FontPfArmaFive8"
	FontPfRondaSeven8                 = "FontPfRondaSeven8"
	FontPfRondaSevenBold8             = "FontPfRondaSevenBold8"
	FontPfRondaSevenBold16            = "FontPfRondaSevenBold16"
	FontPfTempestaFive8               = "FontPfTempestaFive8"
	FontPfTempestaFiveBold8           = "FontPfTempestaFiveBold8"
	FontPfTempestaFiveCompressed8     = "FontPfTempestaFiveCompressed8"
	FontPfTempestaFiveCompressedBold8 = "FontPfTempestaFiveCompressedBold8"
	FontPfTempestaFiveCondensed8      = "FontPfTempestaFiveCondensed8"
	FontPfTempestaFiveCondensedBold8  = "FontPfTempestaFiveCondensedBold8"
	FontPfTempestaFiveExtended8       = "FontPfTempestaFiveExtended8"
	FontPfTempestaFiveExtendedBold8   = "FontPfTempestaFiveExtendedBold8"
)

func RegisterFonts(eng *Engine) {
	eng.Do(RegisterFont{
		ID:   Font04B_03_7px,
		Path: "proc-shared/04B-03-7px.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font09x5,
		Path: "proc-shared/Font09x5.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font09x6,
		Path: "proc-shared/Font09x6.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font09x7,
		Path: "proc-shared/Font09x7.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font14x8,
		Path: "proc-shared/Font14x8.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font14x9,
		Path: "proc-shared/Font14x9.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font14x10,
		Path: "proc-shared/Font14x10.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font18x10,
		Path: "proc-shared/Font18x10.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font18x11,
		Path: "proc-shared/Font18x11.dmd",
	})
	eng.Do(RegisterFont{
		ID:   Font18x12,
		Path: "proc-shared/Font18x12.dmd",
	})

	eng.Do(RegisterFont{
		ID:   FontPfArmaFive8,
		Size: 8,
		Path: "pf-fonts/pf_arma_five.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfRondaSeven8,
		Size: 8,
		Path: "pf-fonts/pf_ronda_seven.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfRondaSevenBold8,
		Size: 8,
		Path: "pf-fonts/pf_ronda_seven_bold.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfRondaSevenBold16,
		Size: 16,
		Path: "pf-fonts/pf_ronda_seven_bold.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFive8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_bold.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveCompressed8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveCompressedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed_bold.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveCondensed8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_condensed.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveCondensedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_condensed_bold.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveExtended8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_extended.ttf",
	})
	eng.Do(RegisterFont{
		ID:   FontPfTempestaFiveExtendedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_extended_bold.ttf",
	})
}
