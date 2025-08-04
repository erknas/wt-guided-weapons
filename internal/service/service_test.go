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

func (m *mockWeaponsProvider) ByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockWeaponsProvider) Search(ctx context.Context, query string) ([]types.SearchResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]types.SearchResult), args.Error(1)
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
			name: "Success",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return(weapons, nil)
				mwi.On("Insert", mock.Anything, weapons).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "Insert error",
			mocks: func(mwa *mockWeaponsAggregator, mwi *mockWeaponsInserter) {
				mwa.On("Aggregate", mock.Anything).Return(weapons, nil)
				mwi.On("Insert", mock.Anything, weapons).Return(errors.New("failed to insert documents"))
			},
			wantErr:     true,
			containsErr: "failed to insert documents",
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to insert documents")
			},
		},
		{
			name: "Aggregate error",
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
			name: "Aggregate context cancelled",
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
			name: "Aggregate context timeout",
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

func TestService_WeaponsByCategory(t *testing.T) {
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
		name:     "Success WeaponsByCategory",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("ByCategory", mock.Anything, "sam-ir").Return(weapons, nil)
		},
		wantErr: false,
	}, {
		name:     "Fail WeaponsByCategory",
		category: "samir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("ByCategory", mock.Anything, "samir").Return([]*types.Weapon{}, errors.New("failed to find documents"))
		},
		wantErr:     true,
		containsErr: "failed to find documents",
		checkErr: func(t *testing.T, err error) {
			assert.Contains(t, err.Error(), "failed to find documents")
		},
	}, {
		name:     "Fail WeaponsByCategory context cancelled",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("ByCategory", mock.Anything, "sam-ir").Return([]*types.Weapon{}, context.Canceled)
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
		name:     "Fail WeaponsByCateogry context timeout",
		category: "sam-ir",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("ByCategory", mock.Anything, "sam-ir").Return([]*types.Weapon{}, context.DeadlineExceeded)
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

			service := &Service{
				provider: mockProvider,
			}

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := service.WeaponsByCategory(ctx, tt.category)

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

func TestService_SearchWeapons(t *testing.T) {
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
		name:  "Success SearchWeapon",
		query: "aim",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("Search", mock.Anything, "aim").Return(results, nil)
		},
		wantErr: false,
	}, {
		name:  "Fail SearchWeapon",
		query: "qn",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("Search", mock.Anything, "qn").Return([]types.SearchResult{}, errors.New("failed to find documents"))
		},
		wantErr:     true,
		containsErr: "failed to find documents",
		checkErr: func(t *testing.T, err error) {
			assert.Contains(t, err.Error(), "failed to find documents")
		},
	}, {
		name:  "Fail SearchWeapon context cancelled",
		query: "hn",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("Search", mock.Anything, "hn").Return([]types.SearchResult{}, context.Canceled)
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
		name:  "Fail WeaponsByCateogry context timeout",
		query: "aim",
		mocks: func(mwp *mockWeaponsProvider) {
			mwp.On("Search", mock.Anything, "aim").Return([]types.SearchResult{}, context.DeadlineExceeded)
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

			service := &Service{
				provider: mockProvider,
			}

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := service.SearchWeapon(ctx, tt.query)

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
