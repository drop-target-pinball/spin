package msg

const (
	ChanSys = "sys"
)

const (
	PingName = "ping"
	PongName = "pong"
)

type Ping struct {
	From string `json:"from" opt:"short=f"`
	ID   string `json:"id" opt:"short=i"`
}

func (m Ping) Name() string { return PingName }
func (m Ping) Chan() string { return ChanSys }

type Pong struct {
	From string `json:"from" opt:"short=f"`
	ID   string `json:"id" opt:"short=i"`
}

func (m Pong) Name() string { return PongName }
func (m Pong) Chan() string { return ChanSys }
