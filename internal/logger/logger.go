package logger

import (
	"context"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLocal = "local"
	envProd  = "production"
)

func New(env string) (*zap.Logger, error) {
	var logger *zap.Logger
	var err error

	switch env {
	case envLocal:
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, err = cfg.Build(
			zap.AddCaller(),
			zap.AddStacktrace(zap.ErrorLevel),
		)
	case envProd:
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = "timestamp"
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		logger, err = cfg.Build(
			zap.AddCaller(),
			zap.AddStacktrace(zap.WarnLevel),
		)
	default:
		logger, err = zap.NewProduction()
	}

	if err != nil {
		return nil, err
	}

	zap.ReplaceGlobals(logger)

	return logger, nil
}

func FromContext(ctx context.Context, layer string) *zap.Logger {
	if logger, ok := ctx.Value("logger").(*zap.Logger); ok {
		return logger.With(zap.String("layer", layer))
	}

	return zap.L().With(zap.String("layer", layer))
}
