package observer

import (
	"context"
	"fmt"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

const period = time.Second * 30

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
	provider      VersionProvider
	parser        VersionParser
	updater       WeaponsUpdater
	log           *zap.Logger
	url           string
	currVersionCh chan version
	newVersionCh  chan version
}

func New(
	provider VersionProvider,
	parser VersionParser,
	updater WeaponsUpdater,
	log *zap.Logger,
	url string,
) *ChangeObserver {
	return &ChangeObserver{
		provider:      provider,
		parser:        parser,
		updater:       updater,
		log:           log,
		url:           url,
		currVersionCh: make(chan version, 1),
		newVersionCh:  make(chan version, 1),
	}
}

func (o *ChangeObserver) Observe(ctx context.Context) {
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

	go func() {
		currVersion, err := o.provider.GetVersion(ctx)
		select {
		case o.currVersionCh <- version{version: currVersion.Version.Version, err: err}:
		case <-ctx.Done():
		}
	}()

	go func() {
		newVersion, err := o.parser.Parse(ctx, o.url)
		select {
		case o.newVersionCh <- version{version: newVersion.Version, err: err}:
		case <-ctx.Done():
		}
	}()

	var currVersion, newVerison version

	for i := 0; i < 2; i++ {
		select {
		case currVersion = <-o.currVersionCh:
		case newVerison = <-o.newVersionCh:
		case <-ctx.Done():
			o.log.Warn("context done while receiving data",
				zap.Error(ctx.Err()),
			)
			return ctx.Err()
		}
	}

	if currVersion.err != nil {
		o.log.Error("GetVersion error",
			zap.Error(currVersion.err),
		)
		return fmt.Errorf("failed to get current version: %w", currVersion.err)
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
		}
		o.log.Info("Weapons updated",
			zap.String("version changed", fmt.Sprintf("%s -> %s", currVersion.version, newVerison.version)),
		)
	}

	o.log.Info("Nothing to update")

	return nil
}
