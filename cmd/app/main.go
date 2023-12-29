package main

import (
	"FiberShop/internal/app"
	"FiberShop/internal/config"
	"FiberShop/internal/lib/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)
	log.Info("starting application", zap.Any("config", cfg))
	application := app.New(log)
	go application.FiberApp.Run()
	//application := app.New(log, cfg.GRPC.Port, cfg.TokenTtl)
	//
	//go application.GRPCSrv.MustRun()
	//
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	log.Info("stopping application", zap.String("signal", sign.String()))
	application.FiberApp.Stop()
	log.Info("application stopped")
}
