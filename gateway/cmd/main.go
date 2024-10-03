package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	pb "github.com/istiak-004/common/api"
	cnf "github.com/istiak-004/gateway/config"
	"github.com/istiak-004/gateway/internals/routers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	//_ "github.com/joho/godotenv/autoload"
)

func main() {
	config := cnf.LoadEnv()
	fmt.Println("Config file path is: +++++++++++++>>", config)
	grpcConnection, err := grpc.NewClient(config.GrpcAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))

	fmt.Println("grpc connection :##################", grpcConnection)

	if err != nil {
		fmt.Printf("Failed to connect to gRPC server: %v", err)
		return
	}
	fmt.Printf("Order service has been started at %v", config.GrpcAddress)
	defer grpcConnection.Close() // close the connection

	// create  a grpc client connection

	orderServiceClient := pb.NewOrderServiceClient(grpcConnection)

	fmt.Println("------------------------------------")
	fmt.Println(orderServiceClient)
	fmt.Println("------------------------------------")

	engine := routers.NewRouter(gin.Default(), orderServiceClient)
	engine.SetupRoutes()

	fmt.Printf("Starting server at port %v", config.PortAddr)
	engine.Run(config.PortAddr) // Start the server on port 8080
}
