package app

import (
	fiberapp "FiberShop/internal/app/fiber"
	"FiberShop/internal/database/postgres"
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
	fiberApp := fiberapp.New(log, authClient, db)
	return &App{
		FiberApp: fiberApp,
	}
}
