package mongodb

import (
	"context"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestMongoDB_UpsertWeapons(t *testing.T) {
	ctx := context.Background()

	db := NewMockDB()

	t.Run("insert new weapons", func(t *testing.T) {
		weapons := []*types.Weapon{
			{Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{Name: "AIM-54", Category: "aam-arh"},
		}

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 2)
	})

	t.Run("update weapons", func(t *testing.T) {
		weapons := []*types.Weapon{
			{Name: "AIM-9L", Category: "aam-ir-all-aspect", Mass: "280"},
			{Name: "AIM-54", Category: "aam-arh", MassAtEndOfBoosterBurn: "450"},
		}

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 2)

		weapons[0].Mass = "270"
		weapons[1].MassAtEndOfBoosterBurn = "400"

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 2)

		for _, weapon := range db.storage {
			if weapon.Name == "AIM-9L" {
				assert.Equal(t, weapon.Mass, "270")
			}
			if weapon.Name == "AIM-54" {
				assert.Equal(t, weapon.MassAtEndOfBoosterBurn, "400")
			}
		}
	})
}

func TestMongoDB_WeaponsByCategory(t *testing.T) {
	ctx := context.Background()

	db := NewMockDB()

	t.Run("find weapons by category", func(t *testing.T) {
		weapons := []*types.Weapon{
			{Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{Name: "AIM-54", Category: "aam-arh"},
		}

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 4)

		category := "aam-ir-all-aspect"

		results, _ := db.WeaponsByCategory(ctx, category)
		assert.Equal(t, 3, len(results))
	})
}

func TestMongoDB_WeaponsByName(t *testing.T) {
	ctx := context.Background()

	t.Run("find weapons by name", func(t *testing.T) {
		db := NewMockDB()
		weapons := []*types.Weapon{
			{Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{Name: "AIM-54", Category: "aam-arh"},
		}

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 4)

		query := "aim"

		results, _ := db.WeaponsByName(ctx, query)
		assert.Len(t, results, 3)
	})

	t.Run("find weapons by name empty results", func(t *testing.T) {
		db := NewMockDB()
		weapons := []*types.Weapon{
			{ID: "1", Name: "AIM-9L", Category: "aam-ir-all-aspect"},
			{ID: "2", Name: "AIM-9M", Category: "aam-ir-all-aspect"},
			{ID: "3", Name: "AAM-3", Category: "aam-ir-all-aspect"},
			{ID: "4", Name: "AIM-54", Category: "aam-arh"},
		}

		_ = db.UpsertWeapons(ctx, weapons)
		assert.Len(t, db.storage, 4)

		query := "abfa1230"

		results, _ := db.WeaponsByName(ctx, query)
		assert.Empty(t, results)
	})
}
