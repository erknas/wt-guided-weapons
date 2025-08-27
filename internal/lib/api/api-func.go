package api

import (
	"context"
	"net/http"
	"time"

	apierrors "github.com/erknas/wt-guided-weapons/internal/lib/api/api-errors"
)

type httpFunc func(w http.ResponseWriter, r *http.Request) error

func MakeHTTPFunc(fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*10)
		defer cancel()

		if err := fn(w, r.WithContext(ctx)); err != nil {
			if apiErr, ok := err.(apierrors.APIError); ok {
				WriteJSON(w, apiErr.StatusCode, apiErr)
			} else {
				errResp := map[string]any{
					"status_code": http.StatusInternalServerError,
					"msg":         "internal sever error",
				}
				WriteJSON(w, http.StatusInternalServerError, errResp)
			}
		}
	}
}
