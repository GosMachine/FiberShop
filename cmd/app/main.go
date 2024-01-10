package main

import (
	"FiberShop/internal/app"
	"FiberShop/internal/config"
	"FiberShop/internal/lib/logger"
	"FiberShop/internal/transport/grpc/auth"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
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

//TODO оптимизировать(горутины, скорость и т.д)
//TODO обрабатывать красивее ошибки на фронте
//TODO отслеживать просмотры страницы и ip
//TODO оптимизировать сервис аутх(кэш и т.д)
//TODO попробовать сделать некторые посты запросы без перезагрузки страницы
