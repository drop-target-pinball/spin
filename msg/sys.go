package msg

const ChanSys = "sys"

const (
	AlertName = "alert"
	LoadName  = "load"
	PingName  = "ping"
	PongName  = "pong"
	ReadyName = "ready"
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

type Load struct {
	ID string `json:"id"`
}

func (m Load) Name() string { return LoadName }
func (m Load) Chan() string { return ChanSys }

type Ready struct {
	ID string `json:"id"`
}

func (m Ready) Name() string { return ReadyName }
func (m Ready) Chan() string { return ChanSys }
