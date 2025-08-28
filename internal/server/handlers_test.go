package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/erknas/wt-guided-weapons/internal/lib/api"
	apierrors "github.com/erknas/wt-guided-weapons/internal/lib/api/api-errors"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

type mockWeaponsServicer struct {
	mock.Mock
}

type mockVersionServicer struct {
	mock.Mock
}

func (m *mockWeaponsServicer) UpdateWeapons(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *mockWeaponsServicer) GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockWeaponsServicer) SearchWeapons(ctx context.Context, query string) ([]types.SearchResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]types.SearchResult), args.Error(1)
}

func (m *mockVersionServicer) GetVersion(ctx context.Context) (types.LastChange, error) {
	args := m.Called(ctx)
	return args.Get(0).(types.LastChange), args.Error(1)
}

func TestHandleGetWeaponsByCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockWeaponsServicer := new(mockWeaponsServicer)
		mockVersionServicer := new(mockVersionServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mockWeaponsServicer, mockVersionServicer, urls, zap.NewNop())

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

		mockWeaponsServicer.On("GetWeaponsByCategory", mock.AnythingOfType("*context.valueCtx"), "gbu-ir").Return(weapons, nil)

		err = server.handleGetWeaponsByCategory(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.Weapons
		err = json.NewDecoder(rr.Result().Body).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, weapons, res.Weapons)

		mockWeaponsServicer.AssertExpectations(t)
	})

	t.Run("category does not exist", func(t *testing.T) {
		mockWeaponsServicer := new(mockWeaponsServicer)
		mockVersionServicer := new(mockVersionServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mockWeaponsServicer, mockVersionServicer, urls, zap.NewNop())

		r := chi.NewRouter()
		r.With(logger.MiddlewareCategoryCheck(server.categories)).Get("/", api.MakeHTTPFunc(server.handleGetWeaponsByCategory))

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("category", "gbuir")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		r.ServeHTTP(rr, req)

		var resp apierrors.APIError
		err = json.NewDecoder(rr.Body).Decode(&resp)
		require.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, rr.Result().StatusCode)
		assert.Equal(t, resp.Message, "category gbuir does not exist")

		mockWeaponsServicer.AssertNotCalled(t, "GetWeaponsByCategory")
	})
}

func TestHandleSearchWeapons(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockWeaponsServicer := new(mockWeaponsServicer)
		mockVersionServicer := new(mockVersionServicer)

		server := New(mockWeaponsServicer, mockVersionServicer, map[string]string{}, zap.NewNop())

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("name", "spice")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		searchResults := []types.SearchResult{
			{Category: "gbu-ir", Name: "SPICE 1000"},
			{Category: "gbu-ir", Name: "SPICE 2000"},
		}

		results := types.SearchResults{Results: searchResults}

		mockWeaponsServicer.On("SearchWeapons", mock.AnythingOfType("*context.valueCtx"), "spice").Return(searchResults, nil)

		err = server.handleSeachWeapons(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.SearchResults
		err = json.NewDecoder(rr.Result().Body).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, results, res)

		mockWeaponsServicer.AssertExpectations(t)
	})
}

func TestHandleGetVersion(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockWeaponsServicer := new(mockWeaponsServicer)
		mockVersionServicer := new(mockVersionServicer)

		server := New(mockWeaponsServicer, mockVersionServicer, map[string]string{}, zap.NewNop())

		rr := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "", nil)
		require.NoError(t, err)

		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("name", "spice")
		req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

		version := types.VersionInfo{Version: "24.0.1.144"}
		lastChange := types.LastChange{Version: version}

		mockVersionServicer.On("GetVersion", mock.AnythingOfType("*context.valueCtx")).Return(lastChange, nil)

		err = server.handleGetVersion(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.LastChange
		err = json.NewDecoder(rr.Result().Body).Decode(&res.Version)
		require.NoError(t, err)
		assert.Equal(t, lastChange, res)

		mockWeaponsServicer.AssertExpectations(t)
	})
}
