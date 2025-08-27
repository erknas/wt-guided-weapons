package weaponsservice

import (
	"context"
	"fmt"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type WeaponsUpserter interface {
	UpsertWeapons(ctx context.Context, weapons []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error)
	WeaponsByName(ctx context.Context, query string) ([]types.SearchResult, error)
}

type WeaponsAggregator interface {
	AggregateWeapons(ctx context.Context) ([]*types.Weapon, error)
}

type VersionUpdater interface {
	UpdateVersion(ctx context.Context) error
}

type WeaponsService struct {
	upserter   WeaponsUpserter
	provider   WeaponsProvider
	aggregator WeaponsAggregator
	updater    VersionUpdater
}

func New(
	upserter WeaponsUpserter,
	provider WeaponsProvider,
	aggregator WeaponsAggregator,
	updater VersionUpdater,
) *WeaponsService {
	return &WeaponsService{
		upserter:   upserter,
		provider:   provider,
		aggregator: aggregator,
		updater:    updater,
	}
}

func (s *WeaponsService) UpdateWeapons(ctx context.Context) error {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.aggregator.AggregateWeapons(ctx)
	if err != nil {
		log.Error("AggregateWeapons error",
			zap.Error(err),
		)
		return fmt.Errorf("failed to aggregate weapons: %w", err)
	}

	if err := s.upserter.UpsertWeapons(ctx, weapons); err != nil {
		log.Error("UpsertWeapons error",
			zap.Error(err),
		)
		return err
	}

	if err := s.updater.UpdateVersion(ctx); err != nil {
		log.Error("UpdateVersion error",
			zap.Error(err),
		)
		return err
	}

	log.Debug("UpdateWeapons complited")

	return nil
}

func (s *WeaponsService) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("WeaponsByCategory error",
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

func (s *WeaponsService) SearchWeapons(ctx context.Context, query string) ([]types.SearchResult, error) {
	log := logger.FromContext(ctx, logger.Service)

	results, err := s.provider.WeaponsByName(ctx, query)
	if err != nil {
		log.Error("WeaponsByName error",
			zap.Error(err),
		)
		return nil, err
	}

	log.Debug("SearchWeapons complited",
		zap.String("query", query),
		zap.Int("total overlaps", len(results)),
	)

	return results, nil
}
