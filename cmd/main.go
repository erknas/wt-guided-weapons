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
	"github.com/erknas/wt-guided-weapons/internal/service/aggregator"
	csvreader "github.com/erknas/wt-guided-weapons/internal/service/csv-reader"
	urlsloader "github.com/erknas/wt-guided-weapons/internal/service/urls-loader"
	weaponmapper "github.com/erknas/wt-guided-weapons/internal/service/weapon-mapper"
	weaponparser "github.com/erknas/wt-guided-weapons/internal/service/weapon-parser"
	"github.com/erknas/wt-guided-weapons/internal/storage/mongodb"
	"go.uber.org/zap"
)

func main() {
	cfg := config.Load()

	logger, err := logger.New(cfg.Env)
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("config loaded",
		zap.Any("cfg", cfg),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	storage, err := mongodb.New(ctx, cfg)
	if err != nil {
		logger.Error("failed to init storage",
			zap.Error(err),
		)
		os.Exit(1)
	}

	logger.Info("storage init")

	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := storage.Close(closeCtx); err != nil {
			logger.Warn("failed to disconnect from storage",
				zap.Error(err),
			)
		}
		logger.Info("storage closed")
	}()

	urls, err := urlsloader.Load(cfg.FileName)
	if err != nil {
		logger.Error("failed to load urls",
			zap.Error(err),
		)
		os.Exit(1)
	}

	parser := weaponparser.New(&csvreader.HTTPReader{}, &weaponmapper.WeaponMapper{})

	aggregator := aggregator.New(urls, parser, logger)

	service := service.New(storage, storage, aggregator)

	server := server.New(service, urls, logger)
	if err := server.Run(ctx, cfg); err != nil {
		logger.Error("server failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
}
