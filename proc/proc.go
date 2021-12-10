package proc

import (
	"log"

	"github.com/drop-target-pinball/go-pinproc"
	"github.com/drop-target-pinball/go-pinproc/wpc"
	"github.com/drop-target-pinball/spin"
)

type Options struct {
	ID                      string
	MachType                pinproc.MachType
	DMDConfig               pinproc.DMDConfig
	SwitchConfig            pinproc.SwitchConfig
	DefaultCoilPulseTime    int
	DefaultFlasherPulseTime int
}

type procSystem struct {
	eng          *spin.Engine
	proc         *pinproc.PROC
	drivers      map[string]spin.Driver
	switches     map[uint8]spin.Switch
	flippers     map[string]spin.Flipper
	events       []pinproc.Event
	source       spin.Display
	dots         []uint8
	frameSize    int
	subFrameSize int
	opts         Options
}

func RegisterSystem(eng *spin.Engine, opts Options) {
	pc, err := pinproc.New(wpc.MachType)
	if err != nil {
		log.Fatalf("unable to connect to P-ROC: %v", err)
	}

	s := &procSystem{
		eng:      eng,
		proc:     pc,
		drivers:  make(map[string]spin.Driver),
		switches: make(map[uint8]spin.Switch),
		flippers: make(map[string]spin.Flipper),
		events:   make([]pinproc.Event, pinproc.MaxEvents),
		opts:     opts,
	}
	eng.RegisterActionHandler(s)
	eng.RegisterServer(s)

	s.proc.Reset(pinproc.ResetFlagUpdateDevice)
	if err := s.proc.SwitchUpdateConfig(opts.SwitchConfig); err != nil {
		log.Fatal(err)
	}
	if err := s.proc.DMDUpdateConfig(opts.DMDConfig); err != nil {
		log.Fatal(err)
	}
	s.subFrameSize = int(opts.DMDConfig.NumColumns) * int(opts.DMDConfig.NumRows) / 8
	s.frameSize = s.subFrameSize * int(opts.DMDConfig.NumSubFrames)
	s.dots = make([]uint8, s.frameSize)
}

func (s *procSystem) HandleAction(action spin.Action) {
	switch act := action.(type) {
	case spin.DriverOff:
		s.driverOff(act)
	case spin.DriverOn:
		s.driverOn(act)
	case spin.DriverPulse:
		s.driverPulse(act)
	case spin.DriverPWM:
		s.driverPWM(act)
	case spin.FlippersOn:
		s.flippersOn(act)
	case spin.FlippersOff:
		s.flippersOff(act)
	case spin.RegisterCoil:
		s.registerCoil(act)
	case spin.RegisterDisplay:
		s.registerDisplay(act)
	case spin.RegisterFlasher:
		s.registerFlasher(act)
	case spin.RegisterFlipper:
		s.registerFlipper(act)
	case spin.RegisterLamp:
		s.registerLamp(act)
	case spin.RegisterMagnet:
		s.registerMagnet(act)
	case spin.RegisterMotor:
		s.registerMotor(act)
	case spin.RegisterSwitch:
		s.registerSwitch(act)
	}
}

func (s *procSystem) Service() {
	n, err := s.proc.GetEvents(s.events)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < n; i++ {
		e := s.events[i]
		if e.EventType != pinproc.EventTypeSwitchClosedDebounced && e.EventType != pinproc.EventTypeSwitchOpenDebounced {
			continue
		}
		sw, ok := s.switches[uint8(e.Value)]
		if !ok {
			spin.Warn("unknown switch: %v", e.Value)
			continue
		}
		released := e.EventType == pinproc.EventTypeSwitchOpenDebounced
		if sw.NC {
			released = !released
		}
		s.eng.Post(spin.SwitchEvent{ID: sw.ID, Released: released})
	}
	if s.source != nil {
		for i := 0; i < len(s.dots); i++ {
			s.dots[i] = 0
		}
		for y := 0; y < s.source.Height(); y++ {
			for x := 0; x < s.source.Width(); x++ {
				color := s.source.At(x, y)
				gray := spin.RGBToGray(color)
				i := (y*s.source.Width() + x) / 8
				b := uint8(1 << ((y*s.source.Width() + x) % 8))
				on := gray > 0
				if on {
					s.dots[(0*s.subFrameSize)+i] |= b
					s.dots[(1*s.subFrameSize)+i] |= b
					s.dots[(2*s.subFrameSize)+i] |= b
					s.dots[(3*s.subFrameSize)+i] |= b
				}
			}
		}
		if err := s.proc.DMDDraw(s.dots); err != nil {
			log.Fatal(err)
		}
	}
	if err := s.proc.DriverWatchdogTickle(); err != nil {
		log.Fatalf("unable to tickle watchdog: %v", err)
	}
	if err := s.proc.FlushWriteData(); err != nil {
		log.Fatalf("unable to flush data: %v", err)
	}
}

