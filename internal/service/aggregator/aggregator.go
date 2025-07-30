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

type parseJob struct {
	category string
	url      string
}

type parseResult struct {
	weapons  []*types.Weapon
	category string
	err      error
}

const numWorkers = 4

func (w *Weapons) Aggregate(ctx context.Context) ([]*types.Weapon, error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	jobs := make(chan parseJob, len(w.urls))
	results := make(chan parseResult, len(w.urls))

	wg := sync.WaitGroup{}

	start := time.Now()
	for i := range numWorkers {
		wg.Add(1)
		w.log.Debug("starting", zap.Int("worker №", i+1))
		go w.worker(ctx, jobs, results, &wg)
	}

	go func() {
		defer close(jobs)

		for category, url := range w.urls {
			select {
			case jobs <- parseJob{category: category, url: url}:
				w.log.Debug("sending job",
					zap.String("table", fmt.Sprintf("%s | %s", category, url)),
				)
			case <-ctx.Done():
				w.log.Error("context cancelled on sending job")
				return
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	var weapons []*types.Weapon
	tables := 0

	for result := range results {
		if result.err != nil {
			w.log.Error("failed to parse table",
				zap.Error(result.err),
				zap.String("category", result.category),
			)
			cancel()
			return nil, fmt.Errorf("failed to parse table: %w", result.err)
		}

		w.log.Info("table successfully parsed",
			zap.String("category", result.category),
		)
		weapons = append(weapons, result.weapons...)
		tables++
	}

	w.log.Info("tables parsing complited",
		zap.Int("total tables", tables),
		zap.Int("total weapons", len(weapons)),
		zap.Duration("took time", time.Since(start)),
	)

	return weapons, nil
}

func (w *Weapons) worker(ctx context.Context, jobs <-chan parseJob, results chan<- parseResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		weapons, err := w.parser.Parse(ctx, job.category, job.url)
		select {
		case results <- parseResult{
			weapons:  weapons,
			category: job.category,
			err:      err,
		}:
		case <-ctx.Done():
			w.log.Error("context cancelled on sending results")
			return
		}
	}
}
