package main

import (
	"FiberShop/internal/app"
	"FiberShop/internal/config"
	"FiberShop/internal/transport/grpc"
	"os"
	"os/signal"
	"syscall"
	"time"

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
	grpc, err := grpc.New(cfg)
	if err != nil {
		log.Error("failed to init grpc clients", zap.Error(err))
		os.Exit(1)
	}
	application := app.New(log, grpc)
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
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	logger, _ := cfg.Build()
	defer logger.Sync()
	return logger
}

//todo microservices
//todo fix google auth
//TODO оптимизировать(горутины, скорость и т.д)
//TODO добавить ссылку для восстановления пароля
