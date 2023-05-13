package config

import (
	"github.com/caarlos0/env/v8"
)

type Config struct {
	EnvironmentProperty string `env:"ENVIRONMENT_PROPERTY" envDefault:"default"`
}

func NewConfig() (cfg Config, err error) {
	if err = env.Parse(&cfg); err != nil {
		return
	}

	return
}
