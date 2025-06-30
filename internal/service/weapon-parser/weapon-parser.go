package weaponparser

import (
	"context"
	"fmt"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

type Reader interface {
	Read(ctx context.Context, url string) ([][]string, error)
}

type Mapper interface {
	Map(data [][]string, category string, weaponIdx int) (*types.Weapon, error)
}

type CSVWeaponParser struct {
	reader Reader
	mapper Mapper
}

func New(reader Reader, mapper Mapper) *CSVWeaponParser {
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
			return nil, fmt.Errorf("mapping CSV to struct failed: %w", err)
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}
