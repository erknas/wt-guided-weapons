package service

import (
	"context"
	"encoding/csv"
	"log/slog"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

const (
	table1 = "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=0"
)

type WeaponsInserter interface {
	Insert(context.Context, []*types.Weapon) error
}

type WeaponsUpdater interface {
	Update(context.Context, []*types.Weapon) error
}

type WeaponsProvider interface {
	Provide(context.Context, string) ([]*types.Weapon, error)
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
	data, err := parseTable(ctx, table1)
	if err != nil {
		return err
	}

	weapons, err := parseWeapons(data)
	if err != nil {
		return err
	}

	return s.inserter.Insert(ctx, weapons)
}

func (s *Service) UpdateWeapons(ctx context.Context) error {
	return nil
}

func (s *Service) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	return nil, nil
}

func structMapping(data [][]string, weaponIdx int) (*types.Weapon, error) {
	headers := make([]string, 0, len(data))

	for i, row := range data {
		headers[i] = strings.TrimSpace(row[0])
	}

	var weapon *types.Weapon

	val := reflect.ValueOf(weapon).Elem()
	typ := val.Type()

	for i := range val.NumField() {
		field := typ.Field(i)
		tag := field.Tag.Get("csv")

		var rowIdx = -1

		for j, header := range headers {
			if strings.Contains(header, tag) {
				rowIdx = j
				break
			}
		}

		if rowIdx == -1 {
			return nil, nil
		}

		if len(data[rowIdx]) < 2 {
			continue
		}

		value := data[rowIdx][weaponIdx]
		val.Field(i).SetString(value)
	}

	return weapon, nil
}

func parseWeapons(data [][]string) ([]*types.Weapon, error) {
	var weapons []*types.Weapon

	for i := range data[0][1:] {
		weapon, err := structMapping(data, i+1)
		if err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

func parseTable(ctx context.Context, url string) ([][]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return csv.NewReader(resp.Body).ReadAll()
}
