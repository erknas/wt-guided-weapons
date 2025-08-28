package observer

import (
	"context"
	"errors"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockVersionProvider struct {
	mock.Mock
}

type mockVersionParser struct {
	mock.Mock
}

type mockWeaponsUpdater struct {
	mock.Mock
}

func (m *mockVersionProvider) GetVersion(ctx context.Context) (types.LastChange, error) {
	args := m.Called(ctx)
	return args.Get(0).(types.LastChange), args.Error(1)
}

func (m *mockVersionParser) Parse(ctx context.Context, url string) (types.VersionInfo, error) {
	args := m.Called(ctx, url)
	return args.Get(0).(types.VersionInfo), args.Error(1)
}

func (m *mockWeaponsUpdater) UpdateWeapons(ctx context.Context) error {
	agrs := m.Called(ctx)
	return agrs.Error(0)
}

func TestObserver_checkVersionChange(t *testing.T) {
	tests := []struct {
		name        string
		mocks       func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater)
		wantErr     bool
		containsErr string
	}{
		{
			name: "success",
			mocks: func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater) {
				mvpr.On("GetVersion", mock.AnythingOfType("*context.timerCtx")).Return(types.LastChange{Version: types.VersionInfo{Version: "2.47"}}, nil)
				mvpa.On("Parse", mock.AnythingOfType("*context.timerCtx"), "test-url").Return(types.VersionInfo{Version: "2.49"}, nil)
				mwu.On("UpdateWeapons", mock.AnythingOfType("*context.timerCtx")).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "same version",
			mocks: func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater) {
				mvpr.On("GetVersion", mock.AnythingOfType("*context.timerCtx")).Return(types.LastChange{Version: types.VersionInfo{Version: "2.47"}}, nil)
				mvpa.On("Parse", mock.AnythingOfType("*context.timerCtx"), "test-url").Return(types.VersionInfo{Version: "2.47"}, nil)
			},
			wantErr: false,
		},
		{
			name: "failed GetVersion error",
			mocks: func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater) {
				mvpr.On("GetVersion", mock.AnythingOfType("*context.timerCtx")).Return(types.LastChange{}, errors.New("failed to get version"))
				mvpa.On("Parse", mock.AnythingOfType("*context.timerCtx"), "test-url").Return(types.VersionInfo{Version: "2.47"}, nil)
			},
			wantErr:     true,
			containsErr: "failed to get current version",
		},
		{
			name: "failed Parse error",
			mocks: func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater) {
				mvpr.On("GetVersion", mock.AnythingOfType("*context.timerCtx")).Return(types.LastChange{Version: types.VersionInfo{Version: "2.47"}}, nil)
				mvpa.On("Parse", mock.AnythingOfType("*context.timerCtx"), "test-url").Return(types.VersionInfo{}, errors.New("failed to read CSV"))
			},
			wantErr:     true,
			containsErr: "failed to get new version",
		},
		{
			name: "failed UpdateWeapons error",
			mocks: func(mvpa *mockVersionParser, mvpr *mockVersionProvider, mwu *mockWeaponsUpdater) {
				mvpr.On("GetVersion", mock.AnythingOfType("*context.timerCtx")).Return(types.LastChange{Version: types.VersionInfo{Version: "2.47"}}, nil)
				mvpa.On("Parse", mock.AnythingOfType("*context.timerCtx"), "test-url").Return(types.VersionInfo{Version: "2.49"}, nil)
				mwu.On("UpdateWeapons", mock.AnythingOfType("*context.timerCtx")).Return(errors.New("failed to aggregate weapons"))
			},
			wantErr:     true,
			containsErr: "failed to update weapons",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mvpr := new(mockVersionProvider)
			mvpa := new(mockVersionParser)
			mwu := new(mockWeaponsUpdater)

			tt.mocks(mvpa, mvpr, mwu)

			observer := New(mvpr, mvpa, mwu, zap.NewNop(), "test-url")

			ctx := context.Background()
			err := observer.checkVersionChange(ctx)

			if tt.wantErr {
				require.Error(t, err)
				if tt.containsErr != "" {
					assert.Contains(t, err.Error(), tt.containsErr)
				}
			} else {
				require.NoError(t, err)
			}

			mvpr.AssertExpectations(t)
			mvpa.AssertExpectations(t)
			mwu.AssertExpectations(t)

			if tt.name == "same version" {
				mwu.AssertNotCalled(t, "UpdateWeapons")
			}
		})
	}
}
