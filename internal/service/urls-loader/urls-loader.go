package urlsloader

import (
	"encoding/json"
	"fmt"
	"os"
)

func Load(fileName string) (map[string]string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", fileName, err)
	}

	urls := make(map[string]string, 30)

	if err := json.Unmarshal(data, &urls); err != nil {
		return nil, fmt.Errorf("failed to decode data: %w", err)
	}

	if len(urls) != 30 {
		return nil, fmt.Errorf("invalid urls, urls count=%d", len(urls))
	}

	return urls, nil
}
