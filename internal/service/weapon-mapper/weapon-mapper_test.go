package weaponmapper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapper(t *testing.T) {
	tests := []struct {
		name        string
		data        [][]string
		category    string
		weaponIdx   int
		expectedErr error
	}{
		{
			name:        "Succes",
			data:        [][]string{{"Category:", "aam-arh", "Name:", "AAM-4"}},
			category:    "aam-arh",
			weaponIdx:   2,
			expectedErr: nil,
		},
		{
			name:        "Invalid data",
			data:        nil,
			expectedErr: errors.New("invalid data"),
		},
		{
			name:        "Invalid weaponIdx",
			data:        [][]string{{"Name:", "AAM-4", "AIM-54A Phoenix"}},
			category:    "aam-arh",
			weaponIdx:   200,
			expectedErr: errors.New("invalid data"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mapper := &WeaponMapper{}

			res, err := mapper.Map(tt.data, tt.category, tt.weaponIdx)

			if tt.expectedErr != nil {
				require.Error(t, err)
				assert.EqualError(t, err, tt.expectedErr.Error())
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, res)
			}
		})
	}
}
