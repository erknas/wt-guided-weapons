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
	upserter VersionUpserter
	provider VersionProvider
	parser   VersionParser
	url      string
}

func New(
	upserter VersionUpserter,
	provider VersionProvider,
	parser VersionParser,
	url string,
) *VersionService {
	return &VersionService{
		upserter: upserter,
		parser:   parser,
		provider: provider,
		url:      url,
	}
}

func (s *VersionService) UpdateVersion(ctx context.Context) error {
	log := logger.FromContext(ctx, logger.Service)

	version, err := s.parser.Parse(ctx, s.url)
	if err != nil {
		log.Error("Parse error",
			zap.Error(err),
		)
		return fmt.Errorf("failed to parse version: %w", err)
	}

	if err := s.upserter.UpsertVersion(ctx, version); err != nil {
		log.Error("UpsertVersion error",
			zap.Error(err),
		)
		return fmt.Errorf("failed to update version: %w", err)
	}

	log.Debug("UpdateVersion complited",
		zap.String("new version", version.Version),
	)

	return nil
}

func (s *VersionService) GetVersion(ctx context.Context) (types.LastChange, error) {
	log := logger.FromContext(ctx, logger.Service)

	version, err := s.provider.Version(ctx)
	if err != nil {
		log.Error("Version error",
			zap.Error(err),
		)
		return types.LastChange{}, fmt.Errorf("failed to get version: %w", err)
	}

	log.Debug("GetVersion complited",
		zap.Any("version", version),
	)

	return version, nil
}
