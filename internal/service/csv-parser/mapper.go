package csvparser

import (
	"reflect"
	"strings"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

const csvTag = "csv"

func MapToWeapon(data [][]string, weaponIdx int) (*types.Weapon, error) {
	headers := make([]string, len(data))

	for i, row := range data {
		headers[i] = strings.TrimSpace(row[0])
	}

	weapon := new(types.Weapon)

	val := reflect.ValueOf(weapon).Elem()
	typ := val.Type()

	for i := range val.NumField() {
		field := typ.Field(i)
		tag := field.Tag.Get(csvTag)
		if tag == "" {
			continue
		}

		rowIdx := headerIndex(headers, tag)
		if rowIdx == -1 || len(data[rowIdx]) < 2 {
			continue
		}

		value := data[rowIdx][weaponIdx]
		val.Field(i).SetString(value)
	}

	return weapon, nil
}

func headerIndex(headers []string, tag string) int {
	for i, header := range headers {
		if strings.Contains(header, tag) {
			return i
		}
	}

	return -1
}
