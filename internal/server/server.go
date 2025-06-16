package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
)

type Servicer interface {
	InsertWeapons(context.Context) error
	UpdateWeapons(context.Context) error
	GetWeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
	GetWeapons(context.Context) ([]*types.Weapon, error)
}

type Server struct {
	svc Servicer
}

func New(svc Servicer) *Server {
	return &Server{
		svc: svc,
	}
}

func (s *Server) Run(ctx context.Context, cfg *config.Config) error {
	router := chi.NewRouter()

	router.Get("/insert", makeHTTPFunc(s.handleInsertWeapon))
	router.Get("/weapons", makeHTTPFunc(s.handleGetWeapons))

	srv := &http.Server{
		Addr:         cfg.ConfigServer.Port,
		Handler:      router,
		ReadTimeout:  cfg.ConfigServer.ReadTimeout,
		WriteTimeout: cfg.ConfigServer.WriteTimeout,
		IdleTimeout:  cfg.ConfigServer.IdleTimeout,
	}

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	errCh := make(chan error, 1)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
		errCh <- nil
	}()

	log.Printf("server starting on http://localhost%s\n", cfg.ConfigServer.Port)

	select {
	case <-quitCh:
		shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		log.Println("server shutdown")
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}

type httpFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPFunc(fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
