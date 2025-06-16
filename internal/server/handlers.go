package server

import (
	"net/http"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

func (s *Server) handleInsertWeapon(w http.ResponseWriter, r *http.Request) error {
	if err := s.svc.InsertWeapons(r.Context()); err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, nil)
}

func (s *Server) handleGetWeapons(w http.ResponseWriter, r *http.Request) error {
	weapons, err := s.svc.GetWeapons(r.Context())
	if err != nil {
		return err
	}

	return writeJSON(w, http.StatusOK, types.Weapons{Weapons: weapons})
}
