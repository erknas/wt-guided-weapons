package versionservice

import (
	"context"
	"fmt"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type VersionUpserter interface {
	UpsertVersion(ctx context.Context, version types.VersionInfo) error
}

type VersionProvider interface {
	Version(ctx context.Context) (types.LastChange, error)
}

type VersionParser interface {
	Parse(ctx context.Context, url string) (types.VersionInfo, error)
}

type VersionService struct {
	versionUpdater  VersionUpserter
	versionParser   VersionParser
	versionProvider VersionProvider
	url             string
}

func New(
	versionUpserter VersionUpserter,
	versionProvider VersionProvider,
	versionParser VersionParser,
	url string,
) *VersionService {
	return &VersionService{
		versionUpdater:  versionUpserter,
		versionParser:   versionParser,
		versionProvider: versionProvider,
		url:             url,
	}
}

func (s *VersionService) UpdateVersion(ctx context.Context) error {
	log := logger.FromContext(ctx, logger.Service)

	version, err := s.versionParser.Parse(ctx, s.url)
	if err != nil {
		log.Error("Parse version failed",
			zap.Error(err),
		)
		return fmt.Errorf("failed to parse version: %w", err)
	}

	if err := s.versionUpdater.UpsertVersion(ctx, version); err != nil {
		log.Error("DB call UpsertVersion failed",
			zap.Error(err),
		)
		return fmt.Errorf("failed to update version: %w", err)
	}

	log.Debug("Service UpdateVersion complited",
		zap.String("new version", version.Version),
	)

	return nil
}

func (s *VersionService) GetVersion(ctx context.Context) (types.LastChange, error) {
	log := logger.FromContext(ctx, logger.Service)

	version, err := s.versionProvider.Version(ctx)
	if err != nil {
		log.Error("DB call Version failed",
			zap.Error(err),
		)
		return types.LastChange{}, fmt.Errorf("failed to get version: %w", err)
	}

	log.Debug("Service GetVersion complited",
		zap.Any("version", version),
	)

	return version, nil
}
