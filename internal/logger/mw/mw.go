package mw

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const requestIDKey = "request_id"

func RequestIDMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get("X-Request-ID")
		if requestID == "" {
			requestID = strings.Replace(uuid.New().String(), "-", "", -1)
		}

		ctx := context.WithValue(r.Context(), requestIDKey, requestID)

		requestLogger := logger.With(zap.String("requestID", requestID))
		ctx = context.WithValue(ctx, "logger", requestLogger)

		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
