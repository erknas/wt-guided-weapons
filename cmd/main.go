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
	"github.com/erknas/wt-guided-weapons/internal/services/version-service/observer"
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

	logger.Info("Config loaded",
		zap.Any("cfg", cfg),
	)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	mongodb, err := mongodb.New(ctx, cfg)
	if err != nil {
		logger.Error("Failed to initialze mongodb",
			zap.Error(err),
		)
		os.Exit(1)
	}

	logger.Info("Mongodb initialized")

	defer func() {
		closeCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		defer cancel()

		if err := mongodb.Close(closeCtx); err != nil {
			logger.Warn("failed to disconnect from storage",
				zap.Error(err),
			)
		}
		logger.Info("mognodb closed")
	}()

	urls, err := urlsloader.Load(cfg.URLs)
	if err != nil {
		logger.Error("Failed to load urls",
			zap.Error(err),
		)
		os.Exit(1)
	}

	reader := csvreader.New()

	versionParser := versionparser.New(reader)
	versionService := versionservice.New(mongodb, mongodb, versionParser, urls["version"])

	weaponsParser := weaponsparser.New(reader, &weaponmapper.WeaponMapper{})
	weaponsAggregator := weaponsaggregator.New(urls, weaponsParser, logger)
	weaponsService := weaponsservice.New(mongodb, mongodb, weaponsAggregator, versionService)

	if err := weaponsService.UpdateWeapons(ctx); err != nil {
		logger.Error("Failed to insert initial data",
			zap.Error(err),
		)
		os.Exit(1)
	}

	observer := observer.New(versionService, versionParser, weaponsService, logger, urls["version"])
	go observer.Observe(ctx)

	server := server.New(weaponsService, versionService, urls, logger)
	if err := server.Run(ctx, cfg); err != nil {
		logger.Error("Server error",
			zap.Error(err),
		)
		os.Exit(1)
	}
}
