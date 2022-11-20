package apiserver

import "iam-mini/apiserver/config"

type Config struct {
	*config.Options
}

func CreateConfigFromOptions(opts *config.Options) (*Config, error) {
	return &Config{opts}, nil
}
