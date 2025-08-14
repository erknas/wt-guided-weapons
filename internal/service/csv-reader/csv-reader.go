package csvreader

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"
)

type HTTPReader struct{}

func (r *HTTPReader) Read(ctx context.Context, url string) ([][]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return csv.NewReader(resp.Body).ReadAll()
}
