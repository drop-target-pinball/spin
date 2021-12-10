package spin

type Config struct {
	SwitchEnterServiceButton    string
	SwitchExitServiceButton     string
	SwitchLeftFlipperButton     string
	SwitchNextServiceButton     string
	SwitchPreviousServiceButton string
	SwitchRightFlipperButton    string
	SwitchStartButton           string

	LampStartButton string

	GI []string
}

type Options struct {
	RegisterEOS bool
}

func DefaultOptions() Options {
	return Options{}
}
