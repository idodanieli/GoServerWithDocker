package redisconfig

import "fmt"

type RedisConfig struct {
	Address string `mapstructure:"address" pkl:"address"`
}

func (rc *RedisConfig) String() string {
	return fmt.Sprintf("Redis: %v", rc.Address)
}
