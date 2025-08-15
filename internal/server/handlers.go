package server

import (
	"net/http"

	"github.com/erknas/wt-guided-weapons/internal/lib/api"
	apierrors "github.com/erknas/wt-guided-weapons/internal/lib/api/api-errors"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (s *Server) handleUpdateWeapons(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	if err := s.svc.UpsertWeapons(r.Context()); err != nil {
		log.Error("UpdateWeapons failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("UpdateWeapons handler complited")

	return api.WriteJSON(w, http.StatusOK, map[string]string{"msg": "OK"})
}

func (s *Server) handleGetWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	category := chi.URLParam(r, "category")

	weapons, err := s.svc.WeaponsByCategory(r.Context(), category)
	if err != nil {
		log.Error("WeaponsByCategory failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("GetWeaponsByCategory handler complited",
		zap.String("category", category),
		zap.Int("total weapons", len(weapons)),
	)

	return api.WriteJSON(w, http.StatusOK, types.Weapons{Weapons: weapons})
}

func (s *Server) handleSeachWeapon(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	query := chi.URLParam(r, "name")

	results, err := s.svc.SearchWeapon(r.Context(), query)
	if err != nil {
		log.Error("SearchWeapon failed",
			zap.Error(err),
		)
		return err
	}

	if len(results) == 0 {
		log.Warn("Empty search results",
			zap.String("query", query),
		)
		return apierrors.EmptySearchResults()
	}

	log.Info("SearchWeapon handler complited",
		zap.String("query", query),
		zap.Int("total weapons found", len(results)),
	)

	return api.WriteJSON(w, http.StatusOK, types.Results{Results: results})
}
