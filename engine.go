package spin

import (
	"github.com/redis/go-redis/v9"
)

type Engine struct {
	Config *Config
	runDB  *redis.Client
	varDB  *redis.Client
}

func NewEngine(settings Settings) (*Engine, error) {
	e := &Engine{
		Config: NewConfig(),
	}

	if err := e.Config.AddFile("project.hcl"); err != nil {
		return nil, err
	}

	e.runDB = redis.NewClient(&redis.Options{Addr: settings.RedisRunAddress})
	e.varDB = redis.NewClient(&redis.Options{Addr: settings.RedisVarAddress})
	return e, nil
}
