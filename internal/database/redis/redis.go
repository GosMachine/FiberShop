package redis

import (
	"FiberShop/internal/transport/grpc/auth"
	"context"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type Redis struct {
	Ctx        context.Context
	Client     *redis.Client
	Log        *zap.Logger
	AuthClient *auth.Client
}

func New(log *zap.Logger, authClient *auth.Client) *Redis {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &Redis{Ctx: ctx, Client: client, Log: log, AuthClient: authClient}
}
