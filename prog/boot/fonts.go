package boot

import "github.com/drop-target-pinball/spin"

const (
	FontPfTempestaFiveCompressedBold8 = "FontPfTempestaFiveCompressedBold8"
)

func RegisterFonts(eng *spin.Engine) {
	eng.Do(spin.RegisterFont{
		ID:   FontPfTempestaFiveCompressedBold8,
		Size: 8,
		Path: "pf-fonts/pf_tempesta_five_compressed_bold.ttf",
	})
}
