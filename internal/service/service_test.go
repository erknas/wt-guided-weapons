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
	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	tests := []struct {
		name        string
		mocks       func(*mockWeaponsAggregator, *mockWeaponsInserter)
		ctx         func() context.Context
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
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
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
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
				defer cancel()
				time.Sleep(time.Millisecond)
				return ctx
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

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
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
	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	tests := []struct {
		name        string
		category    string
		mocks       func(*mockWeaponsProvider)
		ctx         func() context.Context
		wantErr     bool
		containsErr string
		checkErr    func(*testing.T, error)
	}{{
		name:     "Success GetWeaponsByCategory",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "sam-ir").Return(weapons, nil)
		},
		wantErr: false,
	}, {
		name:     "Fail GetWeaponsByCategory",
		category: "samir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "samir").Return([]*types.Weapon{}, errors.New("failed to find documents"))
		},
		wantErr:     true,
		containsErr: "failed to get weapons by category",
		checkErr: func(t *testing.T, err error) {
			assert.Contains(t, err.Error(), "failed to find documents")
		},
	}, {
		name:     "Fail GetWeaponsByCategory context cancelled",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "sam-ir").Return([]*types.Weapon{}, context.Canceled)
		},
		ctx: func() context.Context {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			return ctx
		},
		wantErr:     true,
		containsErr: "failed to get weapons by category",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.Canceled)
		},
	}, {
		name:     "Fail GetWeaponsByCateogry context timeout",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "sam-ir").Return([]*types.Weapon{}, context.DeadlineExceeded)
		},
		ctx: func() context.Context {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
			defer cancel()
			time.Sleep(time.Millisecond)
			return ctx
		},
		wantErr:     true,
		containsErr: "failed to get weapons by category",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProvider := new(mockWeaponsProvider)
			tt.mocks(mockProvider)

			service := &Service{
				provider: mockProvider,
			}

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := service.GetWeaponsByCategory(ctx, tt.category)

			if tt.wantErr {
				require.Error(t, err)
				assert.Empty(t, res)
				assert.Contains(t, err.Error(), tt.containsErr)
				if tt.checkErr != nil {
					tt.checkErr(t, err)
				}
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, res)
			}

			mockProvider.AssertExpectations(t)
		})
	}
}
