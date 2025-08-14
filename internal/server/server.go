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
	apierrors "github.com/erknas/wt-guided-weapons/internal/lib/api/api-errors"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Servicer interface {
	UpsertWeapons(ctx context.Context) error
	WeaponsByCategory(ctx context.Context, category string) ([]*types.Weapon, error)
	SearchWeapon(ctx context.Context, query string) ([]types.SearchResult, error)
}

type Server struct {
	svc        Servicer
	categories map[string]struct{}
	log        *zap.Logger
}

func New(svc Servicer, urls map[string]string, log *zap.Logger) *Server {
	categories := make(map[string]struct{}, len(urls))

	for category := range urls {
		categories[category] = struct{}{}
	}

	return &Server{
		svc:        svc,
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

	s.log.Info("starting server", zap.String("port", cfg.ConfigServer.Port))

	select {
	case <-quitCh:
		shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return err
		}
		s.log.Info("server shutdown")
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
		r.Get("/upsert", makeHTTPFunc(s.handleUpdateWeapons))
		r.With(logger.MiddlewareCategoryCheck(s.categories)).Get("/weapons/{category}", makeHTTPFunc(s.handleGetWeaponsByCategory))
		r.Get("/weapons/search/{name}", makeHTTPFunc(s.handleSeachWeapon))
	})

	r.Handle("/*", http.FileServer(http.Dir("./static")))
}

type httpFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPFunc(fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			if apiErr, ok := err.(apierrors.APIError); ok {
				api.WriteJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal sever error",
				}
				api.WriteJSON(w, http.StatusInternalServerError, errResp)
			}
		}
	}
}
