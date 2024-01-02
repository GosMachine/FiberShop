package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(env string) *zap.Logger {
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	outputPaths := []string{"stdout"}
	if env == "prod" {
		//outputPaths = []string{"logs/logfile.txt"}
		level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}
	cfg := zap.Config{
		Encoding:          "json",
		DisableStacktrace: true,
		Level:             level,
		OutputPaths:       outputPaths,
		EncoderConfig:     zap.NewProductionEncoderConfig(),
	}
	logger, _ := cfg.Build()
	defer logger.Sync()
	return logger
}
