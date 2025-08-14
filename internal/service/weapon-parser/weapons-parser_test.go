package weaponparser

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockReader struct {
	mock.Mock
}

type mockMapper struct {
	mock.Mock
}

func (m *mockReader) Read(ctx context.Context, url string) ([][]string, error) {
	args := m.Called(ctx, url)
	return args.Get(0).([][]string), args.Error(1)
}

func (m *mockMapper) Map(data [][]string, category string, weaponIdx int) (*types.Weapon, error) {
	args := m.Called(data, category, weaponIdx)
	return args.Get(0).(*types.Weapon), args.Error(1)
}

func TestParse(t *testing.T) {
	testWeapon := &types.Weapon{Category: "category", Name: "QN502C"}
	testData := [][]string{
		{"Name:", "QN502C"},
		{"Name:", "Spike LR2"},
	}

	tests := []struct {
		name     string
		mocks    func(*mockReader, *mockMapper)
		ctx      func() context.Context
		wantErr  bool
		checkErr func(*testing.T, error)
	}{
		{
			name: "success",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return(testData, nil)
				mm.On("Map", testData, "category", 1).Return(testWeapon, nil)
			},
			wantErr: false,
		},
		{
			name: "fail Read error",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("failed to create new request"))
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to read CSV")
			},
		},
		{
			name: "fail Read error",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("failed to make HTTP request"))
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to read CSV")
			},
		},
		{
			name: "fail Read error",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("unexpected status code"))
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to read CSV")
			},
		},
		{
			name: "fail Read context timeout",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, context.DeadlineExceeded)
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
				defer cancel()
				time.Sleep(time.Millisecond)
				return ctx
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, context.DeadlineExceeded)
			},
		},
		{
			name: "fail Map error",
			mocks: func(mr *mockReader, mm *mockMapper) {
				mr.On("Read", mock.Anything, "test-url").Return(testData, nil)
				mm.On("Map", testData, "category", 1).Return(testWeapon, errors.New("invalid data"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockReader := new(mockReader)
			mockMapper := new(mockMapper)
			tt.mocks(mockReader, mockMapper)

			parser := New(mockReader, mockMapper)

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := parser.Parse(ctx, "category", "test-url")

			if tt.wantErr {
				require.Error(t, err)
				assert.Empty(t, res)
				if tt.checkErr != nil {
					tt.checkErr(t, err)
				}
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, res)
			}

			mockReader.AssertExpectations(t)
			mockMapper.AssertExpectations(t)
		})
	}
}
