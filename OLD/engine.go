package spin

import (
	"context"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/redis/go-redis/v9"
)

// Engine provides the run environment and the main game loop for execution.
type Engine struct {
	Config    *Config
	DevMode   bool
	Settings  *Settings
	StartTime time.Time
	init      bool
	devices   []Device
	modules   map[string]struct{}
	runDB     *redis.Client
	varDB     *redis.Client
	shutdown  chan struct{}
}

// NewEngine creates an engine using the specified settings.
func NewEngine(settings *Settings) *Engine {
	if settings.Dir == "" {
		settings.Dir = "./project"
	}

	// Disable all default formatting as a specific format will be used
	// instead.
	log.SetFlags(0)

	e := &Engine{
		Config:    NewConfig(),
		DevMode:   settings.DevMode,
		Settings:  settings,
		StartTime: time.Now(),
		modules:   make(map[string]struct{}),
		shutdown:  make(chan struct{}),
	}
	return e
}

// Init reads in project configuration files, clears the runtime database,
// and starts goroutines for devices as necessary.
func (e *Engine) Init() error {
	if e.init {
		return fmt.Errorf("already initialized")
	}
	if err := e.Config.AddFile(e.PathTo(e.Settings.ConfigFile)); err != nil {
		return err
	}
	e.Settings.Merge(e.Config.Settings)

	// e.runDB = redis.NewClient(&redis.Options{Addr: e.Settings.RedisRunAddress})
	// e.varDB = redis.NewClient(&redis.Options{Addr: e.Settings.RedisVarAddress})

	// Runtime database should be cleared out on each start. This also
	// verifies that the database is up and running. Send a ping to the
	// variable database to see if that is also running.
	ctx := context.Background()
	if resp := e.runDB.FlushAll(ctx); resp.Err() != nil {
		return fmt.Errorf("unable to flush runtime database: %v", resp.Err())
	}
	if resp := e.varDB.Ping(ctx); resp.Err() != nil {
		return fmt.Errorf("unable to connect to variable database: %v", resp.Err())
	}

	for _, conf := range e.Config.AudioDevices {
		d, err := NewDevice(conf.Type, conf)
		if err != nil {
			return err
		}
		if d.Init(e) {
			go func() {
				for d.Process(e) {
				}
			}()
			e.devices = append(e.devices, d)
		}
	}

	e.Log("spin version %v (%v)", Version, Date)

	// stream := e.NewStreamClient()
	// for _, id := range e.Config.Load {
	// 	if _, exists := e.modules[id]; exists {
	// 		continue
	// 	}
	// 	e.modules[id] = struct{}{}
	// 	e.Debug("loading module: %v", id)
	// 	if err := stream.Send(Load{ID: id}); err != nil {
	// 		e.Error(err)
	// 	}
	// }
	// if err := stream.Send(Load{}); err != nil {
	// 	e.Error(err)
	// }

	e.init = true
	return nil
}

// Process is the main body of the run loop. This method is called 60 times
// a second (every 16.67 milliseconds). Calling this method is useful for
// testing but normal use is to simply call Run.
func (e *Engine) Process(t time.Time) {
}

// Run executes the game engine loop. This loop repeats until Shutdown
// is called.
func (e *Engine) Run() {
	if !e.init {
		panic("not initialized")
	}
	e.Debug("ready")
	ticker := time.NewTicker(16670 * time.Microsecond)

	var done bool
	for !done {
		select {
		case <-e.shutdown:
			done = true
		case t := <-ticker.C:
			e.Process(t)
		}
	}
	e.Debug("shutdown complete")
}

// Shutdown sends a message to terminate the main run loop.
func (e *Engine) Shutdown() {
	e.shutdown <- struct{}{}
}

// PathTo returns a path that in the joined value of the project directory
// in Settings.Dir and the provided name.
func (e *Engine) PathTo(name string) string {
	return path.Join(e.Settings.Dir, name)
}

// Error writes a message to the log file and then panics. This method should
// be called on unrecoverable errors when the program should exit and be
// restarted by systemd.
func (e *Engine) Error(args ...any) {
	tsMsg, msg := e.logMsg(args...)
	log.Print(tsMsg)
	panic(msg)
}

// Warn writes a message to the log. If DevMode is set to true, this will then
// panic, otherwise execution will continue. This method should be called on
// errors that are not serious enough to exit the application but should be
// immediately addressed by the programmer (for example, a missing sound file)
func (e *Engine) Warn(args ...any) {
	tsMsg, msg := e.logMsg(args...)
	log.Print(tsMsg)
	if e.DevMode {
		panic(msg)
	}
}

// Log writes a message to the log.
func (e *Engine) Log(args ...any) {
	tsMsg, _ := e.logMsg(args...)
	log.Print(tsMsg)
}

// Debug writes a message to the log if DevMode is true.
func (e *Engine) Debug(args ...any) {
	if e.DevMode {
		tsMsg, _ := e.logMsg(args...)
		log.Print(tsMsg)
	}
}

func (e *Engine) logMsg(args ...any) (string, string) {
	if len(args) == 0 {
		return "", ""
	}
	now := time.Now()
	diff := float64(now.Sub(e.StartTime).Milliseconds()) / 1000

	format, others := fmt.Sprintf("%v", args[0]), args[1:]
	msg := fmt.Sprintf(format, others...)
	return fmt.Sprintf("[%10.3f] %v", diff, msg), msg
}
