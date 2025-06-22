package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/config"
	"github.com/erknas/wt-guided-weapons/internal/logger"
	"github.com/erknas/wt-guided-weapons/internal/server/lib"
	apierrors "github.com/erknas/wt-guided-weapons/internal/server/lib/api-errors"
	"github.com/erknas/wt-guided-weapons/internal/service/tables"
	"github.com/erknas/wt-guided-weapons/internal/types"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Servicer interface {
	InsertWeapons(context.Context) error
	GetWeaponsByCategory(context.Context, string) ([]*types.Weapon, error)
}

type Server struct {
	svc        Servicer
	log        *zap.Logger
	categories map[string]struct{}
}

func New(svc Servicer, log *zap.Logger) (*Server, error) {
	tables, err := tables.Load()
	if err != nil {
		return nil, err
	}

	categories := make(map[string]struct{})
	for category := range tables.Tables {
		categories[category] = struct{}{}
	}

	return &Server{
		svc:        svc,
		log:        log,
		categories: categories,
	}, nil
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

	s.log.Info("starting server", zap.String("port", fmt.Sprintf("http://localhost%s", cfg.ConfigServer.Port)))

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
		r.Post("/insert", makeHTTPFunc(s.handleInsertWeapon))
		r.With(logger.MiddlewareCategoryCheck(s.categories)).Get("/weapons/{category}", makeHTTPFunc(s.handleGetWeaponsByCategory))
	})

	r.Handle("/*", http.FileServer(http.Dir("./static")))
}

type httpFunc func(w http.ResponseWriter, r *http.Request) error

func makeHTTPFunc(fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
		defer cancel()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			if apiErr, ok := err.(apierrors.APIError); ok {
				lib.WriteJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal sever error",
				}
				lib.WriteJSON(w, http.StatusInternalServerError, errResp)
			}
		}
	}
}
