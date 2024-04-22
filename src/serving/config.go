package serving

type ServerConfig struct {
	Port        int          `mapstructure:"port"`
	Hostname    string       `mapstructure:"hostname"`
	RedisConfig *RedisConfig `mapstructure:"redis"`
	Paths       []string     `mapstructure:"paths"`
}

type RedisConfig struct {
	Address string `mapstructure:"address"`
}
