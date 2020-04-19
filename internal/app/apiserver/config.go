package apiserver

import "github.com/seblex/rest-api-go/internal/app/store"

type Config struct {
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8020",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}