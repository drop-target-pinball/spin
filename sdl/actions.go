package sdl

import "github.com/drop-target-pinball/spin"

type RegisterKey struct {
	spin.Action
	Key       string
	Mod       string
	EventDown spin.Event
}

type RegisterButton struct {
	spin.Action
	Button     string
	ActionDown spin.Action
}
