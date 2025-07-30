package aggregator

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

type mockTableParser struct {
	mock.Mock
}

func (m *mockTableParser) Parse(ctx context.Context, category, url string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category, url)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func TestAggregate(t *testing.T) {
	urls := map[string]string{
		"aam-sarh": "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=128448244",
		"aam-arh":  "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=650249168",
	}

	aamSarh := []*types.Weapon{
		{Name: "AIM-7C Sparrow", Category: "aam-sarh"},
		{Name: "AIM-7D Sparrow", Category: "aam-sarh"},
	}

	aamArh := []*types.Weapon{
		{Name: "AAM-4", Category: "aam-arh"},
		{Name: "AIM-54A Phoenix", Category: "aam-arh"},
	}

	tests := []struct {
		name     string
		mocks    func(*mockTableParser)
		ctx      func() context.Context
		wantErr  bool
		checkErr func(*testing.T, error)
	}{
		{
			name: "Success",
			mocks: func(mtp *mockTableParser) {
				mtp.On("Parse", mock.Anything, "aam-sarh", urls["aam-sarh"]).Return(aamSarh, nil)
				mtp.On("Parse", mock.Anything, "aam-arh", urls["aam-arh"]).Return(aamArh, nil)
			},
			wantErr: false,
		},
		{
			name: "Parse error",
			mocks: func(mtp *mockTableParser) {
				mtp.On("Parse", mock.Anything, "aam-sarh", urls["aam-sarh"]).Return([]*types.Weapon{}, errors.New("failed to read CSV"))
				mtp.On("Parse", mock.Anything, "aam-arh", urls["aam-arh"]).Return(aamArh, nil)
			},
			wantErr: true,
			checkErr: func(t *testing.T, err error) {
				assert.Contains(t, err.Error(), "failed to parse table")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockParser := new(mockTableParser)
			tt.mocks(mockParser)

			aggregator := New(urls, mockParser, zap.NewNop())

			ctx := context.Background()
			if tt.ctx != nil {
				ctx = tt.ctx()
			}

			res, err := aggregator.Aggregate(ctx)

			if tt.wantErr {
				require.Error(t, err)
				tt.checkErr(t, err)
			} else {
				require.NoError(t, err)
				assert.NotEmpty(t, res)
				assert.ElementsMatch(t, append(aamSarh, aamArh...), res)
				assert.Len(t, res, 4)
			}

			mock.AssertExpectationsForObjects(t)
		})
	}
}
