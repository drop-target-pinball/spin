package spin

import "github.com/drop-target-pinball/coroutine"

type Config struct {
	CoilTrough string

	SwitchDrain                 string
	SwitchEnterServiceButton    string
	SwitchExitServiceButton     string
	SwitchLeftFlipperButton     string
	SwitchNextServiceButton     string
	SwitchPreviousServiceButton string
	SwitchRightFlipperButton    string
	SwitchShooterLane           string
	SwitchStartButton           string
	SwitchTrough                []string
	SwitchTroughJam             string
	SwitchWillDrain             []string
	PlayfieldSwitches           []coroutine.Event
	LampStartButton             string

	GI []string

	NumBalls int
}

type Options struct {
	RegisterEOS bool
}

func DefaultOptions() Options {
	return Options{}
}
