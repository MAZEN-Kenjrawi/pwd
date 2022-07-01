package infrastructure

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env string `env:"ENV" envDefault:"dev"`

	Storage struct {
		Mode string `env:"STORAGE_MODE" envDefault:"filesystem"`
		Url  string `env:"STORAGE_URL"`
	}
}

func NewConfig() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)

	return cfg, err
}
