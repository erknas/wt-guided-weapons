package csvreader

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
)

type Reader interface {
	Read(ctx context.Context, url string) ([][]string, error)
}

type HTTPReader struct {
	client *http.Client
}

func New() *HTTPReader {
	return &HTTPReader{
		client: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				MaxIdleConns:        10,
				MaxIdleConnsPerHost: 15,
				IdleConnTimeout:     90 * time.Second,
			},
		},
	}
}

func (r *HTTPReader) Read(ctx context.Context, url string) ([][]string, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	req.Header.Set("Accept", "text/csv, application/csv, text/plain")

	resp, err := r.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d, status: %s", resp.StatusCode, resp.Status)
	}

	reader := csv.NewReader(resp.Body)
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read csv: %w", err)
	}

	return data, nil
}
