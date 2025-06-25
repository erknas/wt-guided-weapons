package service

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

var (
	ErrCategoryNotExists = errors.New("category does not exist")
)

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
}

type TableParser interface {
	Parse(context.Context, string, string) ([]*types.Weapon, error)
}

type Service struct {
	inserter WeaponsInserter
	provider WeaponsProvider
	parser   TableParser
	urls     map[string]string
	log      *zap.Logger
}

func New(
	inserter WeaponsInserter,
	provider WeaponsProvider,
	parser TableParser,
	urls map[string]string,
	log *zap.Logger,
) *Service {

	return &Service{
		inserter: inserter,
		provider: provider,
		parser:   parser,
		urls:     urls,
		log:      log,
	}
}

func (s *Service) InsertWeapons(ctx context.Context) error {
	start := time.Now()
	log := logger.FromContext(ctx, logger.Service)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &sync.WaitGroup{}
	errCh := make(chan error, len(s.urls))
	dataCh := make(chan []*types.Weapon, len(s.urls))

	for category, url := range s.urls {
		wg.Add(1)
		go func(category, url string) {
			defer wg.Done()

			if ctx.Err() != nil {
				return
			}

			data, err := s.parser.Parse(ctx, category, url)
			if err != nil {
				select {
				case errCh <- err:
					log.Error("parse table failed",
						zap.Error(err),
						zap.String("category", category),
						zap.String("table_url", url),
					)
				case <-ctx.Done():
				}
				return
			}

			select {
			case dataCh <- data:
				log.Debug("parse table completed",
					zap.String("category", category),
					zap.String("table_url", url),
					zap.Int("weapons_count", len(data)),
				)
			case <-ctx.Done():
			}

		}(category, url)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(dataCh)
		log.Debug("all goroutines complited")
	}()

	var weapons []*types.Weapon
	successfulTables := 0

	for range s.urls {
		select {
		case <-ctx.Done():
			log.Warn("context cancelled",
				zap.Error(ctx.Err()),
				zap.Int("complited tables", successfulTables),
			)
			return ctx.Err()
		case err := <-errCh:
			cancel()
			return err
		case data := <-dataCh:
			weapons = append(weapons, data...)
			successfulTables++
		}
	}

	log.Info("tabels parsing complited",
		zap.Int("total successful tables parsed", successfulTables),
		zap.Int("total failed tables", len(s.urls)-successfulTables),
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

	if _, exists := s.urls[category]; !exists {
		log.Warn("category does not exist",
			zap.String("category", category),
		)
		return nil, ErrCategoryNotExists
	}

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
