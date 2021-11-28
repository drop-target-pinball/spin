package builtin

import "github.com/drop-target-pinball/spin"

const (
	// FontBm3   = "FontBm3"
	// FontBm8   = "FontBm8"
	// FontBm10w = "FontBm10w"
	// FontBm10  = "FontBm10"
	// FontBmsf = "FontBmsf"

	Font04B_03_7px = "04B-03-7px"
)

func RegisterFonts(eng *spin.Engine) {
	// eng.Do(spin.RegisterFont{
	// 	ID:   FontBm3,
	// 	Path: "proc-shared/bm3.dmd",
	// })
	// eng.Do(spin.RegisterFont{
	// 	ID:   FontBm8,
	// 	Path: "proc-shared/bm8.dmd",
	// })
	// eng.Do(spin.RegisterFont{
	// 	ID:   FontBm10,
	// 	Path: "proc-shared/bm10.dmd",
	// })
	// eng.Do(spin.RegisterFont{
	// 	ID:   FontBm10w,
	// 	Path: "proc-shared/bm10w.dmd",
	// })
	// eng.Do(spin.RegisterFont{
	// 	ID:   FontBmsf,
	// 	Path: "proc-shared/bmsf.dmd",
	// })
	eng.Do(spin.RegisterFont{
		ID:   Font04B_03_7px,
		Path: "proc-shared/04B-03-7px.dmd",
	})
}
