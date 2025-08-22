package weaponsparser

import (
	"context"
	"fmt"

	csvreader "github.com/erknas/wt-guided-weapons/internal/lib/csv-reader"
	"github.com/erknas/wt-guided-weapons/internal/types"
)

type Mapper interface {
	Map(data [][]string, category string, weaponIdx int) (*types.Weapon, error)
}

type CSVWeaponParser struct {
	reader csvreader.Reader
	mapper Mapper
}

func New(reader csvreader.Reader, mapper Mapper) *CSVWeaponParser {
	return &CSVWeaponParser{
		reader: reader,
		mapper: mapper,
	}
}

func (p *CSVWeaponParser) Parse(ctx context.Context, category, url string) ([]*types.Weapon, error) {
	data, err := p.reader.Read(ctx, url)
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	var weapons []*types.Weapon

	for i := range data[0][1:] {
		weapon, err := p.mapper.Map(data, category, i+1)
		if err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}
