package app

import (
	fiberapp "FiberShop/internal/app/fiber"
	"FiberShop/internal/database/postgres"
	"go.uber.org/zap"
)

type App struct {
	FiberApp *fiberapp.App
}

func New(log *zap.Logger) *App {
	db, err := postgres.New()
	print(db)
	if err != nil {
		panic(err)
	}
	fiberApp := fiberapp.New(log)
	return &App{
		FiberApp: fiberApp,
	}
}