func (s *procSystem) driverOn(act spin.DriverOn) {
	driver, ok := s.drivers[act.ID]
	if !ok {
		spin.Warn("no such driver: %v", act.ID)
		return
	}
	if driver.Type == spin.Coil || driver.Type == spin.Flasher {
		spin.Warn("cannot enable: %v", act.ID)
		return
	}
	addr := driver.Addr.(uint8)
	s.proc.DriverEnable(addr)
}

func (s *procSystem) driverOff(act spin.DriverOff) {
	driver, ok := s.drivers[act.ID]
	if !ok {
		spin.Warn("no such driver: %v", act.ID)
		return
	}
	addr := driver.Addr.(uint8)
	s.proc.DriverDisable(addr)
}

func (s *procSystem) driverPulse(act spin.DriverPulse) {
	driver, ok := s.drivers[act.ID]
	if !ok {
		spin.Warn("no such driver: %v", act.ID)
		return
	}
	time := uint8(act.Time)
	if time == 0 && driver.Type == spin.Coil {
		time = uint8(s.opts.DefaultCoilPulseTime)
	}
	if time == 0 && driver.Type == spin.Flasher {
		time = uint8(s.opts.DefaultFlasherPulseTime)
	}
	if time <= 0 {
		spin.Warn("invalid pulse time: %v", time)
		return
	}
	addr := driver.Addr.(uint8)
	s.proc.DriverPulse(addr, time)
}

func (s *procSystem) driverPWM(act spin.DriverPWM) {
	driver, ok := s.drivers[act.ID]
	if !ok {
		spin.Warn("no such driver: %v", act.ID)
		return
	}
	timeOn := uint8(act.TimeOn)
	timeOff := uint8(act.TimeOff)
	if timeOff == 0 {
		spin.Warn("PWM off time cannot be zero")
		return
	}
	addr := driver.Addr.(uint8)
	s.proc.DriverPatter(addr, timeOn, timeOff, 0, false)
}

