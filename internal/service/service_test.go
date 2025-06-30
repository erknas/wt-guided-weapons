package service

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockWeaponsAggregator struct {
	mock.Mock
}

type mockWeaponsInserter struct {
	mock.Mock
}

type mockWeaponsProvider struct {
	mock.Mock
}

func (m *mockWeaponsAggregator) Aggregate(ctx context.Context) ([]*types.Weapon, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockWeaponsInserter) Insert(ctx context.Context, weapons []*types.Weapon) error {
	args := m.Called(ctx, weapons)
	return args.Error(0)
}

func (m *mockWeaponsProvider) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func TestService_InsertWeapons(t *testing.T) {
	ctx := context.Background()
	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	tests := []struct {
		name        string
		mocks       func(*mockWeaponsAggregator, *mockWeaponsInserter)
		wantErr     bool
		containsErr string
		checkErr    func(*testing.T, error)
	}{
		{
			name: "Success InsertWeapons",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return(weapons, nil)
				mwi.On("Insert", mock.Anything, weapons).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Fail InsertWeapons Insert error",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return(weapons, nil)
				mwi.On("Insert", mock.Anything, weapons).Return(errors.New("failed to insert documents"))
			},
			wantErr:     true,
			containsErr: "failed to insert weapons",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to insert documents")
			},
		},
		{
			name: "Fail InsertWeapons Aggregate error",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return([]*types.Weapon{}, errors.New("failed to parse table"))
			},
			wantErr:     true,
			containsErr: "failed to aggregate weapons",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to parse table")
			},
		},
		{
			name: "Fail InsertWeapons Aggregate context cancelled",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return([]*types.Weapon{}, fmt.Errorf("failed to parse table: %w", context.Canceled))
			},
			wantErr:     true,
			containsErr: "failed to aggregate weapons",
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, context.Canceled)
			},
		},
		{
			name: "Fail InsertWeapons Aggregate context timeout",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return([]*types.Weapon{}, fmt.Errorf("failed to parse table: %w", context.DeadlineExceeded))
			},
			wantErr:     true,
			containsErr: "failed to aggregate weapons",
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, context.DeadlineExceeded)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockAggregator := new(mockWeaponsAggregator)
			mockInserter := new(mockWeaponsInserter)
			tt.mocks(mockAggregator, mockInserter)

			service := &Service{
				aggregator: mockAggregator,
				inserter:   mockInserter,
			}

			err := service.InsertWeapons(ctx)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.containsErr)
				if tt.checkErr != nil {
					tt.checkErr(t, err)
				}
			} else {
				require.NoError(t, err)
			}

			mockAggregator.AssertExpectations(t)
			mockInserter.AssertExpectations(t)
		})
	}
}

func TestService_GetWeaponsByCategory(t *testing.T) {
	ctx := context.Background()

	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	t.Run("Success GetWeaponsByCategory", func(t *testing.T) {
		mockProvider := new(mockWeaponsProvider)
		service := &Service{
			provider: mockProvider,
		}

		mockProvider.On("WeaponsByCategory", mock.Anything, "sam-ir").Return(weapons, nil)

		res, err := service.GetWeaponsByCategory(ctx, "sam-ir")
		require.NoError(t, err)
		assert.Equal(t, weapons, res)
		mockProvider.AssertExpectations(t)
	})

	t.Run("Fail GetWeaponsByCategory", func(t *testing.T) {
		mockProvider := new(mockWeaponsProvider)
		service := &Service{
			provider: mockProvider,
		}

		category := "sam-ir"
		expectedErr := errors.New("failed to find documents")
		mockProvider.On("WeaponsByCategory", mock.Anything, category).Return([]*types.Weapon{}, expectedErr)

		res, err := service.GetWeaponsByCategory(ctx, category)
		require.Error(t, err)
		assert.EqualError(t, err, fmt.Sprintf("failed to get weapons by category %s: %v", category, expectedErr))
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, res)
		mockProvider.AssertExpectations(t)
	})

	t.Run("Fail GetWeaponsByCategory context cancelled", func(t *testing.T) {
		mockProvider := new(mockWeaponsProvider)
		service := &Service{
			provider: mockProvider,
		}

		category := "sam-ir"
		expectedErr := context.Canceled
		mockProvider.On("WeaponsByCategory", mock.Anything, category).Return([]*types.Weapon{}, expectedErr)

		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		res, err := service.GetWeaponsByCategory(ctx, category)
		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, res)
		mockProvider.AssertExpectations(t)
	})

	t.Run("Fail GetWeaponsByCategory context timeout", func(t *testing.T) {
		mockProvider := new(mockWeaponsProvider)
		service := &Service{
			provider: mockProvider,
		}

		category := "sam-ir"
		expectedErr := context.Canceled
		mockProvider.On("WeaponsByCategory", mock.Anything, category).Return([]*types.Weapon{}, expectedErr)

		ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
		defer cancel()

		time.Sleep(time.Millisecond)

		res, err := service.GetWeaponsByCategory(ctx, category)
		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, res)
		mockProvider.AssertExpectations(t)
	})
}
