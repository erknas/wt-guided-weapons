package service

import (
	"context"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	csvparser "github.com/erknas/wt-guided-weapons/internal/service/csv-parser"
	"github.com/erknas/wt-guided-weapons/internal/service/tables"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
}

type Service struct {
	inserter WeaponsInserter
	provider WeaponsProvider
	log      *zap.Logger
}

func New(inserter WeaponsInserter, provider WeaponsProvider, log *zap.Logger) *Service {
	return &Service{
		inserter: inserter,
		provider: provider,
		log:      log,
	}
}

func (s *Service) InsertWeapons(ctx context.Context) error {
	start := time.Now()
	log := logger.FromContext(ctx, logger.Service)

	wg := &sync.WaitGroup{}

	tables, err := tables.Load()
	if err != nil {
		return err
	}

	errCh := make(chan error, len(tables.Tables))
	dataCh := make(chan []*types.Weapon, len(tables.Tables))

	for category, url := range tables.Tables {
		wg.Add(1)
		go func(category, url string) {
			defer wg.Done()

			data, err := csvparser.ParseTable(ctx, category, url)
			if err != nil {
				log.Error("parse table failed",
					zap.Error(err),
					zap.String("category", category),
					zap.String("table_url", url),
				)
				errCh <- err
				return
			}

			log.Debug("parse table complited",
				zap.String("category", category),
				zap.String("table_url", url),
				zap.Int("weapons_count", len(data)),
			)

			dataCh <- data
		}(category, url)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(dataCh)
		log.Debug("channels closed")
		log.Debug("all goroutine complited")
	}()

	var weapons []*types.Weapon
	successfulTables := 0

	for range tables.Tables {
		select {
		case <-ctx.Done():
			log.Warn("parsing cancelled",
				zap.Error(ctx.Err()),
			)
			return ctx.Err()
		case err := <-errCh:
			if err != nil {
				return err
			}
		case data := <-dataCh:
			weapons = append(weapons, data...)
			successfulTables++
		}
	}

	log.Info("tabels parsing complited",
		zap.Int("total successful tables parsed", successfulTables),
		zap.Int("total failed tables", len(tables.Tables)-successfulTables),
		zap.Int("weapons_count", len(weapons)),
		zap.Duration("duration", time.Since(start)),
	)

	if err := s.inserter.Insert(ctx, weapons); err != nil {
		log.Error("insert operation failed",
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("service call failed",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, err
	}

	log.Debug("service call complited",
		zap.String("category", category),
		zap.Int("weapons_count", len(weapons)),
	)

	return weapons, nil
}
