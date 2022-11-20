package config

import (
	"github.com/spf13/pflag"
	"iam-mini/apiserver/server"
	"iam-mini/generic/options"
)

type Options struct {
	RedisOpt    *opt.RedisOpt     `json:"redis"    mapstructure:"redis"`
	InsecureOpt *opt.InsecureOpt  `json:"insecure"    mapstructure:"insecure"`
	JwtOpt      *opt.JwtOptions   `json:"jwt"    mapstructure:"jwt"`
	MySQLOpt    *opt.MySQLOptions `json:"mysql"    mapstructure:"mysql"`
}

func NewOpts() *Options {
	return &Options{
		RedisOpt:    opt.NewRedisOpt(),
		InsecureOpt: opt.NewInsecureOpt(),
		JwtOpt:      opt.NewJwtOptions(),
		MySQLOpt:    opt.NewMySQLOptions(),
	}
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {
	o.RedisOpt.AddFlags(fs)
	o.InsecureOpt.AddFlags(fs)
	o.JwtOpt.AddFlags(fs)
	o.MySQLOpt.AddFlags(fs)
}

func (o *Options) ApplyTo(config *server.Config) {
	o.InsecureOpt.ApplyTo(config.Insecure)
}
