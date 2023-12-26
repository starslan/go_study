package config

import (
	"github.com/caarlos0/env/v10"
	"log"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS" envDefault:"8080"`
	BaseURL       string `env:"BASE_URL" envDefault:"http://localhost"`
}

func AppConfig() Config {
	var cfg = Config{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
