package csvparser

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

type Weapons struct {
	urls map[string]string
	log  *zap.Logger
}

func New(urls map[string]string, log *zap.Logger) *Weapons {
	return &Weapons{
		urls: urls,
		log:  log,
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

			data, err := w.ParseTableByURL(ctx, category, url)
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
			return nil, ctx.Err()
		case err := <-errCh:
			cancel()
			return nil, err
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

func (w *Weapons) ParseTableByURL(ctx context.Context, category, url string) ([]*types.Weapon, error) {
	data, err := w.readCSV(ctx, url)
	if err != nil {
		return nil, err
	}

	var weapons []*types.Weapon

	for i := range data[0][1:] {
		weapon, err := mapCSVToStruct(data, category, i+1)
		if err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

func (w *Weapons) readCSV(ctx context.Context, url string) ([][]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to table [%s]: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return csv.NewReader(resp.Body).ReadAll()
}