func (s *procSystem) flippersOn(act spin.FlippersOn) {
	flipperIDs := act.FlipperIDs
	if len(flipperIDs) == 0 {
		flipperIDs = make([]string, 0)
		for id := range s.flippers {
			flipperIDs = append(flipperIDs, id)
		}
	}

	for _, flipperID := range flipperIDs {
		flipper, ok := s.flippers[flipperID]
		if !ok {
			spin.Warn("no such flipper: %v", flipperID)
			continue
		}
		var pulseState, holdOnState, holdOffState pinproc.DriverState
		s.proc.DriverGetState(flipper.PowerCoilAddr.(uint8), &pulseState)
		s.proc.DriverGetState(flipper.HoldCoilAddr.(uint8), &holdOnState)
		s.proc.DriverGetState(flipper.HoldCoilAddr.(uint8), &holdOffState)

		pinproc.DriverStatePulse(&pulseState, 35)
		pinproc.DriverStatePulse(&holdOnState, 0)
		pinproc.DriverStateDisable(&holdOffState)

		if err := s.proc.SwitchUpdateRule(
			flipper.SwitchAddr.(uint8),
			pinproc.EventTypeSwitchClosedNondebounced,
			pinproc.SwitchRule{},
			[]pinproc.DriverState{pulseState, holdOnState},
			true); err != nil {
			log.Fatal(err)
		}
		if err := s.proc.SwitchUpdateRule(
			flipper.SwitchAddr.(uint8),
			pinproc.EventTypeSwitchOpenNondebounced,
			pinproc.SwitchRule{},
			[]pinproc.DriverState{holdOffState},
			true); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *procSystem) flippersOff(act spin.FlippersOff) {
	flipperIDs := act.FlipperIDs
	if flipperIDs == nil {
		flipperIDs = make([]string, 0)
		for id := range s.flippers {
			flipperIDs = append(flipperIDs, id)
		}
	}

	for _, flipperID := range flipperIDs {
		flipper, ok := s.flippers[flipperID]
		if !ok {
			spin.Warn("no such flipper: %v", flipperID)
			continue
		}

		// Flipper buttons are optos so open == button pressed
		if err := s.proc.SwitchUpdateRule(
			flipper.SwitchAddr.(uint8),
			pinproc.EventTypeSwitchOpenNondebounced,
			pinproc.SwitchRule{},
			[]pinproc.DriverState{},
			true); err != nil {
			log.Fatal(err)
		}
		if err := s.proc.SwitchUpdateRule(
			flipper.SwitchAddr.(uint8),
			pinproc.EventTypeSwitchClosedNondebounced,
			pinproc.SwitchRule{},
			[]pinproc.DriverState{},
			true); err != nil {
			log.Fatal(err)
		}
	}
}

func (s *procSystem) registerCoil(act spin.RegisterCoil) {
	driver := spin.Driver{
		ID:   act.ID,
		Type: spin.Coil,
		Addr: act.Addr,
	}
	s.drivers[act.ID] = driver
}

func (s *procSystem) registerDisplay(act spin.RegisterDisplay) {
	if act.ID != s.opts.ID {
		return
	}
	s.source = act.Display
}

func (s *procSystem) registerFlasher(act spin.RegisterFlasher) {
	driver := spin.Driver{
		ID:   act.ID,
		Type: spin.Flasher,
		Addr: act.Addr,
	}
	s.drivers[act.ID] = driver
}

func (s *procSystem) registerFlipper(act spin.RegisterFlipper) {
	flipper := spin.Flipper{
		ID:            act.ID,
		SwitchAddr:    act.SwitchAddr,
		PowerCoilAddr: act.PowerCoilAddr,
		HoldCoilAddr:  act.HoldCoilAddr,
	}
	s.flippers[act.ID] = flipper
}

func (s *procSystem) registerLamp(act spin.RegisterLamp) {
	driver := spin.Driver{
		ID:   act.ID,
		Type: spin.Lamp,
		Addr: act.Addr,
	}
	s.drivers[act.ID] = driver
}

func (s *procSystem) registerMagnet(act spin.RegisterMagnet) {
	driver := spin.Driver{
		ID:   act.ID,
		Type: spin.Magnet,
		Addr: act.Addr,
	}
	s.drivers[act.ID] = driver
}

func (s *procSystem) registerMotor(act spin.RegisterMotor) {
	driver := spin.Driver{
		ID:   act.ID,
		Type: spin.Motor,
		Addr: act.Addr,
	}
	s.drivers[act.ID] = driver
}

func (s *procSystem) registerSwitch(act spin.RegisterSwitch) {
	sw := spin.Switch{
		ID: act.ID,
		NC: act.NC,
	}
	addr := act.Addr.(uint8)
	s.switches[addr] = sw

	rule := pinproc.SwitchRule{NotifyHost: true}
	if err := s.proc.SwitchUpdateRule(addr, pinproc.EventTypeSwitchClosedDebounced, rule, nil, false); err != nil {
		log.Fatal(err)
	}
	if err := s.proc.SwitchUpdateRule(addr, pinproc.EventTypeSwitchOpenDebounced, rule, nil, false); err != nil {
		log.Fatal(err)
	}
}