package server

import (
	"context"
	"fmt"
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
	InsertWeapons(context.Context, *types.WeaponParams) error
	UpdateWeapons(context.Context, *types.WeaponParams) error
	GetWeaponsByCategory(context.Context, string) ([]*types.WeaponParams, error)
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

	log.Printf("server starting on [http://localhost:%s\n]", cfg.ConfigServer.Port)

	select {
	case <-quitCh:
		shutdownCtx, cancel := context.WithTimeout(ctx, time.Second*10)
		defer cancel()

		if err := srv.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutdown error: %w", err)
		}
		log.Println("server shutdown")
	case err := <-errCh:
		if err != nil {
			return err
		}
	}

	return nil
}
