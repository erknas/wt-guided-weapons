package server

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	apierrors "github.com/erknas/wt-guided-weapons/internal/lib/api/api-errors"
	"github.com/erknas/wt-guided-weapons/internal/logger"
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

func (m *mockServicer) UpsertWeapons(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

func (m *mockServicer) WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error) {
	args := m.Called(ctx, category)
	return args.Get(0).([]*types.Weapon), args.Error(1)
}

func (m *mockServicer) SearchWeapon(ctx context.Context, query string) ([]types.SearchResult, error) {
	args := m.Called(ctx, query)
	return args.Get(0).([]types.SearchResult), args.Error(1)
}

func TestHandleGetWeaponsByCategory(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServicer := new(mockServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mockServicer, urls, zap.NewNop())

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

		mockServicer.On("WeaponsByCategory", mock.AnythingOfType("*context.valueCtx"), "gbu-ir").Return(weapons, nil)

		err = server.handleGetWeaponsByCategory(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.Weapons
		err = json.NewDecoder(rr.Result().Body).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, weapons, res.Weapons)

		mockServicer.AssertExpectations(t)
	})

	t.Run("category does not exist", func(t *testing.T) {
		mockServicer := new(mockServicer)
		urls := map[string]string{"gbu-ir": "test-url"}

		server := New(mockServicer, urls, zap.NewNop())

		r := chi.NewRouter()
		r.With(logger.MiddlewareCategoryCheck(server.categories)).Get("/", makeHTTPFunc(server.handleGetWeaponsByCategory))

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

		mockServicer.AssertNotCalled(t, "WeaponsByCategory")
	})
}

func TestHandleSearchWeapon(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockServicer := new(mockServicer)

		server := New(mockServicer, map[string]string{}, zap.NewNop())

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

		results := types.Results{Results: searchResults}

		mockServicer.On("SearchWeapon", mock.AnythingOfType("*context.valueCtx"), "spice").Return(searchResults, nil)

		err = server.handleSeachWeapon(rr, req)
		require.NoError(t, err)

		assert.Equal(t, rr.Result().StatusCode, http.StatusOK)

		var res types.Results
		err = json.NewDecoder(rr.Result().Body).Decode(&res)
		require.NoError(t, err)
		assert.Equal(t, results, res)

		mockServicer.AssertExpectations(t)
	})
}
