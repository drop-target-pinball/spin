package jd

import "github.com/drop-target-pinball/spin"

var PlayerNavEvents = []spin.Event{
	spin.SwitchEvent{ID: LeftFlipperButton},
	spin.SwitchEvent{ID: RightFlipperButton},
	spin.SwitchEvent{ID: LeftFireButton},
	spin.SwitchEvent{ID: RightFireButton},
}

func Load(eng *spin.Engine) {
	RegisterKeys(eng)
}
