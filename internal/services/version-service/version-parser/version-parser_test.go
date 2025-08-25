package versionparser

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type mockReader struct {
	mock.Mock
}

func (m *mockReader) Read(ctx context.Context, url string) ([][]string, error) {
	args := m.Called(ctx, url)
	return args.Get(0).([][]string), args.Error(1)
}

func TestParse(t *testing.T) {
	testData := [][]string{{"asd"}, {""}, {"123"}, {"Last change in stats was in 2.47.0.123"}, {"♥"}}

	tests := []struct {
		name        string
		mock        func(mr *mockReader)
		ctx         func() context.Context
		wantErr     bool
		containsErr string
	}{
		{
			name: "success",
			mock: func(mr *mockReader) {
				mr.On("Read", mock.Anything, "test-url").Return(testData, nil)
			},
			wantErr: false,
		},
		{
			name: "fail Read error",
			mock: func(mr *mockReader) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("failed to create new request"))
			},
			wantErr:     true,
			containsErr: "failed to read CSV",
		},
		{
			name: "fail Read error",
			mock: func(mr *mockReader) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("failed to make HTTP request"))
			},
			wantErr:     true,
			containsErr: "failed to read CSV",
		},
		{
			name: "fail Read error",
			mock: func(mr *mockReader) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, errors.New("unexpected status code"))
			},
			wantErr:     true,
			containsErr: "failed to read CSV",
		},
		{
			name: "fail Read context timeout",
			mock: func(mr *mockReader) {
				mr.On("Read", mock.Anything, "test-url").Return([][]string{}, context.DeadlineExceeded)
			},
			ctx: func() context.Context {
				ctx, cancel := context.WithTimeout(context.Background(), time.Nanosecond)
				defer cancel()
				time.Sleep(time.Millisecond)
				return ctx
			},
			wantErr:     true,
			containsErr: context.DeadlineExceeded.Error(),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := new(mockReader)
			tt.mock(mr)

			parser := New(mr)

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			results, err := parser.Parse(ctx, "test-url")

			if tt.wantErr {
				require.Error(t, err)
				assert.Empty(t, results)
				assert.Contains(t, err.Error(), tt.containsErr)
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, results)
			}

			mr.AssertExpectations(t)
		})
	}
}
