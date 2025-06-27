package service

import (
	"context"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
}

type WeaponsAggregator interface {
	Aggregate(context.Context) ([]*types.Weapon, error)
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
		return err
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

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("failed to provide weapons",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, err
	}

	log.Debug("GetWeaponsByCategory complited",
		zap.String("category", category),
		zap.Int("total weapons", len(weapons)),
	)

	return weapons, nil
}
