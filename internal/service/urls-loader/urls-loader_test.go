package urlsloader

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	tmpDir := t.TempDir()

	tests := []struct {
		name        string
		prepareFile func(t *testing.T) string
		wantErr     bool
		errContains string
		wantLen     int
	}{
		{
			name: "Success with 30 URLs",
			prepareFile: func(t *testing.T) string {
				filePath := filepath.Join(tmpDir, "success.json")
				urls := make(map[string]string, 30)
				for i := 0; i < 30; i++ {
					urls[fmt.Sprintf("cat%d", i)] = fmt.Sprintf("http://example.com/%d", i)
				}
				data, err := json.Marshal(urls)
				require.NoError(t, err)
				err = os.WriteFile(filePath, data, 0644)
				require.NoError(t, err)
				return filePath
			},
			wantLen: 30,
		},
		{
			name: "File not exists",
			prepareFile: func(t *testing.T) string {
				return filepath.Join(tmpDir, "not_exists.json")
			},
			wantErr:     true,
			errContains: "failed to read file",
		},
		{
			name: "Invalid JSON",
			prepareFile: func(t *testing.T) string {
				filePath := filepath.Join(tmpDir, "invalid.json")
				err := os.WriteFile(filePath, []byte("{invalid}"), 0644)
				require.NoError(t, err)
				return filePath
			},
			wantErr:     true,
			errContains: "failed to decode data",
		},
		{
			name: "Invalid URLs count",
			prepareFile: func(t *testing.T) string {
				filePath := filepath.Join(tmpDir, "invalid_count.json")
				urls := map[string]string{"cat1": "http://example.com/1"}
				data, err := json.Marshal(urls)
				require.NoError(t, err)
				err = os.WriteFile(filePath, data, 0644)
				require.NoError(t, err)
				return filePath
			},
			wantErr:     true,
			errContains: "invalid urls, urls count=1",
		},
		{
			name: "Empty file",
			prepareFile: func(t *testing.T) string {
				filePath := filepath.Join(tmpDir, "empty.json")
				err := os.WriteFile(filePath, []byte(""), 0644)
				require.NoError(t, err)
				return filePath
			},
			wantErr:     true,
			errContains: "failed to decode data",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filePath := tt.prepareFile(t)

			result, err := Load(filePath)

			if tt.wantErr {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errContains)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Len(t, result, tt.wantLen)
			}
		})
	}
}
