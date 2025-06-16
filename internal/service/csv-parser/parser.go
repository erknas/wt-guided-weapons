package csvparser

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	"time"

	"github.com/erknas/wt-guided-weapons/internal/types"
)

func ParseTable(ctx context.Context, url string) ([]*types.Weapon, error) {
	data, err := readCSV(ctx, url)
	if err != nil {
		return nil, err
	}

	var weapons []*types.Weapon

	for i := range data[0][1:] {
		weapon, err := MapToWeapon(data, i+1)
		if err != nil {
			return nil, err
		}
		weapons = append(weapons, weapon)
	}

	return weapons, nil
}

func readCSV(ctx context.Context, url string) ([][]string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request to table [%s]: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return csv.NewReader(resp.Body).ReadAll()
}
