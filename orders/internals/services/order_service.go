package services

import (
	"context"

	"github.com/istiak-004/orders/internals/types"
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
