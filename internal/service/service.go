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
	WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error)
	Search(ctx context.Context, name string) (map[string]string, error)
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
		return fmt.Errorf("failed to insert weapons: %w", err)
	}

	log.Debug("InsertWeapons complited")

	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("failed to provide weapons",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, fmt.Errorf("failed to get weapons by category %s: %w", category, err)
	}

	log.Debug("GetWeaponsByCategory complited",
		zap.String("category", category),
		zap.Int("total weapons", len(weapons)),
	)

	return weapons, nil
}

func (s *Service) SearchWeapon(ctx context.Context, name string) (map[string]string, error) {
	log := logger.FromContext(ctx, logger.Service)

	result, err := s.provider.Search(ctx, name)
	if err != nil {
		log.Error("failed to find weapon",
			zap.Error(err),
		)
		return nil, fmt.Errorf("failed to find weapon %s: %w", name, err)
	}

	return result, nil
}
