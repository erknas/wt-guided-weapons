package mongodb

import (
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateWeaponID(t *testing.T) {
	weapon1 := &types.Weapon{
		Name:            "Kh-29L",
		Category:        "agm-salh",
		AdditionalNotes: "Variant used by Su-17M2",
	}

	weapon2 := &types.Weapon{
		Name:            "Kh-29L",
		Category:        "agm-salh",
		AdditionalNotes: "Variant used by everything else",
	}

	weaponID1 := generateWeaponID(weapon1)
	weaponID2 := generateWeaponID(weapon2)

	assert.Len(t, weaponID1, 16)
	assert.Len(t, weaponID2, 16)
	assert.Regexp(t, "^[a-f0-9]{16}$", weaponID1)
	assert.Regexp(t, "^[a-f0-9]{16}$", weaponID2)
	assert.NotEqual(t, weaponID1, weaponID2)
}
