package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	csvreader "github.com/erknas/wt-guided-weapons/internal/lib/csv-reader"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/server"
	versionservice "github.com/erknas/wt-guided-weapons/internal/services/version-service"
	versionparser "github.com/erknas/wt-guided-weapons/internal/services/version-service/version-parser"
	weaponsservice "github.com/erknas/wt-guided-weapons/internal/services/weapons-service"
	weaponmapper "github.com/erknas/wt-guided-weapons/internal/services/weapons-service/weapon-mapper"
	weaponsparser "github.com/erknas/wt-guided-weapons/internal/services/weapons-service/weapon-parser"
	weaponsaggregator "github.com/erknas/wt-guided-weapons/internal/services/weapons-service/weapons-aggregator"
	"github.com/erknas/wt-guided-weapons/internal/storage/mongodb"
	urlsloader "github.com/erknas/wt-guided-weapons/internal/urls-loader"
	"go.uber.org/zap"
)

func main() {
	configPath := flag.String("config", "local.yaml", "path to the config")
	flag.Parse()

	cfg := config.MustLoad(*configPath)

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

	urls, err := urlsloader.Load(cfg.URLs)
	if err != nil {
		logger.Error("failed to load urls",
			zap.Error(err),
		)
		os.Exit(1)
	}

	logger.Info("urls loaded",
		zap.Int("total", len(urls)),
	)

	reader := csvreader.New()

	versionParser := versionparser.New(reader)
	versionService := versionservice.New(storage, storage, versionParser, urls["version"])

	weaponsParser := weaponsparser.New(reader, &weaponmapper.WeaponMapper{})
	weaponsAggregator := weaponsaggregator.New(urls, weaponsParser, logger)
	weaponsService := weaponsservice.New(storage, storage, weaponsAggregator, versionService)

	server := server.New(weaponsService, versionService, urls, logger)
	if err := server.Run(ctx, cfg); err != nil {
		logger.Error("server failed",
			zap.Error(err),
		)
		os.Exit(1)
	}
}
