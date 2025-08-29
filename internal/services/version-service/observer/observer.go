package observer

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/storage/mongodb"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

const period = time.Minute * 30

type version struct {
	version string
	err     error
}

type VersionProvider interface {
	GetVersion(ctx context.Context) (types.LastChange, error)
}

type VersionParser interface {
	Parse(ctx context.Context, url string) (types.VersionInfo, error)
}

type WeaponsUpdater interface {
	UpdateWeapons(ctx context.Context) error
}

type ChangeObserver struct {
	provider VersionProvider
	parser   VersionParser
	updater  WeaponsUpdater
	log      *zap.Logger
	url      string
}

func New(
	provider VersionProvider,
	parser VersionParser,
	updater WeaponsUpdater,
	log *zap.Logger,
	url string,
) *ChangeObserver {
	return &ChangeObserver{
		provider: provider,
		parser:   parser,
		updater:  updater,
		log:      log,
		url:      url,
	}
}

func (o *ChangeObserver) Observe(ctx context.Context) {
	_, err := o.provider.GetVersion(ctx)
	if err != nil && errors.Is(err, mongodb.ErrNoVersion) {
		o.log.Info("Inserting initial data")
		err := o.updater.UpdateWeapons(ctx)
		if err != nil {
			o.log.Error("Failed to insert initial data",
				zap.Error(err),
			)
		}
	}

	ticker := time.NewTicker(period)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := o.checkVersionChange(ctx); err != nil {
				o.log.Error("checkVersionChage error",
					zap.Error(err),
				)
			}
		case <-ctx.Done():
		}
	}
}

func (o *ChangeObserver) checkVersionChange(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	var currVersion, newVerison version

	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func() {
		defer wg.Done()

		ver, err := o.provider.GetVersion(ctx)
		currVersion = version{version: ver.Version.Version, err: err}
	}()

	go func() {
		defer wg.Done()

		ver, err := o.parser.Parse(ctx, o.url)
		newVerison = version{version: ver.Version, err: err}
	}()

	wg.Wait()

	if currVersion.err != nil {
		o.log.Warn("GetVersion error",
			zap.Error(currVersion.err),
		)
	}

	if newVerison.err != nil {
		o.log.Error("Parse error",
			zap.Error(newVerison.err),
		)
		return fmt.Errorf("failed to get new version: %w", newVerison.err)
	}

	o.log.Debug("versions",
		zap.String("current version", currVersion.version),
		zap.String("new version", newVerison.version),
	)

	if currVersion.version != newVerison.version {
		if err := o.updater.UpdateWeapons(ctx); err != nil {
			o.log.Error("UpdateWeapons error",
				zap.Error(err),
			)
			return fmt.Errorf("failed to update weapons")
		}
		o.log.Info("Weapons updated",
			zap.String("version changed", fmt.Sprintf("%s -> %s", currVersion.version, newVerison.version)),
		)
	}

	o.log.Info("Nothing to update")

	return nil
}
