package service

import (
	"context"
	"fmt"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type WeaponsInserter interface {
	Insert(ctx context.Context, weapons []*types.Weapon) error
}

type WeaponsProvider interface {
	ByCategory(ctx context.Context, category string) ([]*types.Weapon, error)
	Search(ctx context.Context, query string) ([]types.SearchResult, error)
}

type WeaponsAggregator interface {
	Aggregate(ctx context.Context) ([]*types.Weapon, error)
}

type Service struct {
	inserter   WeaponsInserter
	provider   WeaponsProvider
	aggregator WeaponsAggregator
}

func New(inserter WeaponsInserter, provider WeaponsProvider, aggregator WeaponsAggregator) *Service {
	return &Service{
		inserter:   inserter,
		provider:   provider,
		aggregator: aggregator,
	}
}

func (s *Service) InsertWeapons(ctx context.Context) error {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.aggregator.Aggregate(ctx)
	if err != nil {
		log.Error("failed to aggregate weapons",
			zap.Error(err),
		)
		return fmt.Errorf("failed to aggregate weapons: %w", err)
	}

	if err := s.inserter.Insert(ctx, weapons); err != nil {
		log.Error("failed to insert weapons",
			zap.Error(err),
		)
		return err
	}

	log.Debug("InsertWeapons complited")

	return nil
}

func (s *Service) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.ByCategory(ctx, category)
	if err != nil {
		log.Error("failed to provide weapons",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, err
	}

	log.Debug("WeaponsByCategory complited",
		zap.String("category", category),
		zap.Int("total weapons", len(weapons)),
	)

	return weapons, nil
}

func (s *Service) SearchWeapon(ctx context.Context, query string) ([]types.SearchResult, error) {
	log := logger.FromContext(ctx, logger.Service)

	results, err := s.provider.Search(ctx, query)
	if err != nil {
		log.Error("failed to find weapon",
			zap.Error(err),
		)
		return nil, err
	}

	log.Debug("SearchWeapon complited",
		zap.String("query", query),
		zap.Int("total overlaps", len(results)),
	)

	return results, nil
}
