package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port     string `env:"PORT" envDefault:"8080"`
	DBConfig DBConfig
	Env      string `env:"ENV" envDefault:"development"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type DBConfig struct {
	DatabaseURL string `env:"DATABASE_URL,required"`
}

func NewConfiguration(ctx context.Context) *Config {
	cfg := new(Config)

	if err := envconfig.Process(ctx, cfg); err != nil {
		log.Fatal(err)
	}

	return cfg
}
