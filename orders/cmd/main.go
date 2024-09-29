package main

import (
	"context"

	"github.com/istiak-004/orders/internals/services"
)

func main() {
	store := services.NewStore()
	service := services.NewOrderService(store)

	service.CreateOrder(context.Background())
}
