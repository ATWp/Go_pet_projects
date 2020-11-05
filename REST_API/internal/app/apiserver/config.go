package apiserver

import (
	"github.com/ATWp/Go_pet_projects/REST_API/internal/app/store"
)

// Config toml
type Config struct {
	//Port
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

// New Config
func NewConfig *Config{
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}