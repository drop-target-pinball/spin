package spin

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
	SwitchTroughJam             string
	SwitchWillDrain             []string

	LampStartButton string

	GI []string

	NumBalls int
}

type Options struct {
	RegisterEOS bool
}

func DefaultOptions() Options {
	return Options{}
}
