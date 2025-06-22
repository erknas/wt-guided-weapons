package server

import (
	"fmt"
	"net/http"

	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/server/lib"
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

	return lib.WriteJSON(w, http.StatusOK, map[string]string{"msg": "OK"})
}

func (s *Server) handleGetWeaponsByCategory(w http.ResponseWriter, r *http.Request) error {
	log := logger.FromContext(r.Context(), logger.Transport)

	category := chi.URLParam(r, "category")
	fmt.Println(category)

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

	return lib.WriteJSON(w, http.StatusOK, types.Weapons{Weapons: weapons})
}
