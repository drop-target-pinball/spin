package game

import "github.com/drop-target-pinball/spin"

const (
	FontBm3   = "FontBm3"
	FontBm8   = "FontBm8"
	FontBm10w = "FontBm10w"
	FontBm10  = "FontBm10"
	FontBmsf  = "FontBmsf"
)

func RegisterFonts(eng *spin.Engine) {
	eng.Do(spin.RegisterFont{
		ID:   FontBm3,
		Path: "proc-shared/bm3.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBm8,
		Path: "proc-shared/bm8.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBm10,
		Path: "proc-shared/bm10.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBm10w,
		Path: "proc-shared/bm10w.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   FontBmsf,
		Path: "proc-shared/bmsf.dmd",
	})
}
