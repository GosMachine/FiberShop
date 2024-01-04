package middleware

import "go.uber.org/zap"

type App struct {
	log *zap.Logger
}

func New(logger *zap.Logger) *App {
	return &App{log: logger}
}
