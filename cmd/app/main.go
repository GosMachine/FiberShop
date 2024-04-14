package main

import (
	"FiberShop/internal/app"
	"FiberShop/internal/config"
	"FiberShop/internal/transport/grpc/auth"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	cfg := config.MustLoad()
	log := setupLogger()
	log.Info("starting application", zap.Any("config", cfg))
	authClient, err := auth.New(context.Background(), cfg.Clients.Auth.Address,
		cfg.Clients.Auth.Timeout, cfg.Clients.Auth.RetriesCount)
	if err != nil {
		log.Error("failed to init auth client", zap.Error(err))
		os.Exit(1)
	}
	application := app.New(log, authClient)
	go application.FiberApp.Run(cfg.Address)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	sign := <-stop
	log.Info("stopping application", zap.String("signal", sign.String()))
	application.FiberApp.Stop()
	log.Info("application stopped")
}

func setupLogger() *zap.Logger {
	cfg := zap.Config{
		Encoding:          "json",
		DisableStacktrace: true,
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:       []string{"stdout"},
		EncoderConfig:     zap.NewProductionEncoderConfig(),
	}
	logger, _ := cfg.Build()
	defer logger.Sync()
	return logger
}

//todo microservices
//TODO оптимизировать(горутины, скорость и т.д)
//TODO оптимизировать сервис аутх(кэш и т.д)
//TODO добавить ссылку для восстановления пароля
