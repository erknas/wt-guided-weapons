package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger/mw"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Servicer interface {
	InsertWeapons(context.Context) error
	UpdateWeapons(context.Context) error
	GetWeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
	GetWeapons(context.Context) ([]*types.Weapon, error)
}

type Server struct {
	svc Servicer
	log *zap.Logger
}

func New(svc Servicer, log *zap.Logger) *Server {
	return &Server{
		svc: svc,
		log: log,
	}
}

func (s *Server) Run(ctx context.Context, cfg *config.Config) error {
	router := chi.NewRouter()

	s.routes(router)

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

	s.log.Debug("starting server", zap.String("port", fmt.Sprintf("http://localhost%s", cfg.ConfigServer.Port)))

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

func (s *Server) routes(r *chi.Mux) {
	r.Use(func(next http.Handler) http.Handler {
		return mw.RequestIDMiddleware(s.log, next)
	})

	r.Route("/api", func(r chi.Router) {
		r.Post("/insert", makeHTTPFunc(s.handleInsertWeapon))
		r.Get("/weapons", makeHTTPFunc(s.handleGetWeapons))
		r.Get("/weapons/{category}", makeHTTPFunc(s.handleGetWeaponsByCategory))
	})

	r.Handle("/*", http.FileServer(http.Dir("./static")))
}

type httpFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPFunc(fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
