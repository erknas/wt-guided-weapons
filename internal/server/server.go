package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/lib/api"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type WeaponsServicer interface {
	UpdateWeapons(ctx context.Context) error
	GetWeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error)
	SearchWeapons(ctx context.Context, query string) ([]types.SearchResult, error)
}

type VersionServicer interface {
	GetVersion(ctx context.Context) (types.LastChange, error)
}

type Server struct {
	weapons    WeaponsServicer
	version    VersionServicer
	categories map[string]struct{}
	log        *zap.Logger
}

func New(
	weapons WeaponsServicer,
	version VersionServicer,
	urls map[string]string,
	log *zap.Logger,
) *Server {
	categories := make(map[string]struct{}, len(urls))

	for category := range urls {
		categories[category] = struct{}{}
	}

	return &Server{
		weapons:    weapons,
		version:    version,
		categories: categories,
		log:        log,
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

	s.log.Info("Starting server", zap.String("port", cfg.ConfigServer.Port))

	select {
	case <-quitCh:
		shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		s.log.Info("Server shutdown")
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) routes(r *chi.Mux) {
	r.Use(logger.MiddlewareRequestID(s.log))
	r.Use(logger.MiddlewareLogger(s.log))

	r.Route("/api", func(r chi.Router) {
		r.Put("/update", api.MakeHTTPFunc(s.handleUpdateWeapons))
		r.With(logger.MiddlewareCategoryCheck(s.categories)).Get("/weapons/{category}", api.MakeHTTPFunc(s.handleGetWeaponsByCategory))
		r.Get("/weapons/search/{name}", api.MakeHTTPFunc(s.handleSeachWeapons))
		r.Get("/version", api.MakeHTTPFunc(s.handleGetVersion))
	})
}
