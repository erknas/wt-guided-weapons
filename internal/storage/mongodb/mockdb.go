package mongodb

import (
	"context"
	"strings"
	"sync"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

type MockDB struct {
	storage map[string]*types.Weapon
	mu      sync.RWMutex
}

func NewMockDB() *MockDB {
	return &MockDB{
		storage: make(map[string]*types.Weapon),
	}
}

func (m *MockDB) UpsertWeapons(ctx context.Context, weapons []*types.Weapon) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	for _, weapon := range weapons {
		if weapon.ID == "" {
			weapon.ID = generateWeaponID(weapon)
		}

		m.storage[weapon.ID] = weapon
	}

	return nil
}

func (m *MockDB) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	weapons := make([]*types.Weapon, 0)

	for _, weapon := range m.storage {
		if weapon.Category == category {
			weapons = append(weapons, weapon)
		}
	}

	return weapons, nil
}

func (m *MockDB) WeaponsByName(ctx context.Context, query string) ([]types.SearchResult, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	results := make([]types.SearchResult, 0)

	for _, weapon := range m.storage {
		if strings.Contains(strings.ToLower(weapon.Name), strings.ToLower(query)) {
			result := types.SearchResult{Category: weapon.Category, Name: weapon.Name}
			results = append(results, result)
		}
	}

	return results, nil
}
