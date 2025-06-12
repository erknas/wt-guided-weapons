package service

import (
	"context"
	"log/slog"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

type WeaponsInserter interface {
	Insert(context.Context, *types.WeaponParams) error
}

type WeaponsUpdater interface {
	Update(context.Context, *types.WeaponParams) error
}

type WeaponsProvider interface {
	Provide(context.Context, string) ([]*types.WeaponParams, error)
}

type Service struct {
	inserter WeaponsInserter
	updater  WeaponsUpdater
	provider WeaponsProvider
	log      *slog.Logger
}

func New(inserter WeaponsInserter, updater WeaponsUpdater, provider WeaponsProvider, log *slog.Logger) *Service {
	return &Service{
		inserter: inserter,
		updater:  updater,
		provider: provider,
		log:      log,
	}
}

func (s *Service) InsertWeapons(ctx context.Context, params *types.WeaponParams) error {
	return nil
}

func (s *Service) UpdateWeapons(ctx context.Context, params *types.WeaponParams) error {
	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.WeaponParams, error) {
	return nil, nil
}
