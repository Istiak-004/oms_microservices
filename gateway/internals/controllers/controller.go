package controllers

import (
	pb "github.com/istiak-004/common/api"
)

type Controller struct {
	grpcClient pb.OrderServiceClient
}

func NewController(grpcClient pb.OrderServiceClient) *Controller {
	return &Controller{
		grpcClient: grpcClient,
	}
}
