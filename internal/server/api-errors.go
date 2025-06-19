package server

import (
	"fmt"
	"net/http"
)

type APIError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func (e APIError) Error() string {
	return fmt.Sprintln(e.Message)
}

func NewApiError(statusCode int, err error) APIError {
	return APIError{
		StatusCode: statusCode,
		Message:    err.Error(),
	}
}

func InvalidCategory(category string) APIError {
	return NewApiError(http.StatusBadRequest, fmt.Errorf("invalid category: %s", category))
}
