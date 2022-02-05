package spin

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"time"

	"github.com/drop-target-pinball/spin/rlog"
)

const (
	DebugStackTrace = "StackTrace"
)

var (
	logger *rlog.Logger // FIXME: This should be in the Engine
)

func init() {
	logger = rlog.New(60*time.Second, log.Default())
}

type loggingSystem struct {
	startTime time.Time
}

func RegisterLoggingSystem(eng *Engine) {
	sys := &loggingSystem{
		startTime: time.Now(),
	}
	eng.RegisterActionHandler(sys)
	eng.RegisterEventHandler(sys)
}

func (s *loggingSystem) HandleAction(action Action) {
	log.Printf("[%v] %v", s.elapsedTime(), FormatAction(action))
	switch act := action.(type) {
	case Debug:
		s.debug(act)
	}
}

func (s *loggingSystem) HandleEvent(evt Event) {
	logger.Printf("[%v] %v", s.elapsedTime(), FormatEvent(evt))

	// FIXME: this is a hack
	sw, ok := evt.(SwitchEvent)
	if ok {
		if sw.ID == "jd.SwitchBuyInButton" {
			dumpLog()
		}
	}
}

func (s *loggingSystem) debug(act Debug) {
	switch act.ID {
	case DebugStackTrace:
		debugStackTrace()
	}
}

func (s *loggingSystem) elapsedTime() string {
	now := time.Now()
	diff := float64(now.Sub(s.startTime).Milliseconds()) / 1000
	return fmt.Sprintf("%10.3f", diff)
}

func Error(format string, a ...interface{}) {
	format = "(*) " + format
	logger.Printf(format, a...)
}

func Warn(format string, a ...interface{}) {
	format = "(!) " + format
	logger.Printf(format, a...)
}

func Info(format string, a ...interface{}) {
	format = "(?) " + format
	logger.Printf(format, a...)
}

func Log(format string, a ...interface{}) {
	logger.Printf(format, a...)
}

func debugStackTrace() {
	profile := pprof.Lookup("goroutine")
	profile.WriteTo(os.Stdout, 1)
}

func dumpLog() {
	timestamp := time.Now().Format("060102-130405")
	dir := os.Getenv("SPIN_DIR") + "/log"
	if err := os.MkdirAll(dir, 0o755); err != nil {
		Error("unable to create directory %v: %v", dir, err)
	}
	file := dir + "/spin-" + timestamp + ".log"
	Log("writing log to file %v", file)
	f, err := os.Create(file)
	if err != nil {
		Error("unable to create file %v: %v", file, err)
	}
	defer f.Close()

	for _, message := range logger.Messages() {
		f.WriteString(message + "\n")
	}
	profile := pprof.Lookup("goroutine")
	profile.WriteTo(f, 1)
}
