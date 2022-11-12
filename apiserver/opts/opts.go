package opts

import (
	"github.com/spf13/pflag"
	"kkgo-app/opt"
)

type Opts struct {
	RedisOpt *opt.RedisOpt `json:"redis"    mapstructure:"redis"`
}

func NewOpts() *Opts {
	return &Opts{
		RedisOpt: opt.NewRedisOpt(),
	}
}

func (o *Opts) AddFlags(fs *pflag.FlagSet) {
	o.RedisOpt.AddFlags(fs)
}
