package opts

import (
	"github.com/spf13/pflag"
	"iam-mini/opt"
)

type Opts struct {
	RedisOpt    *opt.RedisOpt    `json:"redis"    mapstructure:"redis"`
	InsecureOpt *opt.InsecureOpt `json:"insecure"    mapstructure:"insecure"`
}

func NewOpts() *Opts {
	return &Opts{
		RedisOpt:    opt.NewRedisOpt(),
		InsecureOpt: opt.NewInsecureOpt(),
	}
}

func (o *Opts) AddFlags(fs *pflag.FlagSet) {
	o.RedisOpt.AddFlags(fs)
	o.InsecureOpt.AddFlags(fs)
}

func (o *Opts) ApplyTo()
