package service

import (
	"context"
	"sync"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	csvparser "github.com/erknas/wt-guided-weapons/internal/service/csv-parser"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"go.uber.org/zap"
)

var Tables map[string]string = map[string]string{
	"aam-ir-rear-aspect": "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=0",
	"aam-ir-all-aspect":  "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1726112384",
	"aam-ir-heli":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=638442034",
	"aam-sarh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=128448244",
	"aam-arh":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=650249168",
	"aam-manual":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=29789551",
	"agm-tv":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1614911062",
	"agm-ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=681584518",
	"agm-gnss":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=838522739",
	"agm-salh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=979430030",
	"agm-losbr":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1015750365",
	"agm-saclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1114677066",
	"agm-mclos":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=738722044",
	"gbu-tv":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1023650302",
	"gbu-ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1902633707",
	"gbu-gnss":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=934904667",
	"gbu-salh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=515799062",
	"sam-arh":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=732148886",
	"sam-ir":             "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=330501771",
	"sam-ir-optical":     "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=352246062",
	"sam-losbr":          "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=816252431",
	"sam-saclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1659987410",
	"atgm-ir":            "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=939030645",
	"atgm-losbr":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1210040000",
	"atgm-saclos":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1557896163",
	"atgm-mclos":         "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1498988859",
	"ashm-arh":           "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1251416353",
	"ashm-saclos":        "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1042355629",
	"sam-ir-naval":       "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1023721375",
	"sam-saclos-naval":   "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=677045752",
}

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	WeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
	Weapons(context.Context) ([]*types.Weapon, error)
}

type Service struct {
	inserter WeaponsInserter
	provider WeaponsProvider
	log      *zap.Logger
}

func New(inserter WeaponsInserter, provider WeaponsProvider, log *zap.Logger) *Service {
	return &Service{
		inserter: inserter,
		provider: provider,
		log:      log,
	}
}

func (s *Service) InsertWeapons(ctx context.Context) error {
	start := time.Now()
	log := logger.FromContext(ctx, logger.Service)

	wg := &sync.WaitGroup{}

	errCh := make(chan error, len(Tables))
	dataCh := make(chan []*types.Weapon, len(Tables))

	for category, url := range Tables {
		wg.Add(1)
		go func(category, url string) {
			defer wg.Done()

			data, err := csvparser.ParseTable(ctx, category, url)
			if err != nil {
				log.Error("parse table failed",
					zap.Error(err),
					zap.String("category", category),
					zap.String("table_url", url),
				)
				errCh <- err
				return
			}

			log.Debug("parse table complited",
				zap.String("category", category),
				zap.String("table_url", url),
				zap.Int("weapons_count", len(data)),
			)

			dataCh <- data
		}(category, url)
	}

	go func() {
		wg.Wait()
		close(errCh)
		close(dataCh)
		log.Debug("channels closed")
		log.Debug("all goroutine complited")
	}()

	var weapons []*types.Weapon
	successfulTables := 0

	for range Tables {
		select {
		case <-ctx.Done():
			log.Warn("parsing cancelled",
				zap.Error(ctx.Err()),
			)
			return ctx.Err()
		case err := <-errCh:
			if err != nil {
				return err
			}
		case data := <-dataCh:
			weapons = append(weapons, data...)
			successfulTables++
		}
	}

	log.Info("tabels parsing complited",
		zap.Int("total successful tables parsed", successfulTables),
		zap.Int("total failed tables", len(Tables)-successfulTables),
		zap.Int("weapons_count", len(weapons)),
		zap.Duration("duration", time.Since(start)),
	)

	if err := s.inserter.Insert(ctx, weapons); err != nil {
		log.Error("insert operation failed",
			zap.Error(err),
		)
		return err
	}

	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.WeaponsByCategory(ctx, category)
	if err != nil {
		log.Error("service call failed",
			zap.Error(err),
			zap.String("category", category),
		)
		return nil, err
	}

	log.Debug("service call complited",
		zap.String("category", category),
		zap.Int("weapons_count", len(weapons)),
	)

	return weapons, nil
}

func (s *Service) GetWeapons(ctx context.Context) ([]*types.Weapon, error) {
	log := logger.FromContext(ctx, logger.Service)

	weapons, err := s.provider.Weapons(ctx)
	if err != nil {
		log.Error("service call failed",
			zap.Error(err),
		)
		return nil, err
	}

	log.Debug("service call complited",
		zap.Int("weapons_amount", len(weapons)),
	)

	return weapons, nil
}
