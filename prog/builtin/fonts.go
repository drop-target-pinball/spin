package builtin

import "github.com/drop-target-pinball/spin"

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
)

func RegisterFonts(eng *spin.Engine) {
	eng.Do(spin.RegisterFont{
		ID:   Font04B_03_7px,
		Path: "proc-shared/04B-03-7px.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font09x5,
		Path: "proc-shared/Font09x5.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font09x6,
		Path: "proc-shared/Font09x6.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font09x7,
		Path: "proc-shared/Font09x7.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font14x8,
		Path: "proc-shared/Font14x8.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font14x9,
		Path: "proc-shared/Font14x9.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font14x10,
		Path: "proc-shared/Font14x10.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font18x10,
		Path: "proc-shared/Font18x10.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font18x11,
		Path: "proc-shared/Font18x11.dmd",
	})
	eng.Do(spin.RegisterFont{
		ID:   Font18x12,
		Path: "proc-shared/Font18x12.dmd",
	})
}
