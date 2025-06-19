package tables

import (
	"encoding/json"
	"os"
)

const fileName = "tables.json"

type Tables struct {
	Tables map[string]string `json:"tables"`
}

func Load() (*Tables, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	tables := new(Tables)

	if err := json.Unmarshal(data, tables); err != nil {
		return nil, err
	}

	return tables, nil
}
