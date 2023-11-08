package spin

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"path"
	"reflect"

	"github.com/redis/go-redis/v9"
)

type Engine struct {
	Config   *Config
	DevMode  bool
	Settings *Settings
	devices  []Device
	modules  map[string]struct{}
	runDB    *redis.Client
	varDB    *redis.Client
}

func NewEngine(settings *Settings) *Engine {
	if settings.Dir == "" {
		settings.Dir = "./project"
	}
	e := &Engine{
		Config:   NewConfig(),
		Settings: settings,
		modules:  make(map[string]struct{}),
	}
	return e
}

func (e *Engine) Init() error {
	if err := e.Config.AddFile(e.PathTo(e.Settings.ConfigFile)); err != nil {
		return err
	}
	e.Settings.Merge(e.Config.Settings)

	e.runDB = redis.NewClient(&redis.Options{Addr: e.Settings.RedisRunAddress})
	e.varDB = redis.NewClient(&redis.Options{Addr: e.Settings.RedisVarAddress})

	// Runtime database should be cleared out on each start. This also
	// verifies that the database is up and running. Send a ping to the
	// variable database to see if that is also running.
	ctx := context.Background()
	if resp := e.runDB.FlushAll(ctx); resp.Err() != nil {
		return resp.Err()
	}
	if resp := e.varDB.Ping(ctx); resp.Err() != nil {
		return resp.Err()
	}

	e.Log("spin version %v - %v", Version, Date)

	for _, conf := range e.Config.AudioDevices {
		d, ok := NewDevice(conf)
		if !ok {
			e.Error("no such handler: %v", conf.Handler)
		}
		if d.Init(e) {
			go func() {
				for d.Process(e) {
				}
			}()
			e.devices = append(e.devices, d)
		}
	}

	for _, id := range e.Config.Load {
		if _, exists := e.modules[id]; exists {
			continue
		}
		e.modules[id] = struct{}{}
		e.Log("loading module: %v", id)
		e.Send(Load{ID: id})
	}

	return nil

}

func (e *Engine) NewQueueClient() *QueueClient {
	return NewQueueClient(e.runDB)
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
	log.Panic(logMsg(args...))
}

// Warn writes a message to the log. If DevMode is set to true, this will then
// panic, otherwise execution will continue. This method should be called on
// errors that are not serious enough to exit the application but should be
// immediately addressed by the programmer (for example, a missing sound file)
func (e *Engine) Warn(args ...any) {
	msg := logMsg(args...)
	if e.DevMode {
		log.Panic(msg)
	}
	log.Print(msg)
}

// Log writes a message to the log.
func (e *Engine) Log(args ...any) {
	log.Print(logMsg(args...))
}

func logMsg(args ...any) string {
	if len(args) == 0 {
		return ""
	}
	format, others := fmt.Sprintf("%v", args[0]), args[1:]
	return fmt.Sprintf(format, others...)
}

func (e *Engine) Send(message any) {
	payload, err := json.Marshal(message)
	if err != nil {
		e.Warn(err)
		return
	}
	ctx := context.Background()
	result := e.runDB.XAdd(ctx, &redis.XAddArgs{
		Stream: MessageQueueKey,
		Values: []any{
			"type", reflect.TypeOf(message).Name(),
			"payload", payload,
		},
	})
	if result.Err() != nil {
		e.Error(result.Err())
	}
}
