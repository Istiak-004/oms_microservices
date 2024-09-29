package services

import "context"

type store struct {
}

func NewStore() *store {
	return &store{}
}

func (s *store) Create(ctx context.Context) error {
	// Implementation for creating a new order
	return nil
}
