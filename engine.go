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
	Settings *Settings
	DevMode  bool
	runDB    *redis.Client
	varDB    *redis.Client
}

func NewEngine(settings *Settings) (*Engine, error) {
	if settings.Dir == "" {
		settings.Dir = "./project"
	}

	e := &Engine{
		Config:   NewConfig(),
		Settings: settings,
	}

	if err := e.Config.AddFile(e.PathTo("project.hcl")); err != nil {
		return nil, err
	}
	e.Settings.Merge(e.Config.Settings)

	e.runDB = redis.NewClient(&redis.Options{Addr: settings.RedisRunAddress})
	e.varDB = redis.NewClient(&redis.Options{Addr: settings.RedisVarAddress})
	return e, nil
}

func (e *Engine) PathTo(name string) string {
	return path.Join(e.Settings.Dir, name)
}

func (e *Engine) Error(args ...any) {
	panic(logMsg(args...))
}

func (e *Engine) Warn(args ...any) {
	msg := logMsg(args...)
	if e.DevMode {
		panic(msg)
	}
	log.Print(msg)
}

func (e *Engine) Note(args ...any) {
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
		Stream: "mq",
		Values: []any{
			"type", reflect.TypeOf(message).Name(),
			"payload", payload,
		},
	})
	if result.Err() != nil {
		e.Error(result.Err())
	}
}
