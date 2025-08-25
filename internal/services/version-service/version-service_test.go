package versionservice

import (
	"context"
	"errors"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockVersionUpserter struct {
	mock.Mock
}

type mockVersionProvider struct {
	mock.Mock
}

type mockVersionParser struct {
	mock.Mock
}

func (m *mockVersionUpserter) UpsertVersion(ctx context.Context, version types.VersionInfo) error {
	args := m.Called(ctx, version)
	return args.Error(0)
}

func (m *mockVersionProvider) Version(ctx context.Context) (types.LastChange, error) {
	args := m.Called(ctx)
	return args.Get(0).(types.LastChange), args.Error(1)
}

func (m *mockVersionParser) Parse(ctx context.Context, url string) (types.VersionInfo, error) {
	args := m.Called(ctx, url)
	return args.Get(0).(types.VersionInfo), args.Error(1)
}

func TestVersionService_UpdateVersion(t *testing.T) {
	version := types.VersionInfo{Version: "2.47.0.114"}

	test := []struct {
		name        string
		mocks       func(*mockVersionParser, *mockVersionUpserter)
		wantErr     bool
		containsErr string
	}{
		{
			name: "success",
			mocks: func(mvp *mockVersionParser, mvu *mockVersionUpserter) {
				mvp.On("Parse", mock.Anything, mock.Anything).Return(version, nil)
				mvu.On("UpsertVersion", mock.Anything, version).Return(nil)
			},
			wantErr: false,
		},
		{
			name: "fail Parse error",
			mocks: func(mvp *mockVersionParser, mvu *mockVersionUpserter) {
				mvp.On("Parse", mock.Anything, mock.Anything).Return(types.VersionInfo{}, errors.New("failed to read CSV"))
			},
			wantErr:     true,
			containsErr: "failed to parse version",
		},
		{
			name: "fail UpsertVersion error",
			mocks: func(mvp *mockVersionParser, mvu *mockVersionUpserter) {
				mvp.On("Parse", mock.Anything, mock.Anything).Return(version, nil)
				mvu.On("UpsertVersion", mock.Anything, version).Return(errors.New("failed to update document"))
			},
			wantErr:     true,
			containsErr: "failed to update version",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			mvp := new(mockVersionParser)
			mvu := new(mockVersionUpserter)
			tt.mocks(mvp, mvu)

			service := New(mvu, nil, mvp, "test-url")

			ctx := context.Background()

			err := service.UpdateVersion(ctx)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.containsErr)
			} else {
				require.NoError(t, err)
			}

			mvp.AssertExpectations(t)
			mvu.AssertExpectations(t)
		})
	}
}

func TestVersionService_GetVersion(t *testing.T) {
	version := types.VersionInfo{Version: "2.47.0.114"}
	lastChange := types.LastChange{Version: version}

	test := []struct {
		name        string
		mocks       func(*mockVersionProvider)
		wantErr     bool
		containsErr string
	}{
		{
			name: "success",
			mocks: func(mvp *mockVersionProvider) {
				mvp.On("Version", mock.Anything).Return(lastChange, nil)
			},
			wantErr: false,
		},
		{
			name: "fail Version error",
			mocks: func(mvp *mockVersionProvider) {
				mvp.On("Version", mock.Anything).Return(types.LastChange{}, errors.New("failed to find document"))
			},
			wantErr:     true,
			containsErr: "failed to get version",
		},
		{
			name: "fail Version error version not found",
			mocks: func(mvp *mockVersionProvider) {
				mvp.On("Version", mock.Anything).Return(types.LastChange{}, errors.New("version not found"))
			},
			wantErr:     true,
			containsErr: "failed to get version",
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			mvp := new(mockVersionProvider)
			tt.mocks(mvp)

			service := New(&mockVersionUpserter{}, mvp, &mockVersionParser{}, "test-url")

			ctx := context.Background()

			results, err := service.GetVersion(ctx)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.containsErr)
				assert.Empty(t, results)
			} else {
				require.NoError(t, err)
				assert.Equal(t, results, lastChange)
			}

			mvp.AssertExpectations(t)
		})
	}

}
