package services

import (
	"context"
	"fmt"

	pb "github.com/istiak-004/common/api"
	"github.com/istiak-004/orders/internals/types"
	"google.golang.org/grpc"
)

type ordersService struct {
	store types.OrderStore
}

func NewOrderService(store types.OrderStore) *ordersService {
	return &ordersService{
		store: store,
	}
}

func (s *ordersService) CreateOrder(context.Context) error {
	return nil
}

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler() *grpcHandler {
	return &grpcHandler{}
}

func NewGRPCServiceHandler(orderServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(orderServer, handler)

}

func (g *grpcHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	fmt.Println("@@@@@@@@@@@@@@@grpc handler create order@@@@@@@@@@@@@@@@@@@@")
	resp := &pb.CreateOrderResponse{
		OrderId:    "123",
		Status:     "running",
		TotalPrice: 1000,
		Items:      []*pb.Item{},
	}

	return resp, nil

}
