package server

import (
	"github.com/spf13/pflag"
	"iam-mini/generic/opt"
)

type Opts struct {
	RedisOpt    *opt.RedisOpt    `json:"redis"    mapstructure:"redis"`
	InsecureOpt *opt.InsecureOpt `json:"insecure"    mapstructure:"insecure"`
	JwtOpt      *opt.JwtOptions  `json:"jwt"    mapstructure:"jwt"`
}

func NewOpts() *Opts {
	return &Opts{
		RedisOpt:    opt.NewRedisOpt(),
		InsecureOpt: opt.NewInsecureOpt(),
		JwtOpt:      opt.NewJwtOptions(),
	}
}

func (o *Opts) AddFlags(fs *pflag.FlagSet) {
	o.RedisOpt.AddFlags(fs)
	o.InsecureOpt.AddFlags(fs)
	o.JwtOpt.AddFlags(fs)
}

func (o *Opts) ApplyTo(config *Config) {
	o.InsecureOpt.ApplyTo(config.Insecure)
}
