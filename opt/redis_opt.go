package opt

import "github.com/spf13/pflag"

type RedisOpt struct {
	Host     string `json:"host"                     mapstructure:"host"                     description:"Redis service host address"`
	Port     int    `json:"port"`
	Username string `json:"username"                 mapstructure:"username"`
	Password string `json:"password"                 mapstructure:"password"`
}

func (o *RedisOpt) Validate() bool {
	return true
}

func (o *RedisOpt) AddFlags(fss *pflag.FlagSet) {
	fss.StringVar(&o.Host, "redis.host", o.Host, "Hostname of your Redis server")
	fss.IntVar(&o.Port, "redis.port", o.Port, "The port the Redis server is listening to")
	fss.StringVar(&o.Username, "redis.username", o.Username, "The username for Redis server")
	fss.StringVar(&o.Password, "redis.password", o.Password, "The username for Redis password")
}

func NewRedisOpt() *RedisOpt {
	return &RedisOpt{
		Host:     "127.0.0.1",
		Port:     6379,
		Username: "",
		Password: "",
	}
}
