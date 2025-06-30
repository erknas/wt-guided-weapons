package aggregator

import (
	"context"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockTableParser struct {
	mock.Mock
}

func (m *mockTableParser) Parse(ctx context.Context, category, url string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category, url)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func TestAggregator(t *testing.T) {
	ctx := context.Background()

	urls := map[string]string{
		"aam-sarh": "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=128448244",
		"aam-arh":  "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=650249168",
	}

	mockParser := new(mockTableParser)

	weapons := &Weapons{
		urls:   urls,
		parser: mockParser,
		log:    zap.NewNop(),
	}

	aamSarh := []*types.Weapon{{Category: "aam-sarh", Name: "AIM-7C Sparrow"}, {Category: "aam-sarh", Name: "AIM-7D Sparrow"}}
	aamArh := []*types.Weapon{{Category: "aam-arh", Name: "AAM-4"}, {Category: "aam-arh", Name: "AIM-54A Phoenix"}}

	t.Run("Aggregator", func(t *testing.T) {
		mockParser.On("Parse", mock.Anything, "aam-sarh", urls["aam-sarh"]).Return(aamSarh, nil)
		mockParser.On("Parse", mock.Anything, "aam-arh", urls["aam-arh"]).Return(aamArh, nil)

		res, err := weapons.Aggregate(ctx)
		require.NoError(t, err)
		assert.Len(t, res, 4)
		mockParser.AssertExpectations(t)
	})
}
