package serverconfig

import (
	"HelloWorld/src/gen/redisconfig"
	"context"
	"github.com/apple/pkl-go/pkl"
)

type ServerConfig struct {
	Port        int                      `mapstructure:"port" pkl:"port"`
	Hostname    string                   `mapstructure:"hostname" pkl:"host"`
	RedisConfig *redisconfig.RedisConfig `mapstructure:"redis" pkl:"redis"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a AppConfig
func LoadFromPath(ctx context.Context, path string) (ret *ServerConfig, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return nil, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a AppConfig
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (*ServerConfig, error) {
	var ret ServerConfig
	if err := evaluator.EvaluateModule(ctx, source, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
