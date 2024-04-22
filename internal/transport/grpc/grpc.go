package grpc

import (
	"FiberShop/internal/config"
	"FiberShop/internal/transport/grpc/auth"
	"FiberShop/internal/transport/grpc/product"
	"context"
)

type Grpc struct {
	Auth    *auth.Client
	Product *product.Client
}

func New(cfg *config.Config) (*Grpc, error) {
	var (
		grpc Grpc
		err  error
	)
	grpc.Auth, err = auth.New(context.Background(), cfg.Clients.Auth.Address,
		cfg.Clients.Timeout, cfg.Clients.RetriesCount)
	if err != nil {
		return nil, err
	}
	grpc.Product, err = product.New(context.Background(), cfg.Clients.Product.Address,
		cfg.Clients.Timeout, cfg.Clients.RetriesCount)
	if err != nil {
		return nil, err
	}
	return &grpc, nil
}
