package config

import "os"

type Config struct {
	PortAddr    string `json:"port_address"`
	GrpcAddress string `json:"grpc_address"`
}

func NewConfig() *Config {
	portAddr := os.Getenv("PORT_ADDR")
	grpcAddress := os.Getenv("GRPC_ADDR")
	return &Config{
		PortAddr:    portAddr,
		GrpcAddress: grpcAddress,
	}
}
