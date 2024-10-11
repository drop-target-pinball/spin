package msg

const ChanAudio = "audio"

const (
	PlayMusicName = "play-music"
)

type PlayMusic struct {
	ID     string `json:"id"`
	Vol    int    `json:"vol"`
	Loops  int    `json:"loops,omitempty"`
	Notify bool   `json:"notify,omitempty"`
}

func (m PlayMusic) Name() string { return PlayMusicName }
func (m PlayMusic) Chan() string { return ChanAudio }
