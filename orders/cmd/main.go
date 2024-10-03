package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	pb "github.com/istiak-004/common/api"
	"github.com/istiak-004/orders/internals/services"
	"google.golang.org/grpc"
)

func main() {

	// grpc server connection
	grpcServer := grpc.NewServer()
	conn, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatal("Failed to connect to listen ", err)
	}
	defer conn.Close()
	grpcHandler := services.NewGRPCHandler()

	pb.RegisterOrderServiceServer(grpcServer, grpcHandler)
	go func() {
		log.Println("Server started on port 3000")
		if err := grpcServer.Serve(conn); err != nil {
			log.Fatal("Failed to serve: ", err)
		}
	}()

	// Graceful shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	log.Println("Stopping gRPC server...")
	grpcServer.GracefulStop()
	// Wait for graceful shutdown

	// store := services.NewStore()
	// service := services.NewOrderService(store)

	// service.CreateOrder(context.Background())

}
