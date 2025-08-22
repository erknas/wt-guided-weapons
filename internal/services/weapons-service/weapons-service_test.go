package weaponsservice

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

type mockWeaponsUpserter struct {
	mock.Mock
}

type mockWeaponsProvider struct {
	mock.Mock
}

type mockWeaponsAggregator struct {
	mock.Mock
}

type mockVersionUpdater struct {
	mock.Mock
}

func (m *mockWeaponsUpserter) UpsertWeapons(ctx context.Context, weapons []*types.Weapon) error {
	args := m.Called(ctx, weapons)
	return args.Error(0)
}

func (m *mockWeaponsProvider) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockWeaponsProvider) WeaponsByName(ctx context.Context, query string) ([]types.SearchResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]types.SearchResult), args.Error(1)
}

func (m *mockWeaponsAggregator) AggregateWeapons(ctx context.Context) ([]*types.Weapon, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockVersionUpdater) UpdateVersion(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func TestWeaponsService_UpdateWeapons(t *testing.T) {
	weapons := []*types.Weapon{
		{Category: "sam-ir", Name: "9M39 Igla"},
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "sam-ir", Name: "FB-10"},
		{Category: "sam-ir", Name: "HN-6"},
	}

	tests := []struct {
		name        string
		mocks       func(*mockWeaponsAggregator, *mockWeaponsUpserter, *mockVersionUpdater)
		ctx         func() context.Context
		wantErr     bool
		containsErr string
		checkErr    func(*testing.T, error)
	}{
		{
			name: "success",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return(weapons, nil)
				mwu.On("UpsertWeapons", mock.Anything, weapons).Return(nil)
				mvu.On("UpdateVersion", mock.Anything).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail Upsert error",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return(weapons, nil)
				mwu.On("UpsertWeapons", mock.Anything, weapons).Return(errors.New("failed to upsert documents"))
			},
			wantErr:     true,
			containsErr: "failed to upsert documents",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to upsert documents")
			},
		},
		{
			name: "fail Aggregate error",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return([]*types.Weapon{}, errors.New("failed to parse table"))
			},
			wantErr:     true,
			containsErr: "failed to aggregate weapons",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to parse table")
			},
		},
		{
			name: "fail UpdateVersion error",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return(weapons, nil)
				mwu.On("UpsertWeapons", mock.Anything, weapons).Return(nil)
				mvu.On("UpdateVersion", mock.Anything).Return(errors.New("failed to update version"))
			},
			wantErr:     true,
			containsErr: "failed to update version",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to update version")
			},
		},
		{
			name: "fail Aggregate context cancelled",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return([]*types.Weapon{}, fmt.Errorf("failed to parse table: %w", context.Canceled))
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
			name: "fail Aggregate context timeout",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return([]*types.Weapon{}, fmt.Errorf("failed to parse table: %w", context.DeadlineExceeded))
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
		{
			name: "fail Upsert context cancelled",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return(weapons, nil)
				mwu.On("UpsertWeapons", mock.Anything, mock.Anything).Return(fmt.Errorf("failed to upsert documents: %w", context.Canceled))
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithCancel(context.Background())
				cancel()
				return ctx
			},
			wantErr:     true,
			containsErr: "failed to upsert documents",
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, context.Canceled)
			},
		},
		{
			name: "fail Upsert context timeout",
			mocks: func(mwa *mockWeaponsAggregator, mwu *mockWeaponsUpserter, mvu *mockVersionUpdater) {
				mwa.On("AggregateWeapons", mock.Anything).Return(weapons, nil)
				mwu.On("UpsertWeapons", mock.Anything, mock.Anything).Return(fmt.Errorf("failed to upsert documents: %w", context.DeadlineExceeded))
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
				defer cancel()
				time.Sleep(time.Millisecond)
				return ctx
			},
			wantErr:     true,
			containsErr: "failed to upsert documents",
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, context.DeadlineExceeded)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockWeaponsAggregator := new(mockWeaponsAggregator)
			mockWeaponsUpserter := new(mockWeaponsUpserter)
			mockVersionUpdater := new(mockVersionUpdater)
			tt.mocks(mockWeaponsAggregator, mockWeaponsUpserter, mockVersionUpdater)

			service := &WeaponsService{
				weaponsAggregator: mockWeaponsAggregator,
				weaponsUpserter:   mockWeaponsUpserter,
				versionUpdater:    mockVersionUpdater,
			}

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			err := service.UpdateWeapons(ctx)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.containsErr)
				if tt.checkErr != nil {
					tt.checkErr(t, err)
				}
			} else {
				require.NoError(t, err)
			}

			mockWeaponsAggregator.AssertExpectations(t)
			mockWeaponsUpserter.AssertExpectations(t)
		})
	}
}

