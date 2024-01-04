package app

import (
	fiberapp "FiberShop/internal/app/fiber"
	"FiberShop/internal/database/postgres"
	"FiberShop/internal/database/redis"
	"FiberShop/internal/transport/grpc/auth"
	"go.uber.org/zap"
)

type App struct {
	FiberApp *fiberapp.App
}

func New(log *zap.Logger, authClient *auth.Client) *App {
	db, err := postgres.New()
	if err != nil {
		panic(err)
	}
	cache := redis.New(log, db)
	fiberApp := fiberapp.New(log, authClient, db, cache)
	return &App{
		FiberApp: fiberApp,
	}
}
