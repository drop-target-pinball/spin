package msg

const ChanSys = "sys"

const (
	AlertName = "alert"
	PingName  = "ping"
	PongName  = "pong"
)

type Alert struct {
	Desc string `json:"desc"`
}

func (m Alert) Name() string { return AlertName }
func (m Alert) Chan() string { return ChanSys }

type Ping struct {
	From string `json:"from"`
	ID   string `json:"id"`
}

func (m Ping) Name() string { return PingName }
func (m Ping) Chan() string { return ChanSys }

type Pong struct {
	From string `json:"from"`
	ID   string `json:"id"`
}

func (m Pong) Name() string { return PongName }
func (m Pong) Chan() string { return ChanSys }
