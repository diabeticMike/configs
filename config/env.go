package config

import (
	"github.com/kelseyhightower/envconfig"
)

func NewEnv() (Config, error) {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		return Config{}, err
	}

	return conf, nil
}
