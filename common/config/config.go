package config

import "os"

type Config struct {
	PortAddr string
}

func NewConfig() *Config {
	portAddr := os.Getenv("PORT_ADDR")
	return &Config{
		PortAddr: portAddr,
	}
}
