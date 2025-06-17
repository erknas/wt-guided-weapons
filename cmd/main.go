package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/server"
	"github.com/erknas/wt-guided-weapons/internal/service"
	"github.com/erknas/wt-guided-weapons/internal/storage"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger, err := logger.New(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	storage, err := storage.New(ctx, cfg)
	if err != nil {
		logger.Error("failed to init storage", zap.Error(err))
		os.Exit(1)
	}

	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := storage.Close(closeCtx); err != nil {
			logger.Error("failed to disconnect from storage", zap.Error(err))
		}
	}()

	service := service.New(storage, storage, storage, logger)

	server := server.New(service, logger)

	if err := server.Run(ctx, cfg); err != nil {
		logger.Error("server failed", zap.Error(err))
		os.Exit(1)
	}
}
