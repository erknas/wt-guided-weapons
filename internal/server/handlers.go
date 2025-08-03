package server

import (
	"net/http"

	"github.com/erknas/wt-guided-weapons/internal/lib/api"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (s *Server) handleInsertWeapon(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	if err := s.svc.InsertWeapons(r.Context()); err != nil {
		log.Error("InsertWeapons failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("handleInsertWeapons complited")

	return api.WriteJSON(w, http.StatusOK, map[string]string{"msg": "OK"})
}

func (s *Server) handleGetWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	category := chi.URLParam(r, "category")

	weapons, err := s.svc.GetWeaponsByCategory(r.Context(), category)
	if err != nil {
		log.Error("GetWeaponsByCategory failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("handleGetWeaponsByCategory complited",
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

	log.Info("handleSearchWeapon complited",
		zap.String("name", query),
		zap.Int("total weapons found", len(results)),
	)

	return api.WriteJSON(w, http.StatusOK, types.Results{Results: results})
}
