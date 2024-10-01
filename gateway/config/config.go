package config

import (
	"fmt"
	"os"

	"github.com/istiak-004/common/config"
	"github.com/joho/godotenv"
)

func LoadEnv() *config.Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading.env file", err)
		os.Exit(1)
	}
	config := config.NewConfig()
	return config
}

// type Config struct {
// 	GrpcConnection *grpc.ClientConn `config:"grpc"`
// }

// func NewConfig(grpcTarget string) *Config {
// 	conn, err := grpc.NewClient(grpcTarget)
// 	if err != nil {
// 		fmt.Printf("Failed to connect to gRPC server: %v", err)
// 		return &Config{}
// 	}
// 	return &Config{
// 		GrpcConnection: conn,
// 	}
// }
