package server

import (
	"net/http"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func (s *Server) handleInsertWeapon(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	if err := s.svc.InsertWeapons(r.Context()); err != nil {
		log.Error("InsertWeapons request failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("InsertWeapons request complited")

	return writeJSON(w, http.StatusOK, map[string]string{"msg": "OK"})
}

func (s *Server) handleGetWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	category := chi.URLParam(r, "category")

	if _, exist := s.categories[category]; !exist {
		log.Warn("invalid category",
			zap.String("category", category),
		)
		return InvalidCategory(category)
	}

	weapons, err := s.svc.GetWeaponsByCategory(r.Context(), category)
	if err != nil {
		log.Error("GetWeaponsByCategory request failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("GetWeaponsByCategory request complited",
		zap.String("category", category),
		zap.Int("weapons_count", len(weapons)),
	)

	return writeJSON(w, http.StatusOK, types.Weapons{Weapons: weapons})
}

func (s *Server) handleGetWeapons(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	weapons, err := s.svc.GetWeapons(r.Context())
	if err != nil {
		log.Error("GetWeapons request failed",
			zap.Error(err),
		)
		return err
	}

	log.Info("GetWeapons request complited",
		zap.Int("weapons_count", len(weapons)),
	)

	return writeJSON(w, http.StatusOK, types.Weapons{Weapons: weapons})
}
