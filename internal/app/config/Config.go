package config

type Config struct {
	ServerAddress string `env:"SERVER_ADDRESS"`
	BaseURL       string `env:"BASE_URL"`
}

var cfg Config

func AppConfig() Config {
	//err := env.Parse(&cfg)
	//if err != nil {
	cfg = Config{ServerAddress: "8080", BaseURL: "http://localhost"}
	//}
	return cfg
}
