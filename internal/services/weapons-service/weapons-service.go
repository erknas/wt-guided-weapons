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
	weaponsUpserter   WeaponsUpserter
	weaponsProvider   WeaponsProvider
	weaponsAggregator WeaponsAggregator
	versionUpdater    VersionUpdater
}

func New(weaponsUpserter WeaponsUpserter, weaponsProvider WeaponsProvider, weaponsAggregator WeaponsAggregator, versionUpdater VersionUpdater) *WeaponsService {
	return &WeaponsService{
		weaponsUpserter:   weaponsUpserter,
		weaponsProvider:   weaponsProvider,
		weaponsAggregator: weaponsAggregator,
		versionUpdater:    versionUpdater,
	}
}

func (s *WeaponsService) UpdateWeapons(ctx context.Context) error {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.weaponsAggregator.AggregateWeapons(ctx)
	if err != nil {
		log.Error("AggregateWeapons failed",
			zap.Error(err),
		)
		return fmt.Errorf("failed to aggregate weapons: %w", err)
	}

	if err := s.weaponsUpserter.UpsertWeapons(ctx, weapons); err != nil {
		log.Error("DB call UpsertWeapons failed",
			zap.Error(err),
		)
		return err
	}

	if err := s.versionUpdater.UpdateVersion(ctx); err != nil {
		log.Error("Service call UpdateVersion failed",
			zap.Error(err),
		)
		return err
	}

	log.Debug("Service UpdateWeapons complited")

	return nil
}

func (s *WeaponsService) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.weaponsProvider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("DB call WeaponsByCategory failed",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, err
	}

	log.Debug("Service GetWeaponsByCategory complited",
		zap.String("category", category),
		zap.Int("total weapons", len(weapons)),
	)

	return weapons, nil
}

func (s *WeaponsService) SearchWeapons(ctx context.Context, query string) ([]types.SearchResult, error) {
	log := logger.FromContext(ctx, logger.Service)

	results, err := s.weaponsProvider.WeaponsByName(ctx, query)
	if err != nil {
		log.Error("DB call WeaponsByName failed",
			zap.Error(err),
		)
		return nil, err
	}

	log.Debug("Service SearchWeapons complited",
		zap.String("query", query),
		zap.Int("total overlaps", len(results)),
	)

	return results, nil
}
