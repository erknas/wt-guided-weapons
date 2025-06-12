package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/server"
	"github.com/erknas/wt-guided-weapons/internal/service"
	"github.com/erknas/wt-guided-weapons/internal/storage"
)

func main() {
	var (
		cfg = config.Load()
		log = logger.New(cfg.Env)
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	storage, err := storage.New(ctx, cfg)
	if err != nil {
		log.Error("failed to init storage", "error", err)
		os.Exit(1)
	}

	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := storage.Close(closeCtx); err != nil {
			log.Error("failed to disconnect from storage", "error", err)
		}
	}()

	service := service.New(storage, storage, storage, log)

	server := server.New(service)
	server.Run(ctx, cfg)
}
