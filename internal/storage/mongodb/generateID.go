package mongodb

import (
	"crypto/sha256"
	"fmt"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

func generateWeaponID(weapon *types.Weapon) string {
	data := fmt.Sprintf("%s-%s-%s", weapon.Name, weapon.Category, weapon.AdditionalNotes)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)[:16]
}
