package handlers

import (
	"FiberShop/internal/transport/grpc/auth"
	"go.uber.org/zap"
)

type Handle struct {
	Log    *zap.Logger
	Client *auth.Client
}

func New(log *zap.Logger, client *auth.Client) *Handle {
	return &Handle{Log: log, Client: client}
}
