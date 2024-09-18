package config

import "github.com/caarlos0/env/v6"

type Config struct {
	TelegamToken string `env:"TELEGRAM_TOKEN"`
}

func NewConfig() (*Config, error) {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
