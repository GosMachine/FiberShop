package handlers

import (
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/transport/grpc"

	"go.uber.org/zap"
)

type Handle struct {
	Log   *zap.Logger
	Grpc  *grpc.Grpc
	Db    *postgres.Storage
	Redis *redis.Redis
}

func New(log *zap.Logger, grpc *grpc.Grpc, db *postgres.Storage, redis *redis.Redis) *Handle {
	return &Handle{Log: log, Grpc: grpc, Db: db, Redis: redis}
}