func TestWeaponsService_GetWeaponsByCategory(t *testing.T) {
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
		name:     "success",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "sam-ir").Return(weapons, nil)
		},
		wantErr: false,
	}, {
		name:     "fail ByCategory error",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByCategory", mock.Anything, "sam-ir").Return([]*types.Weapon{}, errors.New("failed to find documents"))
		},
		wantErr:     true,
		containsErr: "failed to find documents",
		checkErr: func(t *testing.T, err error) {
			assert.Contains(t, err.Error(), "failed to find documents")
		},
	}, {
		name:     "fail ByCategory context cancelled",
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
		containsErr: "context canceled",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.Canceled)
		},
	}, {
		name:     "fail ByCategory context timeout",
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
		containsErr: "context deadline exceeded",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProvider := new(mockWeaponsProvider)
			tt.mocks(mockProvider)

			service := &WeaponsService{
				weaponsProvider: mockProvider,
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

func TestWeaponsService_SearchWeapons(t *testing.T) {
	results := []types.SearchResult{
		{Category: "sam-ir", Name: "AIM-9X"},
		{Category: "aam-ir-rear-aspect", Name: "AIM-9B"},
		{Category: "aam-ir-all-aspect", Name: "AIM-9M"},
	}

	tests := []struct {
		name        string
		query       string
		mocks       func(*mockWeaponsProvider)
		ctx         func() context.Context
		wantErr     bool
		containsErr string
		checkErr    func(*testing.T, error)
	}{{
		name:  "success",
		query: "aim",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByName", mock.Anything, "aim").Return(results, nil)
		},
		wantErr: false,
	}, {
		name:  "fail Search error",
		query: "qn",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByName", mock.Anything, "qn").Return([]types.SearchResult{}, errors.New("failed to find documents"))
		},
		wantErr:     true,
		containsErr: "failed to find documents",
		checkErr: func(t *testing.T, err error) {
			assert.Contains(t, err.Error(), "failed to find documents")
		},
	}, {
		name:  "fail Search context cancelled",
		query: "hn",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByName", mock.Anything, "hn").Return([]types.SearchResult{}, context.Canceled)
		},
		ctx: func() context.Context {
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			return ctx
		},
		wantErr:     true,
		containsErr: "context canceled",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.Canceled)
		},
	}, {
		name:  "fail Search context timeout",
		query: "aim",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("WeaponsByName", mock.Anything, "aim").Return([]types.SearchResult{}, context.DeadlineExceeded)
		},
		ctx: func() context.Context {
			ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
			defer cancel()
			time.Sleep(time.Millisecond)
			return ctx
		},
		wantErr:     true,
		containsErr: "context deadline exceeded",
		checkErr: func(t *testing.T, err error) {
			assert.ErrorIs(t, err, context.DeadlineExceeded)
		},
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockProvider := new(mockWeaponsProvider)
			tt.mocks(mockProvider)

			service := &WeaponsService{
				weaponsProvider: mockProvider,
			}

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := service.SearchWeapons(ctx, tt.query)

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
				assert.Len(t, res, 3)
			}

			mockProvider.AssertExpectations(t)
		})
	}
}
