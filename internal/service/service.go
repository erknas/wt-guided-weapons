package service

import (
	"context"
	"log/slog"
	"sync"

	"github.com/erknas/wt-guided-weapons/internal/logger/sl"
	csvparser "github.com/erknas/wt-guided-weapons/internal/service/csv-parser"
	"github.com/erknas/wt-guided-weapons/internal/types"
)

var tables map[string]string = map[string]string{
	"aam_ir_rear_aspect": "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=0",
	"aam_ir_all_aspect":  "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1726112384",
	"aam_ir_heli":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=638442034",
	"aam_sarh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=128448244",
	"aam_arh":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=650249168",
	"aam_manual":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=29789551",
	"agm_tv":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1614911062",
	"agm_ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=681584518",
	"agm_gnss":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=838522739",
	"agm_salh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=979430030",
	"agm_losbr":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1015750365",
	"agm_saclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1114677066",
	"agm_mclos":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=738722044",
	"gbu_tv":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1023650302",
	"gbu_ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1902633707",
	"gbu_gnss":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=934904667",
	"gbu_salh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=515799062",
	"sam_arh":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=732148886",
	"sam_ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=330501771",
	"sam_ir_optical":     "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=352246062",
	"sam_losbr":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=816252431",
	"sam_saclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1659987410",
	"atgm_ir":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=939030645",
	"atgm_losbr":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1210040000",
	"atgm_saclos":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1557896163",
	"atgm_mclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1498988859",
	"ashm_arh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1251416353",
	"ashm_saclos":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1042355629",
	"sam_ir_naval":       "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1023721375",
	"sam_saclos_naval":   "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=677045752",
}

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsUpdater interface {
	Update(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
	Weapons(context.Context) ([]*types.Weapon, error)
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

func (s *Service) InsertWeapons(ctx context.Context) error {
	wg := &sync.WaitGroup{}

	errCh := make(chan error, len(tables))
	dataCh := make(chan []*types.Weapon, len(tables))

	for _, url := range tables {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			data, err := csvparser.ParseTable(ctx, url)
			if err != nil {
				s.log.Error("failed to parse table", sl.Err(err))
				errCh <- err
				return
			}

			dataCh <- data
		}(url)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(dataCh)
		slog.Debug("channels closed")
	}()

	var weapons []*types.Weapon

	for range tables {
		select {
		case <-ctx.Done():
			slog.Warn("parsing cancelled", sl.Err(ctx.Err()))
			return ctx.Err()
		case err := <-errCh:
			if err != nil {
				return err
			}
		case data := <-dataCh:
			weapons = append(weapons, data...)
		}
	}

	return s.inserter.Insert(ctx, weapons)
}

func (s *Service) UpdateWeapons(ctx context.Context) error {
	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	return nil, nil
}

func (s *Service) GetWeapons(ctx context.Context) ([]*types.Weapon, error) {
	weapons, err := s.provider.Weapons(ctx)
	if err != nil {
		s.log.Error("failed to get weapons", sl.Err(err))
		return nil, err
	}

	s.log.Info("get weapons OK", "amount", len(weapons))

	return weapons, nil
}
