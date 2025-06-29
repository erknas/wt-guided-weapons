package service

import (
	"context"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockWeaponsAggregator struct {
	mock.Mock
}

type MockWeaponsInserter struct {
	mock.Mock
}

type MockWeaponsProvider struct {
	mock.Mock
}

func (m *MockWeaponsAggregator) Aggregate(ctx context.Context) ([]*types.Weapon, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *MockWeaponsInserter) Insert(ctx context.Context, weapons []*types.Weapon) error {
	args := m.Called(ctx, weapons)
	return args.Error(0)
}

func (m *MockWeaponsProvider) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func TestService_InsertWeapons(t *testing.T) {
	ctx := context.Background()

	mockAggregator := new(MockWeaponsAggregator)
	mockInserter := new(MockWeaponsInserter)

	service := &Service{
		aggregator: mockAggregator,
		inserter:   mockInserter,
	}

	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	t.Run("Success InsertWeapons", func(t *testing.T) {
		mockAggregator.On("Aggregate", ctx).Return(weapons, nil)
		mockInserter.On("Insert", ctx, weapons).Return(nil)

		err := service.InsertWeapons(ctx)
		require.NoError(t, err)
		mockAggregator.AssertExpectations(t)
		mockInserter.AssertExpectations(t)
	})
}

func TestService_GetWeaponsByCategory(t *testing.T) {
	ctx := context.Background()

	mockProvider := new(MockWeaponsProvider)

	service := &Service{
		provider: mockProvider,
	}

	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	t.Run("Success GetWeaponsByCategory returns slice of weapons", func(t *testing.T) {
		mockProvider.On("WeaponsByCategory", ctx, "sam-ir").Return(weapons, nil)

		res, err := service.GetWeaponsByCategory(ctx, "sam-ir")
		require.NoError(t, err)
		assert.Equal(t, weapons, res)
		mockProvider.AssertExpectations(t)
	})
}
