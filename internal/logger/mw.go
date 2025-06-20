package logger

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/server/lib"
	apierrors "github.com/erknas/wt-guided-weapons/internal/server/lib/api-errors"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	Transport = "transport"
	Service   = "service"
	Storage   = "storage"
)

func MiddlewareRequestID(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-ID")

			if requestID == "" {
				requestID = strings.Replace(uuid.New().String(), "-", "", -1)
			}

			requestLogger := logger.With(zap.String("requestID", requestID))
			ctx := context.WithValue(r.Context(), "logger", requestLogger)

			w.Header().Set("X-Request-ID", requestID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func MiddlewareLogger(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			log := FromContext(r.Context(), "middleware/logger")

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			start := time.Now()
			defer func() {
				log.Info("request complited",
					zap.String("method", r.Method),
					zap.String("path", r.URL.Path),
					zap.String("user-agent", r.UserAgent()),
					zap.String("remote-addr", r.RemoteAddr),
					zap.Int("status code", ww.Status()),
					zap.Int("bytes written", ww.BytesWritten()),
					zap.Duration("duration", time.Since(start)),
				)
			}()

			next.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

func MiddlewareCategoryCheck(categories map[string]struct{}) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			category := chi.URLParam(r, "category")

			fmt.Println("Category", category)

			if _, exists := categories[category]; !exists {
				lib.WriteJSON(w, http.StatusBadRequest, apierrors.InvalidCategory(category))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
