package config

import (
	"github.com/caarlos0/env/v10"
	"log"
)

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

var cfg Config

func AppConfig() Config {
	cfg = Config{ServerAddress: "8080", BaseURL: "http://localhost112233"}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
