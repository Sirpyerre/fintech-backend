package config

import (
	"context"
	"log"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Port        string `env:"PORT,default=8000"`
	DBConfig    DBConfig
	Env         string `env:"ENV,default=development"`
	LogLevel    string `env:"LOG_LEVEL,default=info"`
	WorkerCount int    `env:"MAX_WORKERS,default=1"`
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
