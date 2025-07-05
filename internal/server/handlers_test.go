package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockServicer struct {
	mock.Mock
}

func (m *mockServicer) InsertWeapons(ctx context.Context) error {
	return nil
}

func (m *mockServicer) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func TestHandleGetWeaponsByCategory(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mocServicer := new(mockServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mocServicer, urls, zap.NewNop())

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("category", "gbu-ir")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		weapons := []*types.Weapon{
			{Category: "gbu-ir", Name: "SPICE 1000"},
			{Category: "gbu-ir", Name: "SPICE 2000"},
		}

		mocServicer.On("GetWeaponsByCategory", mock.AnythingOfType("*context.valueCtx"), "gbu-ir").Return(weapons, nil)

		err = server.handleGetWeaponsByCategory(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.Weapons
		err = json.NewDecoder(rr.Result().Body).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, weapons, res.Weapons)

		mocServicer.AssertExpectations(t)
	})

	t.Run("Category does not exist", func(t *testing.T) {
		mockServicer := new(mockServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mockServicer, urls, zap.NewNop())

		rr := httptest.NewRecorder()

		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("category", "gbuir")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		mockServicer.On("GetWeaponsByCategory", mock.Anything, mock.Anything).Return([]*types.Weapon{}, nil)

		err = server.handleGetWeaponsByCategory(rr, req)
		require.Error(t, err)
		assert.Equal(t, err.Error(), "category gbuir does not exist")

		assert.Equal(t, rr.Result().StatusCode, http.StatusBadRequest)

		mockServicer.AssertNotCalled(t, "GetWeaponsByCategory")
	})
}
