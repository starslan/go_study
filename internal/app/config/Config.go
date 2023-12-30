package config

import (
	"flag"
	"github.com/caarlos0/env/v10"
	"log"
)

var sa = "localhost:8080"
var bu = "http://localhost:8080"
var fp = ""

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
	FilePath      string `env:"FILE_STORAGE_PATH"`
}

func NewConfig() Config {
	var cfg = Config{}
	baseURL := flag.String("b", bu, "a string")
	path := flag.String("f", fp, "a string")
	serverAddress := flag.String("a", sa, "a string")
	flag.Parse()
	cfg.BaseURL = *baseURL
	cfg.ServerAddress = *serverAddress
	cfg.FilePath = *path
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}
