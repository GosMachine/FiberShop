package redis

import (
	"FiberShop/internal/database/postgres"
	"context"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Ctx    context.Context
	Client *redis.Client
	Log    *zap.Logger
	Db     *postgres.Storage
}

func New(log *zap.Logger, db *postgres.Storage) *Redis {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Redis{Ctx: ctx, Client: client, Log: log, Db: db}
}
