package aggregator

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type WeaponParser interface {
	Parse(ctx context.Context, category string, url string) ([]*types.Weapon, error)
}

type Weapons struct {
	urls   map[string]string
	parser WeaponParser
	log    *zap.Logger
}

func New(urls map[string]string, parser WeaponParser, log *zap.Logger) *Weapons {
	return &Weapons{
		urls:   urls,
		parser: parser,
		log:    log,
	}
}

func (w *Weapons) Aggregate(ctx context.Context) ([]*types.Weapon, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := &sync.WaitGroup{}

	errCh := make(chan error, len(w.urls))
	dataCh := make(chan []*types.Weapon, len(w.urls))

	start := time.Now()
	for category, url := range w.urls {
		wg.Add(1)
		go func(category, url string) {
			defer wg.Done()

			if ctx.Err() != nil {
				return
			}

			data, err := w.parser.Parse(ctx, category, url)
			if err != nil {
				select {
				case errCh <- err:
					w.log.Error("parse table failed",
						zap.Error(err),
						zap.String("category", category),
						zap.String("table_url", url),
					)
				case <-ctx.Done():
				}
				return
			}

			select {
			case dataCh <- data:
				w.log.Debug("parse table completed",
					zap.String("category", category),
					zap.String("table_url", url),
					zap.Int("weapons_count", len(data)),
				)
			case <-ctx.Done():
			}

		}(category, url)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(dataCh)
		w.log.Debug("all goroutines complited")
	}()

	var weapons []*types.Weapon
	successfulTables := 0

	for range w.urls {
		select {
		case <-ctx.Done():
			w.log.Warn("context cancelled",
				zap.Error(ctx.Err()),
				zap.Int("complited tables", successfulTables),
			)
			return nil, fmt.Errorf("context error: %w", ctx.Err())
		case err := <-errCh:
			cancel()
			return nil, fmt.Errorf("failed to parse table: %w", err)
		case data := <-dataCh:
			weapons = append(weapons, data...)
			successfulTables++
		}
	}

	w.log.Info("tabels parsing complited",
		zap.Int("total tables parsed", successfulTables),
		zap.Int("weapons_count", len(weapons)),
		zap.Duration("duration", time.Since(start)),
	)

	return weapons, nil
}
